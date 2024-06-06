package analyzer_test

import (
	"go-printf-func-name/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

// TestAll 是一个测试函数，用于运行静态代码分析器测试。
func TestAll(t *testing.T) {
	// os.Getwd 获取当前工作目录的完整路径。
	wd, err := os.Getwd()
	if err != nil {
		// 如果获取工作目录失败，使用 t.Fatalf 停止测试并报告错误。
		t.Fatalf("Failed to get wd: %s", err)
	}

	// filepath.Join 用于构造路径。这里它构造了指向 "testdata" 目录的路径。
	// filepath.Dir 获取路径中目录的部分，这里调用两次是为了向上移动两级目录。
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")

	// analysistest.Run 运行分析器测试。
	// 参数 t 是测试句柄，testdata 是测试数据目录，
	// Analyzer_package_name 是要测试的分析器，"name" 是测试用例的名字。
	analysistest.Run(t, testdata, analyzer.Analyzer_package_name, "name")
}
