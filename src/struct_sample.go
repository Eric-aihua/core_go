package main

/*
对struct进行演示
*/
import (
	"fmt"
	"reflect"
)

//定义一个movie的结构体
type Movie struct {
	name   string
	rating float32
}

//定义一个嵌套结构的struct

type Address struct {
	country string
	city    string
	street  string
}

type Student struct {
	name    string
	address Address
}

//初始化
func InitMovie() {
	// 所有属性都用默认值
	m1 := Movie{}
	//初始化的时候指定属性的值
	m2 := Movie{
		name:   "Hero",
		rating: 6.5,
	}
	//逐个属性赋值
	m3 := Movie{}
	m3.name = "Logan"
	m3.rating = 8.8
	//使用new的方式创建movie
	m4 := new(Movie)
	m4.name = "Revenant"
	m4.rating = 7.5
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
}

func NestedStruct() {
	student := Student{
		name:    "Eric",
		address: Address{country: "china", city: "wuhan", street: "guanggu"}}
	fmt.Println(student.name)
	fmt.Println(student.address.country)
	fmt.Println(student)
	//打印结构体类型
	fmt.Println(reflect.TypeOf(student))
}

func main() {
	//	InitMovie()
	NestedStruct()
}
