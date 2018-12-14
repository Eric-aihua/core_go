package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
在cmd1的基础上演示子命令的使用方式：
*/

func main() {
	//自定义使用说明
	flag.Usage = func() {
		usageText := `自命令示例，包括如下两个自命令：
		1. uppercase: 将对应的输入转换为全部大写，参数如下
				-s	指定需要输出的字符串
		2. uppercase: 将对应的输入转换为全部小写，参数如下
				-s	指定需要输出的字符串
	`
		fmt.Fprintf(os.Stderr, "%s\n", usageText)
	}
	//创建子命令
	upperCaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowerCaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 {
		//没有给子命令时，打印usage信息
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "uppercase":
		s := upperCaseCmd.String("s", "default", "转换参数")
		upperCaseCmd.Parse(os.Args[2:])
		fmt.Print(strings.ToUpper(*s))
	case "lowercase":
		s := lowerCaseCmd.String("s", "default", "转换参数")
		lowerCaseCmd.Parse(os.Args[2:])
		fmt.Print(strings.ToLower(*s))
	default:
		flag.Usage()
	}

}
