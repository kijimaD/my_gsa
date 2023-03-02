package trashcomment

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "trashcomment",
	Doc:      "transhcomment is for detecting useless function comment.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

const syntaxLen = 3 // exclude string "// "

const warnLen = 6

// 複数行に対応できてない.
func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, d := range f.Decls {
			funcdecl, ok := d.(*ast.FuncDecl)

			if !ok {
				continue
			}

			if funcdecl.Doc == nil {
				continue
			}

			for _, c := range funcdecl.Doc.List {
				if len(c.Text) < warnLen+syntaxLen {
					pass.Reportf(c.Pos(), "useless function comment!")
				}
			}
		}
	}

	return nil, nil
}
