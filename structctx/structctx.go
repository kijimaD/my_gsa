package structctx

import (
	"errors"
	"go/types"

	"github.com/gostaticanalysis/ident"
	"golang.org/x/tools/go/analysis"
)

var errAssert = errors.New("error: type assert failed")

var Analyzer = &analysis.Analyzer{
	Name: "structctx",
	Doc:  "detect struct ctx.",
	Run:  run,
	Requires: []*analysis.Analyzer{
		ident.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	m, ok := pass.ResultOf[ident.Analyzer].(ident.Map)
	if !ok {
		// type assertできないものはスルー
	}

	for o := range m {
		if ctxInField(o) {
			pass.Reportf(o.Pos(), "struct include context!")
		}
	}

	return nil, nil
}

func ctxInField(o types.Object) bool {
	// 未定義だとスルー
	// なくてもOK?
	if o == nil {
		return false
	}

	// ユニバースブロックで定義されてるとスルー
	// なくてもOK?
	if o.Parent() == types.Universe {
		return false
	}

	// エクスポートされてるとスルー
	if o.Exported() {
		return false
	}

	// 変数ではない・フィールドではない・無名関数だとスルー
	v, ok := o.(*types.Var)
	if !ok || !v.IsField() || v.Anonymous() {
		return false
	}

	var st types.Type
	// 構造体一覧をループ
	for n, s := range structs(v.Pkg()) {
		// structがフィールドvを持つとき(vがフィールドの場合)
		if hasField(s, v) {
			// スコープから構造体の識別子名で検索して、特定
			// stは構造体の型名。T1, T2, T3...
			st = v.Pkg().Scope().Lookup(n).Type()
			break
		}
	}

	if st == nil {
		return false
	}

	// 渡した型のポインター型を得る
	// *T1, *T2, *T3...
	stptr := types.NewPointer(st)

	// importしてるものでループ
	// テストコードだとb, context, fmt...
	for _, pkg := range v.Pkg().Imports() {
		// インポートしてるやつのインターフェース一覧でループ
		for _, i := range interfaces(pkg) {
			// 構造体のフィールドがインターフェースを実装してるものがあったらスルー
			if !i.Empty() &&
				(types.Implements(st, i) ||
					types.Implements(stptr, i)) {
				return false
			}
		}
	}

	// 警告
	return true
}

// copy from github.com/gostaticanalysis/analysisutil
// パッケージの構造体の一覧を返す
func structs(pkg *types.Package) map[string]*types.Struct {
	structs := map[string]*types.Struct{}

	for _, n := range pkg.Scope().Names() {
		o := pkg.Scope().Lookup(n)
		if o != nil {
			s, ok := o.Type().Underlying().(*types.Struct)
			if ok {
				structs[n] = s
			}
		}
	}

	return structs
}

// copy from github.com/gostaticanalysis/analysisutil
// structがフィールドを持っているか判定
func hasField(s *types.Struct, f *types.Var) bool {
	if s == nil || f == nil {
		return false
	}

	for i := 0; i < s.NumFields(); i++ {
		// 構造体->フィールドと、引数のvarを比較
		if s.Field(i) == f {
			return true
		}
	}

	return false
}

// copy from github.com/gostaticanalysis/analysisutil
// パッケージのインターフェース一覧取得
func interfaces(pkg *types.Package) map[string]*types.Interface {
	ifs := map[string]*types.Interface{}

	for _, n := range pkg.Scope().Names() {
		o := pkg.Scope().Lookup(n)
		if o != nil {
			i, ok := o.Type().Underlying().(*types.Interface)
			if ok {
				ifs[n] = i
			}
		}
	}

	return ifs
}
