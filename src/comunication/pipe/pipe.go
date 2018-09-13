//os/exec包实现linux系统下管道
//底层实现为linux系统管道
//管道是单向的
//命名管道可多路复用
//go io.Pipe提供基于内存的有原子性操作的管道
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

//1 匿名管道:ps aux | grep go
//1.1 Run 方法
//2 named管道:任何进程均可以通过命名管道交换数据
//2.1 write未完成时，reade进程被堵塞
func pipe1Run() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "go")
	//定义cmd1的输出为cmd2的输入，即同时满足读写
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer
	cmd1.Stdout = &buf1
	cmd1.Run()
	fmt.Println("here2")
	cmd2.Stdin = &buf1
	cmd2.Stdout = &buf2
	cmd2.Run()
	fmt.Println("final res is : ", buf2.String())
}

//1.2 start wait方法
func pipe1StartWait() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "go")
	//定义cmd1的输出为cmd2的输入，即同时满足读写
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer
	cmd1.Stdout = &buf1
	if err := cmd1.Start(); err != nil {
		fmt.Println("Error, cmd1 start:", err)
		return
	}
	//堵塞直到cmd1 start结束
	if err := cmd1.Wait(); err != nil {
		fmt.Println("Error, cmd1 wait:", err)
		return
	}
	cmd2.Stdin = &buf1
	cmd2.Stdout = &buf2
	if err := cmd2.Start(); err != nil {
		fmt.Println("Error, cmd2 start:", err)
		return
	}
	//堵塞直到cmd2 start结束
	if err := cmd2.Wait(); err != nil {
		fmt.Println("Error, cmd2 wait:", err)
		return
	}
	fmt.Println("final res is : ", buf2.String())
}

//linux command test, OK
func test1() {
	cmd1 := exec.Command("ls")
	res, err := cmd1.Output()
	if err != nil {
		fmt.Println("cmd1, run error:", err)
	}
	fmt.Println("res: ", string(res))
}

//2 os包命名管道
func pipe2() {
	//pipe管道返回读写*File及错误
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Println("pipe err: ", err)
	}
	//reader 与writer并发
	//reader放置某个子协程，writer放置主携程
	//reader 与writer分别放置两个不同的子携程输出为空
	go func() {
		output := make([]byte, 100)
		reader.Read(output)
		fmt.Println("pipe2 output res: ", string(output))
	}()
	_, err1 := writer.Write([]byte("123445"))
	if err1 != nil {
		fmt.Println("pipe err: ", err1)
	}
	time.Sleep(1 * time.Second)
}

//3 io包管道基于内存原子级操作
//all goroutine is asleep
func pipe3() {
	//pipe管道返回读写*File及错误
	reader, writer := io.Pipe()
	//reader 与writer并发
	go func() {
		output := make([]byte, 100)
		reader.Read(output)
		fmt.Println("pipe3 output res: ", string(output))
	}()
	_, err1 := writer.Write([]byte("123445"))
	if err1 != nil {
		fmt.Println("pipe err concurrent1: ", err1)
	}
	time.Sleep(100 * time.Millisecond)
}
func pipe31() {
	//pipe管道返回读写*File及错误
	reader, writer := io.Pipe()
	//reader 与writer并发
	//原子级操作，一次全部写成功，不会单个写
	go func() {
		_, err1 := writer.Write([]byte("123445"))
		if err1 != nil {
			fmt.Println("pipe err concurrent1: ", err1)
		}
	}()

	go func() {
		var err1 error
		input := make([]byte, 26)
		for i := 65; i <= 90; i++ {
			input[i-65] = byte(i)
		}
		fmt.Println("input: ", string(input))
		_, err1 = writer.Write(input)
		if err1 != nil {
			fmt.Println("pipe err concurrent1: ", err1)
		}
	}()
	output := make([]byte, 40)
	reader.Read(output)
	fmt.Println("pipe3 output res: ", string(output))
	time.Sleep(1 * time.Second)
}
func main() {
	//pipe1Run()
	//pipe1StartWait()
	//test1()
	//pipe2()
	//pipe3()
	pipe31()
}
