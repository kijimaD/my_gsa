package main

import (
	"bytes"
	"fmt"

	"github.com/kijimaD/my_gsa/argcount"
	"github.com/kijimaD/my_gsa/gophersample"
	"github.com/kijimaD/my_gsa/structctx"
	"github.com/kijimaD/my_gsa/trashcomment"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(trashcomment.Analyzer, gophersample.Analyzer, argcount.Analyzer, structctx.Analyzer)

	demo("1", "2", "3", "4", "5")
}

// warn.
func demo(a, b, c, d, e string) {
	gopher := "sample"
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, gopher)
}
