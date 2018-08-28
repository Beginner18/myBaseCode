/************************************
1 go语言方法
2 两个不同结构的可以实现相同的方法
***********************************/
package main
import (
	"fmt"
)
type try1 struct{
	str string
	num int
}
type try2 struct{
	str string
	str1 string
}
func (a try1)pri(){
	fmt.Println(a.str)
	fmt.Println(a.num)
}
func (a try2)pri(){
	fmt.Println(a.str)
	fmt.Println(a.str1)
}
func main(){
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
}

