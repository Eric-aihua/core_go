package testpackage

import (
	"bytes"
	"strings"
)

// 使用拼接的方式生成字符串
func StringFromAssign(n int)  {
	s:=""
	for i:=0;i<n;i++{
		s+="a"
	}
}

// 使用Buffer的方式生成字符串
func StringFromBuffer(n int)  string{
	s :=bytes.Buffer{}
	for i:=0;i<n;i++{
		s.WriteString("a")
	}
	return s.String()
}

// 使用Join的方式生成字符串
func StringFromJoin(n int) string {
	//申明切片
	s := make([] string,n)
	for i:=0;i<n;i++{
		s=append(s,"a")
	}
	//将切片中的元素都join起来
	return strings.Join(s,"")
}
/*
1.如果包名是main:func main可以执行，但是会报一个文件夹有两个包的定义的警告
2.如果包名是testpackage,main函数不会执行
*/
//func main() {
//	fmt.Println(StringFromBuffer(100))
//}

