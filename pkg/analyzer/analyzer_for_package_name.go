package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer_package_name = &analysis.Analyzer{
	Name:     "avoidBadPackageNames",
	Doc:      "Checks that packages are not named as 'err' or 'count'.",
	Run:      runPackageName,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func runPackageName(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		ident := node.(*ast.Ident)
		pkgName := ident.Name
		if pkgName == "err" || pkgName == "count" {
			pass.Reportf(ident.Pos(), "package named '%s' is discouraged, consider renaming it", pkgName)
		} else {
			// fmt.Println(pkgName)
		}
	})

	return nil, nil
}
