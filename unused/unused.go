// https://github.com/gostaticanalysis/unused

package unused

import (
	"go/types"

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
	return false
}
