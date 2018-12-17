package main

//演示字符串的一些相关特性
import (
	"bytes"
	"fmt"
	"strings"
)

/*
字符串分为:
解释性字符串：用""表示,会对特殊字符进行处理，例如\t
原生性字符串：用``表示,不会对特殊字符进行处理
*/
func StringType() {
	s1 := "This is a String\n,new line"
	s2 := `This is a String\n,new line
	        new line again`
	fmt.Println(s1)
	fmt.Println(s2)
}

//字符串拼接
func StringConcat() {
	//使用+号进行拼接
	s1 := "Name:" + "Eric"
	//如果有很多字符串可以用buffer进行拼接
	//	var s2 bytes.Buffer
	s2 := bytes.Buffer{}
	for i := 0; i < 100; i++ {
		s2.WriteString("s")
	}
	fmt.Println(s1)
	fmt.Println(s2.String())

}

// 主要演示关于编码相关的一些操作
func StringCode() {
	s1 := "hello"
	s2 := "我爱你"
	fmt.Println(len(s1))      //5
	fmt.Println(len(s2))      //9,每个汉字三个字节
	fmt.Println(s1[1])        //101
	fmt.Printf("%q\n", s1[1]) //输出字符:e
	fmt.Printf("%b\n", s1[1]) //输出二进制:1100101
	fmt.Println(s1[2:])       //切片操作:llo
	fmt.Println(s2[1])        //136
}

// 字符串常用操作
func StringOperations() {
	s1 := "Hello  "
	fmt.Println(strings.TrimSpace(s1)) //去除空格
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.Index(s1, "ll"))
}

//字符串数组拼接为字符串
func ArrayJoin() {
	arr := []string{"a", "b", "c"}
	fmt.Println(strings.Join(arr, ""))
}

func ByteConvert() {
	s := "sunaihua"
	// 将字符串转换为byte数组
	bs := []byte(s)
	fmt.Println(bs)
	//byte数组转为str
	fmt.Print(string(bs))
}

func main() {
	//StringType()
	//	StringConcat()
	//	StringCode()
	//StringOperations()
	//ArrayJoin()
	ByteConvert()
}
