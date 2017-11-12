package main

import "fmt"

//在同一个目录中只能有一个包名，所有文件的内容跟在同一个文件中没有差别
func intseq() func() int{
	i := 0
	return func()int{
		i += 1
		return i
	}
}

func fact(n int) int{
	if n == 0{
		return 1
}
    return n * fact(n-1)
}

func gobyexamplechapter5(){
	nextInt := intseq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intseq()
	fmt.Println(newInts())
	fmt.Println(fact(10))
}