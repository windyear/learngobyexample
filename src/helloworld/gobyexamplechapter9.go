package main

import (
	"time"
	"fmt"
)

//超时应用于需要等待的IO操作或者连接外部资源等等
//select可以实现多路复用
func testtimeout(){
	c1 := make(chan string)
	go func(){
		time.Sleep(time.Second * 4)
		c1 <- "after 2 second"
	}()
	//select只选取一个操作
	select{
	    case result := <- c1:
		fmt.Print(result)
		case <- time.After(time.Second * 3):
			fmt.Println("timeout 3")
	}
}
//利用default可以实现非阻塞的多路复用，如果没有情况已经发生，则调用默认的情况
func testselectdefault(){
	msg := make(chan string)
	select{
	case message := <- msg:
		fmt.Println("send message: ", message)
	default:
		fmt.Println("no message received")
	}
	//一个非阻塞的发送方法
	message := "windyear have sent message."
	select{
	case msg <- message:
		fmt.Println(<-msg)
	default:
		fmt.Println("no message have sent")
	}

}

//用range可以遍历通道的值
func rangechan(){
	msg := make(chan string, 2)
	msg <- "one"
	msg <- "two"
	close(msg)
	//如果不关闭通道则会一直阻塞
	for message := range msg{
		fmt.Println("receieve message: ", message)
	}

}

func closechan(){
	mes := make(chan int, 5)
	done := make(chan string)
	//一个负责接受通道数据的匿名函数
	go func(){
		for{
			if j, more := <- mes;more{
				fmt.Println("jobs receieve", j)
			}else{
				fmt.Println("all jobs have been receive")
				done <- "receive all jobs"
				return
			}
		}
	}()

	//发送数据到通道
	for i := 0; i < 4; i++{
		mes <- i
	}
	//发送完毕关闭通道
	close(mes)

	<- done
}
func gobyexamplechapter9(){
	testtimeout()
	testselectdefault()
	rangechan()
	closechan()
}