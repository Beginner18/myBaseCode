/*1 适用于函数和方法,只能跟方法或者函数
函数中嵌套defer改变defer执行顺序，函数中defer按栈执行，该函数正常执行
*/
/*2 defer语句延迟执行直到函数return或者panic
return函数不是原子级，包括写入返回参数及
函数返回，defer语句在写入返回参数后函数返
回前执行
*/
/*3 defer语句反顺序执行一般用于成对
打开、关闭、连接、断开连接、加锁、释放锁
释放资源
*/
/*4 defer语句所在处变量入栈，只是函数结束再处理
	该语句，若值传递则后续值改变不影响defer语句，
	若地址传递则影响defer语句
	defer函数参数在defer语句执行时初始化
	defer函数调用时具体执行
4.1 defer记录函数的进入和退出，trace函数
*/
/*5 方法为一种数据类型，方法定义后，通过赋值给具有方法的
数据类型，则方法的类型与赋值类型一致，可以执行方法操作
*/
/*6 Each time a "defer" statement executes, the function value
and parameters to the call are evaluated as usualand saved anew
but the actual function is not invoked.
anew: 重新再
*/
/*7 defer函数调用函数返回值与函数输入参数按正常函数处理(复制新参数)，
闭包函数变量更新defer函数中变量更新，但实际函数没有激活。
*/
/*8 函数返回的过程是这样子的：先给返回值赋值，然后调用defer表达式，
最后才是返回到调用函数中。利用闭包defer函数(指针)可修改返回值的值.
*/
/*9 defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，
修改返回值，使最终的函数返回值与你想象的不一致
*/
/*10 defer+panic
无panic的函数按defer正常执行: defer定义，func　return/panic执行
defer
panic 的函数仅执行panic之后语句不执行
*/
//
package main

import (
	"fmt"
	"log"
	"time"
)

//定义方法
type defMeth struct {
	err bool
	str string
}

//定义接口，接口为一种数据类型，提供操作
type in interface {
	ini()
	prinOut()
}

func (def *defMeth) ini() {
	def.err = true
	def.str = "hei"
}
func (def *defMeth) prinOut() {
	fmt.Println("just print: ", def)
}

type defError struct {
	err error
	str string
}

func (def *defError) ini() {
	def.err = fmt.Errorf("this is an err: %s\n", "true")
	def.str = "this is a bug"
}
func (def *defError) prinOut() {
	fmt.Println("just print: ", def)
}
func main() {
	var s1 defMeth
	var s2 defError
	var def in
	a := 1
	//defer+闭包处理, a为主函数变量，a值变化该defer函数变量会更新
	defer func() { fmt.Println("close package, test the value of a:", a) }()
	defer fmt.Println("fmt a is: ", a)
	//值传递,defer函数入栈，变量值已赋值，只是最后执行
	defer func(a int) { fmt.Println("no name func a is: ", a) }(a)
	//地址传递：传递的为地址，若地址处值改变则操作值改变，map传递的地址
	defer func(a *int) { fmt.Println("no name func a address is: ", *a) }(&a)
	//为接口赋值变量类型及变量地址
	def = &s1
	//a:=1
	def.ini()
	//defer a=2 //错误defer后必须是函数或者方法或者接口
	def.prinOut()
	//fmt.Println("result", bigSlowOperation())
	//接口
	def = &s2
	def.ini()
	//函数中嵌套defer改变defer执行顺序，函数中defer按栈执行，该函数正常执行
	def.prinOut()
	a++
	fmt.Println(a)
	fmt.Println("the value of bigSlow func: ", bigSlowOperation())
	//10 test defer panic
	deferPanic()

}

func bigSlowOperation() int {
	log.Printf("main starts %s", "\n")
	i := 1
	//非匿名函数调用参数i已确定:1
	//defer fmt.Println(i)
	//匿名函数可调用返回值:2
	defer func() { fmt.Println(i) }()
	defer trace("bigSlowOperation")() // don't forget the
	//extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation
	//by sleeping
	i++
	return i
}

//返回值为函数，函数需要正常入栈，
//函数输入参数及返回地址按正常函数处理,
//defer trace("")()
//函数为trace中包装的匿名函数，需正常处理该函数,
//的形参及返回地址
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//10 defer+panic: what happends after a defer func panic
func deferPanic() {
	//正常执行
	defer fmt.Println("the last defer panic func")
	defer fmt.Println("the 3st defer panic func")
	//panic　func 仅执行panic语句
	defer func() { panic("test defer panic"); fmt.Println("the func panics") }()
	//正常执行
	defer fmt.Println("the 2st defer panic func")
	defer fmt.Println("the 1st defer panic func")
}
