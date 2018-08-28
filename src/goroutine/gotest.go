package main
import (
	"fmt"
	"time"
)
// 函数Publish在给定时间过期后打印text字符串到标准输出
// 该函数并不会阻塞而是立即返回
func Publish(text string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println("BREAKING NEWS:", text)
    }()    // 注意这里的括号。必须调用匿名函数
}
func main() {
    Publish("A goroutine starts a new thread of execution.", 5*time.Second)
    fmt.Println("Let’s hope the news will published before I leave.")
    // 等待发布新闻
    time.Sleep(10 * time.Second)
    fmt.Println("Ten seconds later: I’m leaving now.")
}

