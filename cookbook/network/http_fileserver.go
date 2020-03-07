package main

import "net/http"

/*
@Time : 2020-03-04 23:11 
@Author : sunaihua

使用http库实现的fileserver,登录浏览器输入localhost:8080就可以对文件进行浏览
*/

func main(){
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080",nil)
}