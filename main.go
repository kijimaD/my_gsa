package main

import (
	gophersample "github.com/kijimaD/my_gsa/gopher"
	"github.com/kijimaD/my_gsa/trashcomment"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(gophersample.Analyzer)
	unitchecker.Main(trashcomment.Analyzer)
}
