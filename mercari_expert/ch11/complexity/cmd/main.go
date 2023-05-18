package main

import (
	"go_practice/mercari_expert/ch11/gocc"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gocc.Analyzer)
}
