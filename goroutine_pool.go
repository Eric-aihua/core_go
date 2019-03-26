package main

import (
	"fmt"
	"strconv"
	"time"
)

//定义任务Task类型,每一个任务Task都可以抽象成一个函数
type Task struct {
	f func() error
}

//通过NewTask来创建一个Task
func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

//执行Task任务的方法
func (t *Task) Execute() {
	t.f() //调用任务所绑定的函数
}

//定义池类型
type Pool struct {
	//对外接收Task的入口
	EntryChannel chan *Task

	//协程池最大worker数量,限定Goroutine的个数
	worker_num int

	//协程池内部的任务就绪队列
	JobsChannel chan *Task
}

//创建一个协程池
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		worker_num:   cap,
		JobsChannel:  make(chan *Task),
	}

	return &p
}

//协程池创建一个worker并且开始工作
func (p *Pool) worker(work_id int) {
	for task := range p.JobsChannel {
		task.Execute()
		//fmt.Println("worker ID ", work_id, " 执行完毕任务")
	}
}

func (p *Pool) run() {
	//  每一个Worker用一个Goroutine承载
	for i := 0; i < p.worker_num; i++ {
		go p.worker(i);
	}
	//2, 从EntryChannel协程池入口取外界传递过来的任务,并且将任务送进JobsChannel中
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
	//3, 执行完毕需要关闭JobsChannel
	close(p.JobsChannel)
	//4, 执行完毕需要关闭EntryChannel
	close(p.EntryChannel)
}

func dataProcess(ch chan string) {
	param := <-ch
	result := "process data " + param
	time.Sleep(2 * time.Second)
	fmt.Println(result)
	//ch <- result
}

func demo() {
	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//创建一个Task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()
	//启动协程池p
	p.run()
}

func fullDemo() {
	//创建一个协程池,最大开启3个协程worker
	p := NewPool(10)
	ch := make(chan string, 1)
	t := NewTask(func() error {
		dataProcess(ch)
		return nil
	})

	//创建一个Task
	go func() {
		for i := 1; i < 30; i++ {
			ch <- strconv.Itoa(i)
			p.EntryChannel <- t
		}
		return
	}()
	//启动协程池p
	p.run()

}

//主函数
func main() {
	//demo()
	fullDemo()
}
