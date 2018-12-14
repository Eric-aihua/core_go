package main

import (
	"testing"
)

type helloStruct struct {
	name        string
	expected string
}

// 定义测试集
var HelloTests = []helloStruct{
	{"eric", "hello eric"},
	{"simon", "hello simon"},
}


//在命令行模式下执行：go test
func TestCToF(t *testing.T) {
	for _, tcase := range HelloTests {
		actual := Hello(tcase.name)
		if actual != tcase.expected {
			t.Errorf("Expected value:%v actual value:%v", tcase.expected, actual)
		}
	}
}

