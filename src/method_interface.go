package main

import (
	"fmt"
	"strconv"
)

//定义一个Roboot接口
type Roboot interface {
	PowerOff()
	PowerOn()
}

type MachineRoboot struct {
	name string
}
type EletricRoboot struct {
	name string
}

// 实现Roboot接口
func (mr *MachineRoboot) PowerOff() {
	fmt.Println(mr.name + " machine roboot poweroff")
}

// 实现Roboot接口
func (mr *MachineRoboot) PowerOn() {
	fmt.Println(mr.name + " machine roboot poweron")
}

// 实现Roboot接口
func (er *EletricRoboot) PowerOff() {
	fmt.Println(er.name + " machine roboot poweroff")
}

// 实现Roboot接口
func (er *EletricRoboot) PowerOn() {
	fmt.Println(er.name + " machine roboot poweron")
}

// 该方法接受Robot类型的参数
func Restart(r Roboot) {
	r.PowerOff()
	r.PowerOn()
}

//定义一个movie的结构体
type Movie struct {
	name   string
	rating float64
}

// 定义movie的方法,movie对象可以字节调用该方法
func (movie *Movie) summary() string {
	summary := "Name:" + movie.name + " Rating:" + strconv.FormatFloat(movie.rating, 'f', 1, 64)
	return summary
}

func MethodSample() {
	m := Movie{
		name:   "Hero",
		rating: 6.5,
	}
	//movie的对象可以直接调用summary方法
	fmt.Println(m.summary())

}

func InterfaceSample() {
	mr := MachineRoboot{
		name: "MMR",
	}
	er := EletricRoboot{
		name: "ERR",
	}
	mr.PowerOn()
	mr.PowerOff()
	er.PowerOn()
	er.PowerOff()
	//使用对象的地址，调用接受roboot实例的方法
	Restart(&mr)
	Restart(&er)
}

func main() {
	//	MethodSample()
	InterfaceSample()
}
