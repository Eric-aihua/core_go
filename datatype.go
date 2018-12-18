package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 主要用来对go语言中的数据类型进行说明

func types_sample() {
	var b bool
	fmt.Println(b)
	// get data type
	fmt.Println(reflect.TypeOf(b))
	var s string = "hello"
	fmt.Println(s + " world")
	// define array
	var strArray [3]string
	strArray[0] = "one"
	strArray[1] = "two"
	strArray[2] = "three"
	fmt.Println(strArray)
	// 类型转换
	bs, err := strconv.ParseBool("true")
	fmt.Println(bs)
	fmt.Println(err)
}

func main() {
	types_sample()
}
