package main

import (
	"fmt"
	"time"
)

func testchan(mes chan string){
	mes <- "test chan in func"
	mes <- "second test chan in func "
}

//利用通道进行函数同步运行，发送一个执行完了的信号
func testsyc(done chan bool){
	fmt.Println("the function testsyc is working......")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

//指定通道的方向,pings通道只能接受数据
func ping(pings chan<- string, mes string){
	pings <- mes
}

func pong(pongs <-chan string, pings chan<- string){
	msg := <- pongs
	pings <- msg
}

//通道选择器
func testselect(){
	c1 := make(chan string)
	c2 := make(chan string)
	go func(){
		time.Sleep(time.Second * 4)
		c1 <- "four seconds"
	}()

	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "two seconds"
	}()
	for i := 0; i < 2; i++{
		select{
		case msg1 := <-c1:
			fmt.Println("received after ", msg1)
			case msg2 := <-c2:
				fmt.Println("received after ", msg2)
		}
	}
}
func gobyexamplechapter8(){
	//测试通道
	message := make(chan string)
	message2 := make(chan string, 2)
	//直接调用匿名函数
	go func(){
		message <- "testchan"
	}()
	mes := <- message
	fmt.Println("test chan:", mes)

	go testchan(message2)

	fmt.Println(<-message2)
	fmt.Println(<-message2)
	done := make(chan bool)
	go testsyc(done)

	//一直阻塞到函数执行完，如果没有接受通道的信号，可能协程还没开始整个程序就结束了
	<- done

	//测试通道的方向性
	//如果不加上缓冲机制便会造成死锁，因为要等接受准备好了才可以发送，而没有用其他协程
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, "passed messsage")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
	testselect()
}