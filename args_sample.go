package main

import (
	"fmt"
	"os"
)

/*
通过os.Args读取程序输入，运行步骤如下：
1：go build args_sample.go
2: 执行：args_sample.exe aa bb cc dd
3: 输出如下：
    index: 0  value: args_sample.exe  //第一个参数为文件名
    index: 1  value: aa
    index: 2  value: bb
    index: 3  value: cc
    index: 4  value: dd

*/

func main() {
	fmt.Print(len(os.Args))
	fmt.Print(os.Args)
	for i,arg :=range os.Args{
		fmt.Println("index:",i," value:",arg)
	}
}
