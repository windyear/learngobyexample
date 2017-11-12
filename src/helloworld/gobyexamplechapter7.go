package main

import (
	"fmt"
	"errors"
)


//可以自己定义错误类型的结构
type argError struct{
	arg int
	prob string
}
//实现Error方法才可以实现自定义错误类型
func (argerror *argError) Error()string{
	return fmt.Sprintf("%d - %s", argerror.arg,argerror.prob)
}

func canwork2(num int)(int, error){
	if num == 10{
		return -1, &argError{num, "can not work with it"}
	}
	return num, nil
}

func canwork(num int)(int, error){
	if num == 10{
		return -1, errors.New("can't work with number 10")
	}
	return num, nil
}

//用于协程的函数
func testgo(from string){
	for i := 0; i < 16; i++{
		fmt.Println(from, ":", i)
	}
}
func gobyexamplechapter7(){
	testarray := []int{1,10}
	//用一个循环去遍历测试函数
	for _, v := range testarray{
		if r, e := canwork(v); e != nil{
			fmt.Println("canwork fail", e)
		}else{
			fmt.Println("canwork work: ", r)
		}
	}

	for _, v := range testarray{
		if r, e := canwork2(v); e != nil{
			fmt.Println("canwork fail", e)
		}else{
			fmt.Println("canwork work: ", r)
		}
	}
    //测试协程
    //因为是同一个函数所以会自动阻塞？
    testgo("main")
    go testgo("another")
    go func(msg string){
    	for i := 0; i < 15; i++ {
			fmt.Println(msg)
		}
	}("going")

    var input string
    fmt.Scanln(&input)
    fmt.Println("done", input)
}