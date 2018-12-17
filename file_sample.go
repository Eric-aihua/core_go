package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// 通过json:XXX,XX 指定在obj->json 时，key的名称，例如Name,转化为json后将变为name
type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,0"`
	Addr string `json:"addr,omitempty"` //如果指定了omitempty,当对象没有对Addr赋值时，在编码的时候将不会放到Json对象中
}

/*
主要用来对文件解析做演示
*/

func main() {
	//readAllFile()
	//writeFile()
	//dirList()
	readConfig()
}

// 读取json的config配置
func readConfig() {
	files, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("read dir error")
	}
	p1 := Person{}
	json.Unmarshal(files, &p1)
	fmt.Println(p1)

}

// 3.列出目录文件
func dirList() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal("read dir error")
	}
	for _, file := range files {
		fmt.Printf("name:%s  mode:%s\n", file.Name(), file.Mode())
	}

}

// 2. 写文件,该方式会覆盖原有的内容，原内容不会保留
func writeFile() {
	err := ioutil.WriteFile("test", []byte("hello world"), 0644)
	if err != nil {
		log.Fatal("write file error")
	}

}

// 1.一次读取整个文件,文件比较大会有问题
func readAllFile() {
	fileBytes, err := ioutil.ReadFile("README2.md")
	if err != nil {
		log.Fatal("open file error")
	}
	fmt.Println(string(fileBytes))
}
