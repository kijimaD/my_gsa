// https://github.com/gostaticanalysis/unused

package unused

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/ident"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "unused",
	Doc:      "unused is detect unused ident.",
	Run:      run,
	Requires: []*analysis.Analyzer{ident.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	identm, ok := pass.ResultOf[ident.Analyzer].(ident.Map)
	if ok {
		// mにはmapで識別子が入っている
		for o := range identm {
			// どこか別で呼び出されるとlen(identm[o])は1より大きくなる
			// 一度しか出現してないと1になる(定義の1回か)
			if !skip(o) && len(identm[o]) == 1 {
				// identm[o]には*ast.Identの配列が入ってる
				n := identm[o][0]
				pass.Reportf(n.Pos(), "%s is unused", n.Name)
			}
		}
	}

	return nil, nil
}

func skip(o types.Object) bool {
	// 公開関数・フィールドは除外
	if o == nil || o.Parent() == types.Universe || o.Exported() {
		return true
	}

	switch o := o.(type) {
	// パッケージ名は1つしかないので除外
	case *types.PkgName:
		return true
	case *types.Var:
		// varがくるときfieldがくるときがある
		// パッケージスコープにない?
		if o.Pkg().Scope() != o.Parent() &&
			!(o.IsField() && !o.Anonymous() && isFieldInNamedStruct(o)) {
			return true
		}
	case *types.Func:
		// main
		// mainパッケージのmain()は呼び出しはない
		if o.Name() == "main" && o.Pkg().Name() == "main" {
			return true
		}

		// init
		if o.Name() == "init" && o.Pkg().Scope() == o.Parent() {
			return true
		}

		// method
		sig, ok := o.Type().(*types.Signature)
		if ok {
			if recv := sig.Recv(); recv != nil {
				for _, i := range analysisutil.Interfaces(o.Pkg()) {
					// インターフェースの中の、関数を実装していれば使っているということになる
					if i == recv.Type() || (types.Implements(recv.Type(), i) && has(i, o)) {
						return true
					}
				}
			}
		}

	}

	return false
}

func has(intf *types.Interface, m *types.Func) bool {
	for i := 0; i < intf.NumMethods(); i++ {
		if intf.Method(i).Name() == m.Name() {
			// インターフェースが実装しているすべてのメソッドの中と、関数の名前が一致するかどうか
			return true
		}
	}
	return false
}

func isFieldInNamedStruct(v *types.Var) bool {
	structs := allNamedStructs(v.Pkg())
	for _, s := range structs {
		for i := 0; i < s.NumFields(); i++ {
			if s.Field(i) == v {
				return true
			}
		}
	}
	return false
}

func allNamedStructs(pkg *types.Package) []*types.Struct {
	var structs []*types.Struct

	for _, n := range pkg.Scope().Names() {
		o := pkg.Scope().Lookup(n)
		if o != nil {
			switch t := o.Type().(type) {
			case *types.Named:
				switch u := t.Underlying().(type) {
				case *types.Struct:
					structs = append(structs, u)
				}
			}
		}
	}

	return structs
}
