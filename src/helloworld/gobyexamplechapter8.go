package main

func gobyexamplechapter8(){
	//测试通道
	message := make(chan string)
	//直接调用匿名函数
	go func(){
		message <- "testchan"
	}()
	
}