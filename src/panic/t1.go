/*1 数组访问越界、空指针引起panic异常
2 panic异常，程序中断运行，并立即执行goroutine
	中defer函数，随后程序崩溃输出日志信息: panic
	value、函数调用堆栈信息；
	panic value通常为某种错误信息，日志信息提供足够诊断工具；
	panic异常与日志信息一并记录至报告；
3 输入值为空接口:
	func panic(v interface{})
4 recover输出值为空接口:
	func recover() interface{}
5 panic及recover英文解释
	5.1 When panic is called, it immediately stops execution of the
	//panic异常，程序回滚goroutine的栈，执行栈中defer函数
	current function and begins unwinding the stack of the goroutine,
	running any deferred functions along the way.
	//若回滚至栈顶，则程序(goroutine)死掉
	If that unwinding reaches the top of the goroutine's stack,
	the program dies.
	//recover函数可以重新控制goroutine,并重回正常执行顺序
	5.2 However, it is possible to use the built-in function recover
	to regain control of the goroutine and resume normal execution.
	//recover只能放置defer函数中，因为panic会回滚defer函数
	A call to recover stops the unwinding and returns the
	argument passed to panic. Because the only code that runs
	while unwinding is inside deferred functions,
	recover is only useful inside deferred functions.
	5.3 recover可用于关闭失败的goroutine而不影响其他goroutine
	One application of recover is to shut down a failing goroutine
	inside a server without killing the other executing goroutines.
*/
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	//defer函数后panic不会输出栈信息
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("panic value: ", p)
		}
	}()
	//添加打印栈信息函数可输出栈信息
	defer printStack()
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
