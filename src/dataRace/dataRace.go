/***************************************************
					数据竞争
[Issue]:
	数据竞争: 两个以上的goroutine并发访问相同的变量，
	且至少其中一个为写操作时发生。
[Solution]:
	1 多个goroutine并发访问，不写变量.
	2 避免从多个goroutine访问变量,
	其他goroutine通过channel通信获得，数据同步，
	不要使用共享数据来通信，使用通信来共享数据。
	3 互斥，同一时刻最多一个goroutine访问.
***************************************************/
/************2 bank问题****************************
// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
    var balance int // balance is confined to teller goroutine
    for {
        select {
        case amount := <-deposits:
            balance += amount
        case balances <- balance:
        }
    }
}

func init() {
    go teller() // start the monitor goroutine
}
**************2 bank问题************************/

/************2 流水线，串行绑定****************
type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
    for {
        cake := new(Cake)
        cake.state = "cooked"
        cooked <- cake // baker never touches this cake again
    }
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
    for cake := range cooked {
        cake.state = "iced"
        iced <- cake // icer never touches this cake again
    }
}
******************2 流水线，串行绑定************/

