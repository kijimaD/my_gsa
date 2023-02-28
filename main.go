package main

import (
	gopher_sample "github.com/kijimaD/my_gsa/gopher"
	"github.com/kijimaD/my_gsa/trashcomment"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(gopher_sample.Analyzer)
	unitchecker.Main(trashcomment.Analyzer)
}
