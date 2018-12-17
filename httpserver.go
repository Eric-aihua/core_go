package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
本代码主要用来演示用go实现httpserver的相关特性
*/

//申请一个处理请求的方法
func index(writer http.ResponseWriter, request *http.Request) {
	//如果跳转到index的url不是"/",则可认为是输入的路径错误，则将其转到404
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	} else {
		log.Print("received request from:", request.RemoteAddr)
		writer.Write([]byte("hello world"))
	}
}

// 1.收到请求，返回hello world
func hello2(writer http.ResponseWriter, request *http.Request) {
	log.Print("received request from:", request.RemoteAddr)
	writer.Write([]byte("hello world2\n"))
}

// 2. 根据不同的格式要求，返回不同格式的结果
// json请求 :curl -is -H "Accept:application/json" localhost:8888/formats
// json请求 :curl -is -H "Accept:application/xml" localhost:8888/formats
func formats(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println(request.Header)
	switch request.Header.Get("Accept") {
	case "application/json":
		writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		writer.Write([]byte(`{"name":"Jeason","age":15,"addr":"usa"}`))
	case "application/xml":
		writer.Header().Set("Content-Type", "application/xml;charset=utf-8")
		writer.Write([]byte(`<msg>hello world</msg>`))
	default:
		writer.Write([]byte("hello world"))
	}
}

// 2. 根据不同请求方法，返回不同的结果
// get请求 : curl -is -XGET 'localhost:8888/methods?name=eric&age=15'
// post请求 :  curl -is -XPOST -d "post request data" localhost:8888/methods
func methods(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		result := ""
		// 读取get的参数
		for k, v := range request.URL.Query() {
			result += "param key:" + k + " param value:" + strings.Join(v, ",") + "" +
				"\n"
		}
		writer.Write([]byte(result))
	case "POST":
		//读取POST请求的参数
		b, e := ioutil.ReadAll(request.Body)
		if e != nil {
			log.Fatal("Process POST method error!")
		} else {
			writer.Write([]byte("post result:" + string(b)))
		}
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Default result"))
	}
}

func main() {
	/*
		路由的定义将符合以下规律
		1. 一个请求没有匹配到任何path,将被路由到/
		2. 路由的匹配需要完全匹配，如果定义的为/hello/,则请求也必须为localhost:8888/hello/
	*/
	http.HandleFunc("/", index)
	http.HandleFunc("/hello2", hello2)
	http.HandleFunc("/formats", formats)
	http.HandleFunc("/methods", methods)
	log.Print("Start to receive request.............")
	http.ListenAndServe(":8888", nil)

}
