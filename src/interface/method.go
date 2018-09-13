/************************************
1 go语言方法
2 两个不同结构的可以实现相同的方法
3 嵌入结构体类型扩展/继承
3.1 继承结构体成员可直接调用被继承结构体成员
结构体,及其方法;
3.2 指针结构体嵌套同结构体嵌套
3.3 继承后，结构体has被继承结构体但并不是同一
类型，函数调用
//4 方法的值
T.pri	//为一函数
***********************************/
package main

import (
	"fmt"
)

type try1 struct {
	str string
	num int
}
type try2 struct {
	str  string
	str1 string
}

func (a try1) pri() {
	fmt.Println(a.str)
	fmt.Println(a.num)
}
func (a try2) pri() {
	fmt.Println(a.str)
	fmt.Println(a.str1)
}

//3 嵌套结构体
type try3 struct {
	try1
	check bool
}

func test3() {
	a := try3{try1{"嵌套结构体", 3}, true}
	a.pri()
	fmt.Println("嵌套结构体，str:", a.str, "num:", a.num)
}

//3.2 指针嵌套结构体体
type try31 struct {
	*try1
	check bool
}

func test31() {
	a := try31{&try1{"**嵌套结构体**", 3}, false}
	a.pri()
	fmt.Println("嵌套结构体，str:", a.str)
}

//3.3 嵌套结构体与被嵌套结构体为不同类型
func test33(a *try1) {
	a.str = "modified"
	a.num = 01
}
func mainTest33() {
	a := try31{&try1{"**嵌套结构体**", 3}, false}
	fmt.Println("before:")
	a.pri()
	//error: 类型不匹配
	//test33(a)
	test33(a.try1)
	fmt.Println("after:")
	a.pri()
}

func main() {
	var a try1 = try1{
		str: "123",
		num: 1,
	}
	a.pri()
	var b try1 = try1{
		str: "123",
		num: 1,
	}
	b.pri()
	test3()
	test31()
	mainTest33()
}
