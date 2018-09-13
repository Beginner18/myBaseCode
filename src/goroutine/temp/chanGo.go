/****************************
非原子级操作，时间片轮询
该文件为go并发编程与管道使用
1 同步管道：管道不带缓冲，
发送方会阻塞直到接收方从管道中接收了值
接送方或者发送方任一方没准备好则堵塞

2 异步通信　队列：管道带缓存，
发送方则会阻塞直到发送的值被拷贝到缓冲区内；
也就是说，这个信息只要还在管道里面没被使用，
那么该线程就会一直堵塞

3 如果缓冲区已满，
则意味着需要等待直到某个接收方获取到一个值。
同上接收方在有值可以接收之前会一直阻塞。

4 close 函数标志着不会再往某个管道发送值。
close之后对channel发送数据导致panic
在调用close之后，并且在之前发送的值都被接收后，
接收操作会返回一个零值，不会阻塞。
一个多返回值的接收操作会额外返回
一个布尔值用来指示返回的值是否发送操作传递的。

5 死锁
没有close（channel)导致另一线程一直等待

6 数据竞争
当两个线程并发地访问同一个变量，
并且其中至少一个访问是写操作时，
数据竞争就发生了

7 避免数据竞争:不要通过共享内存来通讯，而是通过通讯来共享内存
避免数据竞争的唯一方式是线程间同步访问所有的共享可变数据。
有几种方式能够实现这一目标。
Go语言中，通常是使用管道或者锁。
(sync和sync/atomic包中还有更低层次的机制暂不讨论）

8 互斥锁:
有时，通过显式加锁，而不是使用管道，来同步数据访问，可能更加便捷。
Go语言标准库为这一目的提供了一个互斥锁 - sync.Mutex。

要想这类加锁起效的话，关键之处在于：
所有对共享数据的访问，不管读写，仅当goroutine持有锁才能操作。
一个goroutine出错就足以破坏掉一个程序，引入数据竞争。
因此，应该设计一个自定义数据结构，具备明确的API，确保所有的同步都在数据结构内部完成。
下例中，我们构建了一个安全、易于使用的并发数据结构，AtomicInt，用于存储一个整型值。
任意数量的goroutine都能通过Add和Value方法安全地访问这个数值。

9 检测数据竞争
以下列方式运行程序：
go run -race *.go
使用局部变量传递避免数据竞争:线程级安全
使用闭包让每个routine使用独有变量

10 select语句:case为I/O操作
select语句是Go语言并发工具集中的终极工具
select用于从一组可能的通讯中选择一个进一步处理
如果任意一个通讯都可以进一步处理， 则从中随机选择一个，执行对应的语句。
否则，如果又没有默认分支（default case），select语句则会阻塞， 直到其中一个通讯完成
随机数生成、select语句为操作设置时间限制


***************************/
package main

import (
	"fmt"
	"sync"
	"time"
)

func chan1() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		ch <- "Hello!"
		ch <- "Hello!"
		//若不关闭ch，第二次调用时造成死锁
		close(ch)
	}()
	fmt.Println("1st", <-ch) // 输出字符串"Hello!"
	fmt.Println("2st", <-ch) // 输出零值 - 空字符串""，不会阻塞
	fmt.Println("3st", <-ch) // 再次打印输出空字符串""
	v, ok := <-ch            // 变量v的值为空字符串""，变量ok的值为false
	fmt.Println("4st", v, ok)
}

//依次读取管道值
func chanOrd() {
	var ch <-chan string = Producer()
	for s := range ch {
		fmt.Println("Consumed", s)
	}
}

func Producer() <-chan string {
	ch := make(chan string)
	go func() {
		ch <- string("海老握り")  // Ebi nigiri
		ch <- string("鮪とろ握り") // Toro nigiri
		close(ch)
	}()
	return ch
}

//管道控制线程
// 在给定时间过期时，Publish函数会打印text变量值到标准输出
// 在text变量值发布后，该函数会关闭管道wait
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	//空管道用来发送消息
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch) // 广播 - 一个关闭的管道都会发送一个零值
	}()
	return ch
}
func chanEmp() {
	wait := Publish("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for the news...")
	//fmt.Println(<-wait)
	<-wait // 等待结束  不然 The news 这个详细输入后 程序就退出了
	fmt.Println("The news is out, time to leave.")
}

