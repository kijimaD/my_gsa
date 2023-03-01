package trashcomment_test

import (
	"testing"

	"github.com/kijimaD/my_gsa/trashcomment"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, trashcomment.Analyzer, "a")
}
