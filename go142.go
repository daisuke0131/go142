package go142

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "go142 is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "go142",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			fmt.Println(n)
			fmt.Println(pass.Fset.Position(n.Pos()))
			fmt.Println("Defs:", pass.TypesInfo.Defs[n])
			fmt.Println("Uses:", pass.TypesInfo.Uses[n])
		}
	})

	return nil, nil
}
