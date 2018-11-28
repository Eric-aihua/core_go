package main

//主要用来演示数组切片,map的功能
import (
	"fmt"
)

//初始完大小后，数组大小不能发生变化
func array_sample() {
	a := []int{1, 2} //等同于a := [2]int{1, 2}
	fmt.Println(a)
	/**
	// a在初始化的时候指定了大小为2,数据的大小不能自动伸缩，当对a[2]进行赋值时，会报index out of range错误
	a[2] = 3
	fmt.Println(a)
	**/
}

// 切片可以理解为增强版的数组，初始化大小后，还可以对数组进行扩容，删除等操作
func slice_sample() {
	//初始化一个长度为2,类型为string的切片
	slice := make([]string, 2)
	fmt.Println(slice)

	slice[0] = "Eric"
	slice[1] = "Simon"
	// 通过append添加新的元素，且一次可以添加多个
	slice = append(slice, "Jay", "Rose", "Tom")
	fmt.Println(slice)

	// 删除第二个元素
	slice = append(slice[:2], slice[2+1:]...)
	fmt.Println(slice)
	//切片的拷贝
	slice_clone := make([]string, 2)
	copy(slice_clone, slice)
	fmt.Println(slice_clone)
}

func main() {
	//	array_sample()
	slice_sample()
}
