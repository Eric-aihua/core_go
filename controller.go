package main

import (
	"fmt"
)

// 主要用来演示go语言中的控制结构，if/else,switch,for等

func if_else_sample() {
	con1 := 10
	// else if 与else必须要与{符号在同一行
	if con1 < 0 {
		fmt.Println("below zero")
	} else if con1 < 5 && con1 > 0 {
		fmt.Println("middle")
	} else {
		fmt.Println("large")
	}
}

// switch 使用样例
func switch_sample() {
	c := 'c'
	switch c {
	case 'a':
		fmt.Println("Letter a!")
	case 'b':
		fmt.Println("Letter b!")
	default:
		fmt.Println("Unknown char!")
	}
}

func for_sample() {
	//使用带有初始值的for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("print Number:%v\n", i)
	}
	//使用range作为集合的for循环
	arr := []int{1, 2, 3, 4, 5}
	for index, item := range arr {
		fmt.Printf("index:%v   value:%v", index, item)
	}

}

func a() {
	fmt.Println("execute a")
}
func a1() {
	fmt.Println("execute a1")
}
func a2() {
	fmt.Println("execute a2")
}

// 演示defer的功能，defer定义的函数会在主函数执行完之后执行，且执行的顺和定义的顺序想反，定义的顺是a->a1->a2,最后执行的顺序是a2->a1->a
func b() {
	defer a()
	defer a1()
	fmt.Println("execute b")
	defer a2()
}

func main() {
	//	if_else_sample()
	//	switch_sample()
		for_sample()
	//b()
}
