package main

import (
	"github.com/kijimaD/my_gsa/gophersample"
	"github.com/kijimaD/my_gsa/trashcomment"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(trashcomment.Analyzer)
	unitchecker.Main(gophersample.Analyzer)
}
