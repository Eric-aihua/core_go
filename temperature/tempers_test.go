package temperature

import (
	"testing"
)

type tempreatureStruct struct {
	i        float64
	expected float64
}

// 定义测试集
var CToFTests = []tempreatureStruct{
	{4.1, 39.38},
	{10, 50},
}

// 定义测试集
var FToCTests = []tempreatureStruct{
	{32, 0},
	{50, 10},
}

//在命令行模式下执行：go test
func TestCToF(t *testing.T) {
	for _, tcase := range CToFTests {
		actual := CtoF(tcase.i)
		if actual != tcase.expected {
			t.Errorf("Expected value:%v actual value:%v", tcase.expected, actual)
		}
	}
}

//在命令行模式下执行：go test
func TestFToC(t *testing.T) {
	for _, tcase := range FToCTests {
		actual := FtoC(tcase.i)
		if actual != tcase.expected {
			t.Errorf("Expected value:%v actual value:%v", tcase.expected, actual)
		}
	}
}
