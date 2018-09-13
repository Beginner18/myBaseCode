//close(nil)
//关闭已关闭通道引起panic
//持续从通道接收数据可用for range语句
//通道关闭for range 语句退出循环
package main

import (
	"fmt"
)

func closeChan(c1 chan interface{}) {
	//c1<-interface{}{}
	close(c1)
	//close(c1)
}
func receiveChan(c1 <-chan interface{}, control chan interface{}) {
	res, ok := <-c1
	fmt.Println("the value of res:", res, "the value of OK:", ok)
	//control<-interface{}{}
	close(control)
	//不可关闭接收通道
	//close(c1)
}
func main() {
	c1 := make(chan interface{}, 0)
	control := make(chan interface{}, 0)
	go closeChan(c1)
	go receiveChan(c1, control)
	<-control
}
