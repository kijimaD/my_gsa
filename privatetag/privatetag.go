package privatetag

import (
	"errors"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var errAssert = errors.New("error: type assert failed")

var Analyzer = &analysis.Analyzer{
	Name:     "privatetag",
	Doc:      "detect private tag.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errAssert
	}

	inspect.Preorder([]ast.Node{new(ast.GenDecl)}, func(n ast.Node) {
		decl, ok := n.(*ast.GenDecl)
		if ok {
			if decl.Tok == token.TYPE {
				for _, s := range decl.Specs {
					typespec, ok := s.(*ast.TypeSpec)
					if ok {
						structtype, ok := typespec.Type.(*ast.StructType)
						if ok {
							for _, field := range structtype.Fields.List {
								for _, name := range field.Names {
									if !ast.IsExported(name.Obj.Name) && len(field.Tag.Value) > 0 {
										pass.Reportf(name.NamePos, "detect private field tag!")
									}
								}
							}
						}
					}
				}
			}
		}
	})

	return nil, nil
}
