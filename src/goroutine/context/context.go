/*
https://segmentfault.com/a/1190000006744213
1　简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等
相关操作
2 控制并发
3 关闭不需要的goroutine
3.1 全局变量法，对应goroutine检测到全局变量对应关闭值则关闭goroutine
3.2 chan+select方法
3.3 go出现关系链，root　goroutine建立子节点，子节点建立其子节点，树结构
context跟踪goroutine
ctx取消，以ctx为参数的goroutine对应ctx.Done channel被关闭,以此接收到关闭
消息

*/
package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

//3.3 对应案例
func eg33() {
	//基于根建立上下文
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//context控制go
	go watch(ctx, "1st")
	go watch(ctx, "2st")
	go watch(ctx, "3st")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	//取消监控
	cancel()
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

//父节点取消则所有子节点取消,ctx.Done() channel接收到信息值
func myContext() {
	var wg sync.WaitGroup
	//ctx, cancel1st := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//ctx, _ := context.WithCancel(context.Background())
	//cancel1st()
	fmt.Println("here my")
	wg.Add(1)
	go child(ctx, wg)
	wg.Wait()
}
func child(ctx context.Context, wg sync.WaitGroup) {
	fmt.Println("here child")
	select {
	case <-ctx.Done():
		fmt.Println("root exit error: ", ctx.Err())
		wg.Done()
		os.Exit(0)
	default:
		fmt.Println("hello handle 1st layer")
		//ctx2st, cancle2st := context.WithCancel(ctx)
		ctx2st, _ := context.WithCancel(ctx)
		go leaf(ctx2st, wg)
		wg.Done()
	}
}
func leaf(ctx context.Context, wg sync.WaitGroup) {
	fmt.Println("here leaf")
	wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("parent exit error: ", ctx.Err())
			wg.Done()
			os.Exit(0)
		default:
			fmt.Println("hello handle 2st layer")
		}
	}
	wg.Done()
}

func main() {
	//eg33()
	myContext()
}
