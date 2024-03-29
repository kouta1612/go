package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestComplexity(t *testing.T) {
	testcases := []struct {
		name       string
		code       string
		complexity int
	}{
		// テストケースをここに追加していく
		{
			name: "simple function",
			code: `
				package main
				func Double(n int) int {
					return n * 2
				}
			`,
			complexity: 1,
		},
		{
			name: "if statement",
			code: `
				package main
				func Double(n int) int {
					if n % 2 == 0 {
						return 0
					}
					return n
				}
			`,
			complexity: 2,
		},
		{
			name: "if & for statement",
			code: `
				package main
				func BubbleSort(l []int) []int {
					for i := 0; i < len(l)-1; i++ {
						for j := i + 1; j < len(l); j++ {
							if l[i] > l[j] {
								l[i], l[j] = l[j], l[i]
							}
						}
					}
					return l
				}
			`,
			complexity: 4,
		},
	}

	for _, testtestcase := range testcases {
		t.Run(testtestcase.name, func(t *testing.T) {
			a := GetAST(t, testtestcase.code)
			c := Count(a)
			if c != testtestcase.complexity {
				t.Errorf("got=%d, want=%d", c, testtestcase.complexity)
			}
		})
	}
}

// 引用: https://github.com/knsh14/gocc/blob/9078b24a5eb4377455473212ec67b8034de1439f/complexity/complexity_test.go
func GetAST(t *testing.T, code string) ast.Node {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range file.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatal("no function declear found")
	return nil
}
