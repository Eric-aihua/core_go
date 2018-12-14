package main

import (
	"flag"
	"fmt"
	"os"
)

/*
演示基本命令行的使用方法，要点如下：
1. 通过flag实现命令行功能
2. 通过flag.[TYPE]的方式对输入的类型进行验证
3. 通过flag的Prase方法对参数进行解析
4. 通过指针使用参数
*/

func main() {
	s := flag.String("s", "默认值", "帮助文档")
	//自定义使用说明
	flag.Usage = func() {
		usageText := `命令行测试程序，使用方法为 cmdline1 [参数]，参数如下
		-s	指定需要输出的字符串
	`
		fmt.Fprintf(os.Stderr, "%s\n", usageText)
	}
	flag.Parse()    //参数转换
	fmt.Println(s)  //打印的地址
	fmt.Println(*s)
}
