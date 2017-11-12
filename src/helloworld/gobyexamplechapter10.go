package main

import (
	"time"
	"fmt"
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
func gobyexamplechapter10(){
	testtimer()
	testticker()
}