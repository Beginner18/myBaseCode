//1 ServrMux 本质上是一个 HTTP 请求路由器
//（或者叫多路复用器，Multiplexor).
//它把收到的请求与一组预先定义的 URL 路径列表做对比，
//然后在匹配到路径的时候调用关联的处理器（Handler).
//2 ListenAndServe(addr string, handler Handler)
//是因为 ServeMux 也有 ServeHTTP 方法，因此它也是个合法的 Handler。
//对我来说，将 ServerMux 用作一个特殊的Handler是一种简化.
//它不是自己输出响应而是将请求传递给注册到它的其他 Handler.
//这乍一听起来不是什么明显的飞跃-但在 Go 中将 Handler 链在一起是非常普遍的用法.
//3 具有ServerHTTP方法的结构体可以作为handler
//4 定义函数，令函数强制转换为handlerFunc, handlerFunc具有ServerHTTP方法

package main

import (
	"log"
	"net/http"
	"time"
)

//自定义结构体为handler
type timeHandler struct {
	format string
}

//定义结构体具有ServerHTTP方法即具有handler功能
func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	//建立serveMux可实现多种URL对应不同handler
	mux := http.NewServeMux()
	//建立重定位handler
	rh := http.RedirectHandler("http://www.baidu.com", 307)
	rhg := http.RedirectHandler("http://www.youku.com", 307)
	//建立具有ServerHTTP方法的结构体
	th := &timeHandler{format: time.RFC1123}
	// Convert the timeHandler function to a HandlerFunc type
	th1 := http.HandleFunc(timeHandler)
	th3339 := &timeHandler{format: time.RFC3339}
	//tens, _ := time.ParseDuration("10s")
	//建立超时handler
	//outTime := http.TimeoutHandler(rhg, tens, "out of time")
	//建立pattern与handler对应关系
	mux.Handle("/foo", rh)
	mux.Handle("/youku", rhg)
	mux.Handle("/time", th)
	mux.Handle("/time1", th1)
	mux.Handle("/time/rfc3339", th3339)
	log.Println("Listening...")
	//建立监听服务器
	http.ListenAndServe(":3000", mux)
}
