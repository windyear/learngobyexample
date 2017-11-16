package main

import (
	"sort"
	"fmt"
	"os"
)

type ByLength []string
//自定义排序程序
func (s ByLength) Len() int{
	return len(s)
}

func (s ByLength) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool{
	return len(s[i]) < len(s[j])
}

func testdefinesort(){
	fruit := []string{"peach", "apple", "watermelon", "tomato", "purple"}
	sort.Sort(ByLength(fruit))
	fmt.Println("fruit: ", fruit)
}
//测试自带排序程序
func testSort(){
	var str = []string{"windyear", "dengli", "play"}
	sort.Strings(str)
	fmt.Println("strings: ", str)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

func testPanic(){
	panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil{
		panic(err)
	}
}

func createFile(p string) *os.File{
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writring")
	fmt.Fprintln(f, "data: haha for test")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}

func testdefer(){
	f := createFile("./test.txt")
	defer closeFile(f)
	writeFile(f)
}
func gobyexamplechapter12(){
    //testSort()
    //testdefinesort()
    //testPanic()
    testdefer()
}
