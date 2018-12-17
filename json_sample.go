package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 通过json:XXX,XX 指定在obj->json 时，key的名称，例如Name,转化为json后将变为name
type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,0"`
	Addr string `json:"addr,omitempty"` //如果指定了omitempty,当对象没有对Addr赋值时，在编码的时候将不会放到Json对象中
}

// 1.将对象进行编码
func encodeJson() {
	p1 := Person{"Eric", 10, "china"}
	p2 := Person{}
	p2.Name = "Simon"

	fmt.Printf("%+v\n", p1)
	//将Person对象编码为Json对象
	jsonObj, err := json.Marshal(p1)
	jsonObj2, err := json.Marshal(p2)
	if err != nil {
		log.Fatal("encod error!")
	}
	fmt.Printf("%s\n", jsonObj)
	fmt.Printf("%s\n", jsonObj2)
}

//2.对json字符串进行解码
func decodeJson() {
	jsonStr := `{"name":"Eric","age":10,"addr":"china"}`
	p1 := Person{}
	json.Unmarshal([]byte(jsonStr), &p1)
	fmt.Printf("%+v\n", p1)
}

//3.解析API的Json

func decodeApiJson() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8888/formats", nil)
	//设置header信息
	req.Header.Add("Accept", "application/json")
	if err != nil {
		log.Fatal("Create request error")
	}
	res, e2 := client.Do(req)
	if e2 != nil {
		log.Fatal("Get response error")
	}
	defer res.Body.Close()
	p1 := Person{}
	json.NewDecoder(res.Body).Decode(&p1)
	fmt.Printf("%+v\n", p1)
}

func main() {
	//encodeJson()
	//decodeJson()
	decodeApiJson()
}
