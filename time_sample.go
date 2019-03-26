package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	printNow()
	//sleep()
	//timeoutSample()
	//ticker()
	//format()
}

func format() {
	const shortForm = "2006-01-01 15:04:05"

	ts1 := "2006-01-01 12:04:05"
	ts2 := "2007-01-01 13:04:05"
	t1, err := time.Parse(shortForm, ts1)
	if err != nil {
		log.Fatal("parse format error!")
	}
	t2, err := time.Parse(shortForm, ts2)
	if err != nil {
		log.Fatal("parse format error!")
	}
	// 格式转换
	fmt.Println(t1)
	fmt.Println(t2)
	//时间比较
	fmt.Println(t1.After(t2))

}

// 计时器
func ticker() {
	c := time.Tick(1 * time.Second)
	for t := range c {
		fmt.Printf("Ticker time:%s\n", t)
	}
}

// 超时功能,可和goroutine一起使用
func timeoutSample() {
	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("timeout print")
		//return
	}
}

// 休眠
func sleep() {
	fmt.Println("begin to sleep 3 second")
	time.Sleep(3 * time.Second)
	fmt.Println("finished sleep 3 second")
}

//打印时间
func printNow() {
	fmt.Println(time.Now().Format("2006-01-02-15"))
	fmt.Printf("%v\n", time.Now())                  //打印当前时间
	fmt.Printf("%v\n", time.Now().Unix())           //打印时间戳
	fmt.Printf("%v\n", time.Now().AddDate(0, 1, 0)) //当前时间加一个月
}
