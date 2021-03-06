//1 interface: 抽象类型，接口类型,仅包含方法集合

//2 接口内嵌，类继承/方法继承

//3 接口实现: 某一类型定义所有接口的方法

//4 指针方法与值方法可以互相调用(隐式转换)，
//值需为可取址的值(临时变量不可以)

//5 接口值: 具体类型(接口动态类型)&&
//类型值(接口动态值)
//	5.1 接口零值: 类型与值均为nil
//	赋值具体类型后，接口类型隐式转换为具体
//	类型且动态值为该类型的copy

//6 接口可比较(取决于动态类型)，可用于map的key或者switch操作数
//见panic/recover,若动态类型不可比较(slice)则panic

//7 接口值为nil与包含nil指针的接口不同

/*8 类型断言:x.(T)
8.1 T为具体类型：检查x的动态类型==T,OK返回T，失败panic
8.2 T为接口类型，检查x的动态类型是否满足T方法，不会获得动态
类型，结果为相同类型的值和部分的接口值，结果有类型T,保护接口值
内部的动态类型和值的部分
var w io.Writer
w = os.Stdout
rw,ok := w.(io.ReadWriter) // success: *os.File has both Read and Write
w = new(ByteCounter)
rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
*/
/*9 switch case
switch x.(type){
case nil:
case int:
case bool:
case float32:
case float64:
case string:
}
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//4
type m4 struct {
	str1 string
	str2 string
}

//指针方法
func (t *m4) mPrintln() {
	fmt.Println("1st str: ",
		t.str1, "\n2st str: ",
		t.str2)
}

//值方法
func (t m4) mVPrintln() {
	fmt.Println("1st str: ",
		t.str1, "\n2st str: ",
		t.str2)
}
func testM4() {
	//4 test
	//临时变量取址，指针调用
	m := m4{
		str1: "str1 of m",
		str2: "str2 of m",
	}
	//临时变量指针调用指针方法
	(&m4{"Pointer-pointer method hi",
		"Pointer-pointer method hello"}).mPrintln()
	//临时变量调用指针方法error
	/*
		(m4{"temp value-pointer method hi",
			"temp value -Pointer method hello"}).mPrintln()
	*/
	//变量调用指针方法
	m.mPrintln()
	//指针调用变量方法
	(&m).mPrintln()
	//临时变量值调用
	m4{"temp value-value method hi",
		"temp value-valuee method hello"}.mVPrintln()
}

//7 8
const debug = false

func test7() {
	var buf *bytes.Buffer
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic: ", p)
		}
	}()
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	//隐式转换io.Writer接口复制buf
	f(buf)
	if debug {
		// ...use buf...
		fmt.Println("here debug")
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	fmt.Printf("the type of out in func f: %T \n",
		out)
	fmt.Println("the value of out in func f:",
		out)
	//8 断言成功返回接口类型T,否则panic
	f, _ := out.(*bytes.Buffer), out.(*os.File)
	defer func() {
		fmt.Println("type assertion: f", f)
	}()

	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
func main() {
	//testM4()
	test7()

}

/*****old************
package main
import (
	"fmt"
	"strconv"
)
//type int1 int
type int1 int
type jie1 struct {
	a1 string
	b1 string
}
//type float float64
type try1 interface{
	str()string
}
func (a int1)str()string{
	return strconv.Itoa(int(a))
	//fmt.Println(a)
}
func (a jie1)str()string{
	var a1 string = ""
	a1 =a1+ a.a1+";"
	a1 +=a.b1+";"
	return a1
	//fmt.Println(a.a1,"2",a.b1)
}
func printTry(a try1){
	fmt.Println(a.str())
}
func main(){
	//var b float64 = 3.2
	var a int1
	a=44
	b := jie1{"1","123"}
	/**********method 1
	var c try1
	c=a
	fmt.Println(c.str())
	c=b
	fmt.Println(c.str())
	*****end of method 1***/
//printTry形参为interface
/****method2*****
	printTry(a)
	printTry(b)
}
*/
