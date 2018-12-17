package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/*主要演示http client的使用*/

const BASEURL string = "http://localhost:8888/methods"

//get方法
func getRequest() {
	res, err := http.Get(BASEURL + "?name=eric&age=15")
	if err != nil {
		log.Fatal("Get Request Error!")
	} else {
		resData, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			log.Fatal("Get Request Error!")
		}
		fmt.Println(string(resData))
	}
}

//发起post 请求
func postRequest() {
	body := strings.NewReader(`{"name","eric"}`)
	res, err := http.Post(BASEURL, "application/json", body)
	if err != nil {
		log.Fatal("Response Request Error!")
	} else {
		resData, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			log.Fatal("Response Request Error!")
		}
		fmt.Println(string(resData))
	}
}

//使用client对象发起请求，该方式可以对request进行更好的控制
func clientRequest() {
	//设置timeout时间
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:8888/formats", nil)
	//设置request的header信息
	req.Header.Add("Accept", "application/json")
	if err != nil {
		log.Fatal("Create request error")
	}
	res, e2 := client.Do(req)
	if e2 != nil {
		log.Fatal("Get response error")
	}
	p1 := Person{}
	json.NewDecoder(res.Body).Decode(&p1)
	fmt.Printf("%+v\n", p1)
}

func main() {
	//getRequest()
	//postRequest()
	clientRequest()
}
