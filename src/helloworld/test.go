package main

import ("fmt"
	    "time"
	    "math"
)
//变量　常量
func gobyexamplechapter1(){
	//go by example: value
	fmt.Println("Hello" + "Deng Li!")

	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("3.0/2.3", 3.0 / 2.3)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//声明变量的不同方式
	var teststring string = "I'm Zhangxuwen!"
	var test1, test2 int = 1, 2
	var d = true
	f := "short"
	fmt.Println(teststring)
	fmt.Println(test1,test2)
	fmt.Println(d,f)

	//常量，出现var的地方都可以使用const
	const const_string string = "constant"
	const n = 500000
	const testn = 3e20/n
	fmt.Println(const_string)
	fmt.Println(int64(testn))
	fmt.Println(math.Sin(n))



}
//for循环　if/else分支
func gobyexamplechapter2(){
	//有三种不同的for循环的方式，省略第一以及第三个参数的时候变成while循环
	i := 1
	for j := 2; j <= 9; j++{
		fmt.Println(i)
	}
	//可以理解为就是while
	for i <= 99{
		fmt.Println(i)
		i = i + 1
	}

	for{
		fmt.Println("loop")
		break
	}

	//if-else分支，没有什么特别的用法，只有一点：就是在条件语句之前声明的变量可以再所有的条件分支中使用
	if 11 % 2 == 0{
		fmt.Println("11 is even")
	}else{
		fmt.Println("11 is odd")
	}
	if num := 12; num < 0{
		fmt.Println(num, "is negative")
	}else if num < 10{
		fmt.Println(num, "has 1 digit")
	}else{
		fmt.Println(num, "has multiple digits")
	}

	//switch语句
	i = 1
	fmt.Println("write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//time包中可以获取当前是星期几
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

}
//数组　切片
func gobyexamplechapter3(){
    var testarray [5]int
    fmt.Println("emp: ", testarray)
    testarray[4] = 100
    fmt.Println("set: ", testarray)
    fmt.Println("get: ", testarray[4])

    b := [5]int{1, 2, 4, 5, 5}
    fmt.Println(b)

    //二维数组
    var c [2][3]int
    fmt.Println("2d", c)

    //slice支持切片操作，仅仅由元素决定
    //编程时这里出现了一个错误，没有采用自动推导类型的:=而写成了＝号
    testslice := make([]string, 3)
    fmt.Println("test slice: ", testslice)
    //切片操作，中间是冒号，注意是半开半闭区间
    testslice2 := testslice[0:3]
    fmt.Println("test slice 2: ", testslice2)
    //slice组成的多维数据结构的内部长度可以不容，比如二维数组的列数可以不同
    twoarray := make([][]int, 3)
    for i := 0; i < 3; i++{
    	innerlen := i + 1
    	//这里应该初始化一个slice
    	twoarray[i] = make([]int, innerlen)
    	for j := 0; j < innerlen; j++{
    		twoarray[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoarray)
    //map是哈希字典
    testmap := make(map[string]int)
    testmap["windyear"] = 1
    testmap["DengLi"] = 2
    testmap["love"] = 3
    fmt.Println("print map: ", testmap)
    testvalueofmap := testmap["DengLi"]
    fmt.Println(testvalueofmap)
    fmt.Println(len(testmap))
    delete(testmap, "DengLi")
    _, prs := testmap["windyear"]
    fmt.Println("return :", prs)
    initmap := map[string]int{"who":1, "are": 2, "you": 3}
    fmt.Println("initmap: ", initmap)
    //range可以迭代各种数据结构,range返回的两个值分别是索引和值
    testnums := []int{2, 3, 4}
    sum := 0
    for _, num := range testnums{
    	sum = num + sum
	}
	fmt.Println("the sum is : ", sum)
	//注意格式化输出使用的是printf函数而不是println函数
	for k, v := range testmap{
		fmt.Printf("%v -> %v\n", k, v)
	}

}

//gobyexamplechapter4
func plus(a int, b int) int {
	return a + b
}

func vals()(int, int){
	return 10, 11
}
func argsvals(nums ...int){
	fmt.Print(nums," ")
	total := 0
	for _, num := range nums{
		total = total + num
	}
	fmt.Println("total: ", total)
}
func gobyexamplechapter4(){
	sum := plus(123, 456)
	fmt.Println("sum: ",sum)

	val1, val2 := vals()
	fmt.Println("return valus", val1, val2)

	argsvals(1,2,3)
	argsvals(4,5,6)
	nums := []int{5, 6, 7, 8,9}
	//要逐个传入数组参数，需要在数组后面加上省略号
	argsvals(nums...)

}
func main(){

	fmt.Println("Hello world! First program in goland")
	test_time := time.Date(2017, time.November, 11, 23,3,5,0, time.UTC )
	fmt.Printf("Go launched at %s\n", test_time.Local())
    //gobyexamplechapter1()
    //gobyexamplechapter2()
    //gobyexamplechapter3()
    //gobyexamplechapter4()
    //gobyexamplechapter5()
    //gobyexamplechapter6()
    //gobyexamplechapter7()
    //gobyexamplechapter8()
    //gobyexamplechapter9()
    //gobyexamplechapter10()
    gobyexamplechapter11()
}
