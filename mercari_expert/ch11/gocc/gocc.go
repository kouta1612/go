package gocc

import (
	"go/ast"

	"go/mercari_expert/ch11/complexity"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gocc",
	Doc:      "checks cyclomatic complexity",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		count := complexity.Count(n)
		if count >= 10 {
			fd := n.(*ast.FuncDecl)
			pass.Reportf(n.Pos(), "function %s complexity=%d", fd.Name, count)
		}
	})

	return nil, nil
}
