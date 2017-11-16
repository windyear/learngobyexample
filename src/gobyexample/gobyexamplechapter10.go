package main

import (
	"time"
	"fmt"
	"sync/atomic"
	"runtime"
)

//定时器就是一段时间之后再执行任务，可以终止
func testtimer(){
	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second * 4)
	go func(){
		<- timer2.C
		fmt.Println("Timer2 expired")
	}()
	//关闭计时器
	stoptimer2 := timer2.Stop()
	if stoptimer2{
		fmt.Println("stop timer2")
	}
}
//打点器就是经过一段时间返回一个通道时间值
func testticker(){
	ticker := time.NewTicker(time.Second)
	go func(){
		for t := range ticker.C{
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Second * 3)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

//测试一个工作池，函数从参数通道中提取任务，然后返回结果给通道，同时启动多个函数并行运行，任务会被并行分配执行
func worker(jobs <-chan int , id int, result chan<- int){
	for j := range jobs{
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func testworker(){
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启动三个阻塞线程
	for w := 1; w <= 3; w++{
		go worker(jobs, w,results)
	}

	//通过通道传入工作任务，其他并行的协程启动执行函数
	for j := 1; j <= 9; j++{
		jobs <- j
	}
	close(jobs)
	//不接收返回值函数会阻塞？？
    for i := 0; i < 9; i++{
    	<-results
	}
}
//通过定时器和通道可以进行速率限制
func testspeed(){
	//请求通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++{
		requests <- i
	}
	close(requests)
	//用打点器作为一个定时器
	limiter := time.Tick(time.Second)
	for req := range requests{
		//定时器
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++{
		burstyLimiter <- time.Now()
	}
	//每秒产生一个脉冲
	go func(){
		for t := range time.Tick(time.Second){
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 15)
	for i := 0; i < 15; i++{
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests{
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

//原子操作，操作的原子性
func testAtomic(){
	var ops uint64 = 0
	//启动50个协程进行操作
	//用for循环一直操作
	for i := 0; i < 5; i++{
		go func(){
			for{
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
			//允许其他协程运行
			//
		}()
	}

	time.Sleep(time.Second * 3)

	//取出值
	opsfinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsfinal)
}
func gobyexamplechapter10(){
	//testtimer()
	//testticker()
	//testworker()
	//testspeed()
	testAtomic()
}