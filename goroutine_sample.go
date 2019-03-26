package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
主要用来演示goroutine的特性
*/

func PrintService(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "PrintService Result"
	fmt.Println("被调用函数的输出")
}

func pinger(ch chan string) {
	t := time.NewTicker(1 * time.Second)
	//	time.NewTicker(1 * time.Second)
	for {
		ch <- "ping...."
		<-t.C
	}
}

func PrintBufferService(ch chan string) {
	for msg := range ch {
		fmt.Println("Channel Buffer item:" + msg)
	}
	fmt.Println("被调用函数的输出")
}

func ProcessInt(i int) {
	time.Sleep(2 * time.Second)
	result := "Result:" + strconv.Itoa(i)
	fmt.Println("被调用函数的输出:" + result)
}

func SimpleFunc() {
	time.Sleep(2 * time.Second)
	fmt.Println("被调用函数的输出")
}

func Ping1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Ping1 Result"
	fmt.Println("Ping1函数的输出")
}
func Ping2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Ping2 Result"
	fmt.Println("Ping2函数的输出")
}

//1:没有使用阻塞机制的协成,SimpleFunc的Print不会被输出，因为go SimpleFunc()的调用是非阻塞的，调用后将直接执行主函数的Print,然后结束程序
func NoStopGoroutine() {
	go SimpleFunc()
	fmt.Println("主函数的输出")
}

//2:使用Time.Sleep对主函数进行阻塞，保证被调用函数Print的执行
func UseTimeGoroutine() {
	go SimpleFunc()
	fmt.Println("主函数的输出")
	time.Sleep(3 * time.Second)
}

//3：使用chan进行消息的传递
func UseChannelGoroutine() {
	//申声明一个string类型的chan,该chan只能发送string类型的消息
	ch := make(chan string)
	go PrintService(ch)
	result := <-ch
	fmt.Println("主函数的输出:" + result)
}

//4：使用chan buffer进行消息的传递
func UseChannelBufferGoroutine() {
	//申声明一个string类型的chan buffer,该buffer可存储2个消息，当超过2个消息的时候会产生错误
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	//	ch <- "world" //当buffer的数量超过2时，会报错
	go PrintBufferService(ch)
	time.Sleep(time.Second * 1) //需要sleep 1s等待调用函数的执行
	fmt.Println("主函数的输出")
}

//5. 使用协成同时处理100个int值
func ForLoopByGorouine() {
	fmt.Println("开始执行For循环")
	for i := 0; i < 100; i++ {
		go ProcessInt(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("For循环结束")
}

//6.非阻塞的goroutine,由于是非阻塞的，只会输出一个ping

func NonBlockingGoroutine() {
	message := make(chan string)
	go pinger(message)
	msg := <-message
	fmt.Println(msg)
}

//7.阻塞的goroutine,将会每隔一秒输出一个ping

func BlockingGoroutine() {
	message := make(chan string)
	go pinger(message)
	for {
		msg := <-message
		fmt.Println(msg)
	}

}

// 创建一个只可以读的chan
func OnlyReadChan(ch <-chan string) {
	// 从ch读取值
	message := <-ch
	//	ch <- "hi " // 不能向ch写
	fmt.Println(message)
}

// 创建一个只可以写的chan
func OnlyWriteChan(ch chan<- string) {

	//	message := <-ch // 不能读
	ch <- "hi " //写
}

// 创建一个只可读可写的chan

//8.演示只读，只写，可读可写类型chan的创建方法
func ChanTypeSample() {
	ch := make(chan string, 1)
	ch <- "hello"
	go OnlyReadChan(ch)
	ch2 := make(chan string, 1)
	go OnlyWriteChan(ch2)
	result := <-ch2
	fmt.Println("从只写chan获取的结果:" + result)
	time.Sleep(1 * time.Second)
}

// 9. 演示使用select来接受不同chan的数据，适用于根据最先收到消息而采取相应的措施，有两点需要注意：（1）select不具有阻塞性，当有多个case时，一个case执后程序将退出
//（2）可以使用time.After方法来防止被调用函数执行时间较长

func SelectSample() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go Ping1(ch1)
	go Ping2(ch2)
	select {
	case msg := <-ch1:
		fmt.Println("收到ch1的消息:" + msg)
	case msg2 := <-ch2:
		fmt.Println("收到ch2的消息:" + msg2)
	case <-time.After(500 * time.Millisecond):
		//如果500毫秒没收到ch1,ch2的消息，程序将退出
		fmt.Println("超时输出")
	}
}

// 10:演示在Select结构中增加一个可以在特定情况下让程序退出的chan,例如：遍历完文件后，或是执行一定的次数后
func SelectWithStopSample() {
	ch1 := make(chan string)
	stop := make(chan bool)
	go pinger(ch1)
	go func() {
		//sleep 10 秒后发出退出的指令
		time.Sleep(10 * time.Second)
		fmt.Println("满足程序退出的条件，准备退出!")
		stop <- true
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println("收到ch1的消息:" + msg)
		case <-stop:
			fmt.Println("收到stop指令,Bye Bye!")
			return

		}
	}
}

func main() {
	//	NoStopGoroutine()
	//	UseTimeGoroutine()
	//	UseChannelGoroutine()
	//  UseChannelBufferGoroutine()
		ForLoopByGorouine()
	//	NonBlockingGoroutine()
	//	BlockingGoroutine()
	//	ChanTypeSample()
	//	SelectSample()
	//SelectWithStopSample()
}
