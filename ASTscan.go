package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[1:] {
		if filePath == "--" { // to be able to run this like "go run main.go -- input.go"
			continue
		}

		f, err := parser.ParseFile(v.fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("Failed to parse file %s: %s", filePath, err)
		}

		ast.Walk(&v, f)
	}
}

// visitor 包含一个指向 token.FileSet 的指针，用于跟踪文件和位置信息。
type visitor struct {
	fset *token.FileSet
}

// Visit 是实现 ast.Visitor 接口的方法，用于访问 AST 中的每个节点。
func (v *visitor) Visit(node ast.Node) ast.Visitor {
	// 如果节点为 nil，说明访问已完成，返回 nil 结束遍历。
	if node == nil {
		return nil
	}

	// 创建一个缓冲区用于存储节点的文本表示。
	var buf bytes.Buffer
	// 使用 printer.Fprint 将节点的格式化表示写入缓冲区。
	printer.Fprint(&buf, v.fset, node)
	// 打印节点的文本表示和结构化表示。
	fmt.Printf("%s |\n %#v\n", buf.String(), node)

	// 返回访问者自身，以继续遍历。
	return v
}
