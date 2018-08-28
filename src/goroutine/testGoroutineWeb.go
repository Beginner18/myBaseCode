package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start := time.Now()
	ch := make(chan string)//传输字符串的通道
	for _,url := range os.Args[1:] {
		go fetch(url,ch)  //启动一个goroutine
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string){
	start := time.Now()
	resp,err := http.Get(url) //获取url对应指针，类似打开文件指针
	if err !=nil {
		ch <- fmt.Sprint(err) //将err发送到ch通道
		return
	}
	nbytes,err := io.Copy(ioutil.Discard, resp.Body) //io.Copy将2复制到1并返回2中总字符数int类型及err
	if err !=nil {
		ch <- fmt.Sprintf("while reading %s: %v",url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}
