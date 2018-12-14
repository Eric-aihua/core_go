package testpackage

/*
测试文件有如下特征：
1：测试文件与功能代码放在一起，且名称为xxxx_test.go
2: 测试的func必须以Test开头
3：测试的func接受一个testing.T作为参数
4：执行测试的时候，在对应的目录下 执行go test命令
5: 执行测试的时候，在对应的目录下 执行go test -cover 命令可以得到代码覆盖率
*/

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

