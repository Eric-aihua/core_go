package main

import (
	"fmt"
)

//单个返回值
func return_single_value(x int, y int) int {
	return x + y
}

//返回数组
func returnArray() []int{
	a := make([]int,2)
	b := []int{1,2}
	a[0]=1
	a[2]=2
	//return a
	return b
}

//多返回值
func returnMultiValue() (string, int) {
	s := "string result"
	i := 3
	return s, i
}


//返回字典
func ReturnMap()  map[string][]string{
	relationGraph  := make(map[string][]string)
	relationGraph["you"] = []string{"alice", "bob", "claire"}
	relationGraph["bob"] = []string{"anuj", "peggy"}
	relationGraph["alice"] = []string{"peggy"}
	relationGraph["claire"] = []string{"thom", "jonny"}
	relationGraph["anuj"] = []string{}
	relationGraph["peggy"] = []string{}
	relationGraph["thom"] = []string{}
	relationGraph["jonny"] = []string{}
	//fmt.Println(relationGraph)
	return relationGraph
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
