package argcount

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "argcount",
	Doc:      "argcount is for detecting too many arguments.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var warnlen = 5

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, d := range f.Decls {
			funcdecl, ok := d.(*ast.FuncDecl)

			if !ok {
				continue
			}

			for _, p := range funcdecl.Type.Params.List {
				l := len(p.Names)
				if l >= warnlen {
					pass.Reportf(funcdecl.Pos(), "too many arguments!")
				}
			}
		}
	}

	return nil, nil
}
