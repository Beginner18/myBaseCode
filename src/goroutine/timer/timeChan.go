//定时器与channel复用,超时控制
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func test1() {
	timer1 := time.NewTimer(2 * time.Second)
	<-time.After(3 * time.Second)
	fmt.Println("present time: ", time.Now())
	fmt.Println("2s: ", <-timer1.C)
	fmt.Println("stop timer: ", timer1.Stop())
}
func timeOut(c1 <-chan int) {
	var time1 *time.Timer
	outTime := 500 * time.Millisecond
	for {
		if time1 == nil {
			time1 = time.NewTimer(outTime)
		} else {
			time1.Reset(outTime)
		}
		select {
		case value, ok := <-c1:
			if ok {
				fmt.Println("the value is:", value)
			} else {
				fmt.Println("channel closed")
				os.Exit(0)
			}
		case <-time1.C:
			fmt.Println("timeOut")
		}
	}
}

//制作心跳包，每隔一段时间执行ls
func timeTick() {
	tick := time.NewTicker(time.Second)
	//for range 可持续接收通道值，只要通道不关闭
	for time := range tick.C {
		fmt.Println("now time", time)
		cmd := exec.Command("/bin/bash", "-c", "ls", "")
		str, _ := cmd.Output()
		fmt.Println("the res is: ", string(str))
	}
}
func main() {
	//test1()
	/*test2
	c1 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			if i == 3 {
				time.Sleep(time.Second)
			}
			c1 <- i
		}
		close(c1)
	}()
	timeOut(c1)
	*/
	timeTick()

}
