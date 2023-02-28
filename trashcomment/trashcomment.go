package trashcomment

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "trashcomment",
	Doc:      "transhcomment is for detecting useless comment.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var syntaxLen = 3 // exclude "// "

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, c := range f.Comments {
			for _, cl := range c.List {
				if len(cl.Text) < 5+syntaxLen {
					pass.Reportf(cl.Pos(), "useless comment!")
				}
			}
		}
	}

	return nil, nil
}
