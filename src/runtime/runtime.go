package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	testFunc()
}

//测试runtime包中函数功能
func testFunc() {
	//defer 匿名函数，函数返回值为匿名函数
	//调用该函数需要两个()
	defer func() func() {
		start := time.Now()
		//sleep
		//time.Second为type Duration const类型
		time.Sleep(2 * time.Second)
		//返回值为匿名函数
		return func() {
			fmt.Println("start time: ", start, "cost time: ", time.Since(start))
		}
	}()()

	//返回GOROOT path
	fmt.Println("GOROOT: ", runtime.GOROOT())
	//返回go version
	fmt.Println("GO version: ", runtime.Version())
	//返回本机CPU核数
	fmt.Println("CPU num: ", runtime.NumCPU())
	//修改程序最大核数为输入参数，并返回之前核数
	//输入参数小于1则保持原数目
	fmt.Println("max procs: ", runtime.GOMAXPROCS(2))
	var mem runtime.MemStats
	//读取内存分布及统计信息,mem中包括系统内存分配及堆栈信息
	runtime.ReadMemStats(&mem)
	fmt.Println("内存统计信息",
		"\n已申请且在使用的字节数: ", mem.Alloc,
		"\n已申请堆字节数: ", mem.HeapAlloc,
		"\n内存分配次数: ", mem.Mallocs)
}
