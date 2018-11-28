package main

/*
对go的变量声明规则进行演示
*/
import (
	"fmt"
	"reflect"
)

/*常量的声明*/
const NAME string = "ERIC"

/*
变量声明的方式
*/
func define_methods() {
	//完整声明
	var v1 string = "full define"
	//一次声明多个变量
	var v2, v3 string = "v2", "v3"
	var (
		v4 string = "v4"
		v5 int    = 123
	)
	//不指定类型的声明
	var v6 = "v6"
	//简短声明
	v7 := "v7"
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println(v7)
}

// 指针的使用方式,参数为指向string类型的指针
func point_sample(x *string) {
	//打印变量内存的地址
	fmt.Println(x)
	//打印变量的值
	fmt.Println(*x)
	//修改指针位置的值
	*x = "other value"
	fmt.Println(*x)
	fmt.Println(x)
}

func not_point_params(x string) {
	// 传递的变量在使用前，会复制一个值相同的变量，两个变量有两个不同的地址
	fmt.Println(&x)

}
func main() {
	// 打印常量
	fmt.Println(NAME)
	//	define_methods()
	point_v := "point sample"
	//point_sample(point_v)
	fmt.Println(reflect.TypeOf(&point_v))
	fmt.Println(&point_v)
	point_sample(&point_v)
	not_point_params(point_v)
}
