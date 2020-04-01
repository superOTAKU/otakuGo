package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

/**
测试filepath包

测试环境：mac os
*/
func TestFilepath(t *testing.T) {
	//获取当前工作目录绝对路径
	cwd, _ := filepath.Abs(".")
	fmt.Println(cwd)

	//获取一个虚拟路径，基于当前工作目录
	testFilePath, _ := filepath.Abs("dir/test.txt")
	fmt.Println(testFilePath)

	//获取基本文件名，不带父路径
	testFileBasePath := filepath.Base(testFilePath)
	if testFileBasePath != "test.txt" {
		t.Error("filepath.Base error")
	} else {
		fmt.Println(testFileBasePath)
	}

	//获取父路径
	testFileDirPath := filepath.Dir(testFilePath)
	if testFileDirPath != cwd+"/"+"dir" {
		t.Error("filepath.Dir error")
	} else {
		fmt.Println(testFileDirPath)
	}

	//basepath为基准，为了到达targetpath，将路径调整为基于basepath的相对路径
	relPath, _ := filepath.Rel("a/b/c/d/", "a/b/test.txt")
	fmt.Println(relPath)
	//从目录 a/b/c/d开始，相对路径../../test.txt，则为 a/b/c/d/../../test.txt，即为a/b/test.txt
	if relPath != "../../test.txt" {
		t.Error("filepath.Rel error")
	}

}
