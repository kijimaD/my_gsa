package gophersample_test

import (
	"testing"

	"github.com/kijimaD/my_gsa/gophersample"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gophersample.Analyzer, "a")
}
