package main

import (
	"fmt"
	"github.com/Eric-aihua/core_go/temperature"
)



func main() {
	//通过引入github上的包，并使用CToF方法
	fmt.Println(temperature.CtoF(10))
	//在temperature中，由于privateFunc以小写字母开头，将被认为是私有方法，该方法对于外部不可见。
	//fmt.Println(temperature.privateFunc())
}