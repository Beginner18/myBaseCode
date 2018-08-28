/***************1 互斥：通道堵塞 同步*********
var (
	// a binary semaphore guarding balance
    sema    = make(chan struct{}, 1)
    balance int
)

func Deposit(amount int) {
    sema <- struct{}{} // acquire token
    balance = balance + amount
    <-sema // release token
}

func Balance() int {
    sema <- struct{}{} // acquire token
    b := balance
    <-sema // release token
    return b
}
***************1 互斥：通道堵塞 同步******/

/***********2 sync.Mutex******************
//monitor/监控：函数(可导出)+互斥锁+变量(不可直接访问)
import "sync"

var (
    mu      sync.Mutex // guards balance
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}

func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}
***************2 sync.Mutex***************

/*************3 defer:函数返回时调用*****
//defer:临界区延伸到函数作用域后
//defer:函数返回之后或发生错误返回时调用
func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}
*************3 defer********************/

/**********4 函数拆分********************
//2中Deposit函数分为两部分:a b
//a为deposit函数，使用时需额外加锁，不可导出包
//b为Deposit，调用deposit加锁可导出包
func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    deposit(-amount)
    if balance < 0 {
        deposit(amount)
        return false // insufficient funds
    }
    return true
}

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    deposit(amount)
}

func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }
*************4 函数拆分：封装***************/

/*
//sync.RWMutex:多度单写
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock
    defer mu.RUnlock()
    return balance
}
*/
/****************main: defer*********************
package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	i++
	fmt.Printf("%dst i: %d\n", i+1, i)
	time.Sleep(1 * time.Second)
}
func main() {
	var i int
	defer hello(i)
	hello(i)
}

******main: defer*************************/
