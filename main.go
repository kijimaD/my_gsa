package main

import (
	"bytes"
	"fmt"

	"github.com/kijimaD/my_gsa/gophersample"
	"github.com/kijimaD/my_gsa/trashcomment"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(trashcomment.Analyzer, gophersample.Analyzer)

	// test ================

	// trashcomment
	// a

	// gophersample
	gopher := "sample"
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, gopher)
}
