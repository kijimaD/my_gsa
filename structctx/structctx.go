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
	if o == nil {
		return false
	}

	// ユニバースブロックで定義されてるとスルー
	if o.Parent() == types.Universe {
		return false
	}

	// エクスポートされてるとスルー
	if o.Exported() {
		return false
	}

	// 変数だとスルー
	v, isVar := o.(*types.Var)
	if !isVar || !v.IsField() || v.Anonymous() {
		return false
	}

	// 警告
	return true
}
