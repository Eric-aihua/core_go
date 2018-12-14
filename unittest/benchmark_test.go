package testpackage

import "testing"

/* 基准测试的func有一下特点：
1. func都已Benchmark开头
2. 使用*testing.B作为参数
3. b.N为迭代次数
4. 执行测试的时候，在对应的目录下 执行go test -bench=.
*/
func BenchmarkStringFromJoin(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		StringFromJoin(100)
	}
}

func BenchmarkStringFromAssign(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		StringFromAssign(100)
	}
}

func BenchmarkStringFromBuffer(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		StringFromBuffer(100)
	}
}