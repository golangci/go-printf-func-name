package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

// Analyzer_package_name 定义了一个静态代码分析器，用于检查不建议使用的包名。
var Analyzer_package_name = &analysis.Analyzer{
	Name:     "avoidBadPackageNames",                                    // 分析器的名称。
	Doc:      "Checks that packages are not named as 'err' or 'count'.", // 文档字符串，描述了分析器的功能。
	Run:      runPackageName,                                            // Run 函数，实现了分析器的逻辑。
	Requires: []*analysis.Analyzer{inspect.Analyzer},                    // 依赖的其他分析器列表。
}

// runPackageName 是分析器的核心函数，检查包名是否为 'err' 或 'count'。
// 使用了pass.ResultOf[inspect.Analyzer]来获取inspect.Analyzer的结果，该结果是一个 inspector.Inspector实例，用于遍历AST节点。
func runPackageName(pass *analysis.Pass) (interface{}, error) {
	// 获取 inspector 实例，它是由 inspect.Analyzer 提供的。
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// nodeFilter 定义了需要访问的 AST 节点类型，这里仅访问标识符。
	nodeFilter := []ast.Node{
		(*ast.Ident)(nil), // 只检查标识符节点。
	}

	// Preorder 方法遍历 AST 节点。
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		ident := node.(*ast.Ident) // 类型断言，获取当前节点的标识符。
		pkgName := ident.Name      // 获取标识符的名称，即包名。

		// 如果包名是 'err' 或 'count'，则报告这一情况。
		if pkgName == "err" || pkgName == "count" {
			pass.Reportf(ident.Pos(), "package named '%s' is discouraged, consider renaming it", pkgName)
		} else {
			// 如果包名不是 'err' 或 'count'，可以在这里添加其他逻辑。
			// fmt.Println(pkgName) // 例如，打印出包名。
		}
	})
	return nil, nil // 返回 nil 表示没有错误。
}
