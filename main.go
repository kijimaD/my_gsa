package main

import (
	gopher_sample "github.com/kijimaD/my_gsa/gopher"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gopher_sample.Analyzer) }
