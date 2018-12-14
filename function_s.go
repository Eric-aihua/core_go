package main

import (
	"fmt"
)

//单个返回值
func return_single_value(x int, y int) int {
	return x + y
}

//多返回值
func returnMultiValue() (string, int) {
	s := "string result"
	i := 3
	return s, i
}

//通过三个点来表达多个参数
func multiParam(numbers ...int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Printf("多输入参数的结果:%v \n", total)
}

// 返回带有名称的返回值,返回值的变量名为x,y
func returnNamedValue() (x, y string) {
	// 此处的x,y变量在func声明的部分已经做了初始化，不能再对变量进行声明
	x = "hello"
	y = "world"
	//不需要显示的指定返回值
	return
}

// 递归调用
func recursion(deep int) {
	if deep < 10 {
		fmt.Println("continue call!")
		deep += 1
		recursion(deep)
	} else {
		fmt.Println("Finished call!")
		return
	}

}

func main() {
	fmt.Printf("单返回值结果:%v \n", return_single_value(3, 5))
	var s, i = returnMultiValue()
	fmt.Printf("多返回值结果:%v-%v\n", s, i)
	multiParam(2, 3, 1, 4, 3, 6)
	fmt.Println(returnNamedValue())
	recursion(1)

}
