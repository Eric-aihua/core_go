package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

/*
1:通过执行go get github.com/golang/example/stringutil 安装stringutil包，并通过import进行使用,go get 的用法如下：
(1)go get xxxxxxx安装一个包
(2)go get -u :更新项目依赖的所有包
(3)go get -u xxx:只更新项目依赖的某个包
(4)go get -u all :更新系统中所有的包
*/
func StringRevert(s string) {
	//Reverse方法位于stringutil/reverse.go文件中，reverse.go文件的package声明为stringutil
	fmt.Println(stringutil.Reverse(s))
}

func main() {
	StringRevert("Eric!")
}
