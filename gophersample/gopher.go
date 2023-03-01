// https://docs.google.com/presentation/d/1I4pHnzV2dFOMbRcpA-XD0TaLcX6PBKpls6WxGHoMjOg/edit#slide=id.g6298a0e67590aa1e_6
// 識別子gopherを検出する

package gophersample

import (
	"errors"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var errAssert = errors.New("error: type assert failed")

var Analyzer = &analysis.Analyzer{
	Name:     "simple",
	Doc:      "simple is simple analyzer.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errAssert
	}

	inspect.Preorder([]ast.Node{new(ast.Ident)}, func(n ast.Node) {
		i, ok := n.(*ast.Ident)
		if ok {
			if i.Name == "gopher" {
				pass.Reportf(n.Pos(), "identifier is gopher!")
			}
		}
	})

	return nil, nil
}