//6 数据竞争
func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		// 译注：注意下面这一行
		n++ // 一次访问: 读, 递增, 写
		close(wait)
	}()
	// 译注：注意下面这一行
	//<-wait
	n++ // 另一次冲突的访问
	<-wait
	fmt.Println(n) // 输出：未指定
}

//7 避免数据竞争
func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0 // 仅为一个goroutine可见的局部变量.
		n++
		ch <- n // 数据从一个goroutine离开...
	}()
	n := <-ch // ...然后安全到达另一个goroutine.
	n++
	fmt.Println(n) // 输出: 2
}

//8 互斥锁
// AtomicInt是一个并发数据结构，持有一个整数值
// 该数据结构的零值为0
type AtomicInt struct {
	mu sync.Mutex // 锁，一次仅能被一个goroutine持有。
	n  int
}

// Add方法作为一个原子操作将n加到AtomicInt
func (a *AtomicInt) Add(n int) {
	a.mu.Lock() // 等待锁释放，然后持有它
	a.n += n
	a.mu.Unlock() // 释放锁
}

// Value方法返回a的值
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	//fmt.Println("*value:",&(a.n),"value:",((a.n)))
	n := a.n
	a.mu.Unlock() // 整个结构被解锁了
	return n
}

//互斥锁，封装的结构体可以实现并发访问
func lockItUp() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1) // 一个访问
		close(wait)
	}()
	n.Add(1) // 另一个并发访问
	<-wait
	fmt.Println(n.Value()) // 输出: 2
}

//测试WaitGroup
//对一组线程，.Done()线程计数减一
//.Wait()阻塞线程直到计数为0
func groupTry(a int) {
	var goGroup sync.WaitGroup
	goGroup.Add(a)
	for i := 0; i < a; i++ {
		//函数调用+定义
		go func(b int) {
			//i与for循环主routine中的i共享内存
			//go routine不分顺序
			fmt.Println("hello ", i+1)
			goGroup.Done()
			fmt.Println("goGroup:", goGroup)
			fmt.Println("b is:", b)
		}(12)
	}
	goGroup.Wait()
	fmt.Println("a", a)
}

func raceGroup() {
	var wg sync.WaitGroup
	wg.Add(5)
	// 译注：注意下面这行代码中的i++
	for i := 0; i < 5; i++ {
		//使用局部变量作为线程间数据传输避免数据竞争
		go func(i int) {
			// 注意下一行代码会输出什么？为什么？
			fmt.Println(i) // 6个goroutine共享变量i
			wg.Done()
		}(i)
	}
	wg.Wait() // 等待所有（5个）goroutine运行结束
	fmt.Println()
}

//使用闭包使每个routine使用唯一变量避免数据竞争
/**********************************************
func alsoCorrect() {
    var wg sync.WaitGroup // 使用 WaitGroup 高级货
    wg.Add(5)
    for i := 0; i < 5; i++ {
        n := i // 为每个闭包创建一个独有的变量
        go func() {
            fmt.Print(n)
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println()
}
**********************************************/
// RandomBits函数 返回一个管道，用于产生一个比特随机序列
func RandomBits() int {

	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0: // 注意：分支没有对应的处理语句
			case ch <- 1:
			}
		}
	}()
	return <-ch
}

//go routine && select
func goSel() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // 为一个未匹配的发送操作提供空间
	//为发送的消息内容提供空间
	info := make(chan string, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, info, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		// 没有待处理的发送操作
		fmt.Println("null")
	}
}

// 函数Seek 发送一个name到match管道或从match管道接收一个peer，结束时通知wait group
func Seek(name string, match chan string, info chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s,and the info is: %s\n", peer, name, <-info)
	case match <- name:
		info <- "This is just a info"
		// 等待某个goroutine接收我的消息
	}
	wg.Done()
}
func main() {
	//chan1()
	//chanOrd()
	//chanEmp()
	//race()
	//sharingIsCaring()
	//lockItUp()
	//groupTry(10)
	//raceGroup()
	/*
		for i:=0;i<10;i++{
			fmt.Print(RandomBits(),",")
			if i==9 {
				fmt.Print("\n")
			}
		}
	*/
	//cur := time.Now()
	//fmt.Println(cur.Local())
	goSel()
}
