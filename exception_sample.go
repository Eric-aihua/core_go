package main

//演示关于异常的处理以及声明的方式
import (
	"errors"
	"fmt"
	"io/ioutil"
)

/*
异常的处理
*/
func ProcessException() {
	//ReadFile的声明为：func ReadFile(filename string) ([]byte, error)。该方法将返回字节数组以及error
	file, err := ioutil.ReadFile("hello")
	//通过判断err是否为空，来确定是否存在错误
	if err == nil {
		fmt.Println(file)
	} else {
		fmt.Printf("Occuring Error when open file,error is: %v", err)
	}
}

//创建异常
func CreateError() {
	//通过errors.New创建error
	err := errors.New("this is a new error")
	if err != nil {
		fmt.Println(err)
	}
	//通过fmt.Errorf创建
	err2 := fmt.Errorf("hi %v", "newerror")
	if err2 != nil {
		fmt.Println(err2)
	}
}

//panic将使程序立刻退出
func PanicDemo() {
	err := errors.New("this is a heavy error")
	if err != nil {
		panic("由于严重系统错误，系统将立刻停止")
	}
	fmt.Println("前一条语句执行了panic,该语句不会执行")
}

func main() {
	//	ProcessException()
	//	CreateError()
	PanicDemo()
}
