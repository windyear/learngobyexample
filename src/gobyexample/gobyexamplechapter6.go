package main

import (
	"fmt"
	"math"
)

type person struct{
	name string
	age int
}

type rect struct{
	height int
	width int
}
func zeroval(val int){
	val = 0
}

//传入指针参数
func zeroptr(val *int){
	*val = 0
}

//结构体
func teststruct(){
	//声明变量的前面要加上var
	var persontest person
	persontest.age = 21
	persontest.name = "DengLi"
	fmt.Println(persontest)
	person1 := person{"windyear", 23}
	fmt.Println(person1)
	fmt.Println(person1.name,person1.age)

}

//结构中的方法可以理解为方法是结构的一个成员，在声明函数的时候加上一个接收器，接收器也分为指针和值类型
//go会自动处理调用过程中的指针和值的关系，没有搞懂？？？
func (r rect) area() int{
	return r.height * r.width
}

func (r *rect) area2() int{
	return r.height * r.width
}

//用接口去定义一组方法
type square struct{
	width, height float64
}

type circle struct{
	radius float64
}

type geometry interface{
	area() float64
	perim() float64
}

//两个结构分别去实现方法
func (s square) area() float64{
	return s.width * s.height
}

func (s square) perim() float64{
	return 2 * s.width + 2 * s.height
}

func (c circle) area() float64{
	return 2 * math.Pi * c.radius * c.radius
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

//测试接口的函数
func measure(g geometry){
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}
func gobyexamplechapter6(){
	test := 1
	zeroval(test)
	fmt.Println("zeroval: ", test)
	zeroptr(&test)
	fmt.Println("zeroptr: ", test)
	//打印指针
	fmt.Println("print ptr: ", &test)
	teststruct()

	testrect := rect{12,14}
	fmt.Println("area: ", testrect.area())
	fmt.Println("*rect area: ", testrect.area2())

	//测试接口的函数,可以传入实现了接口的实例
	testsquare := square{2,3}
	testcircle := circle{5}
	measure(testsquare)
	measure(testcircle)
}