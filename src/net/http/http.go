//1 servrMux与Handler最为重要
//2 servrMux: 请求路由(多路复用路由)，收到的请求与一组
//预先定义的URL路径列表做对比然后在匹配到路径的时候调
//用关联的处理器,即关联模式和Handler遇到匹配模式则直接
//调用Handler
//3 Handler:负责输出HTTP响应的头和正文，满足http.Handler
//接口的对象都可作为一个处理器，即对象只要有签名的
//ServerHTTPservrMux方法: ServerHTTP(ResoseWriter, *Request)
//serveMux满足此需求
//4 主函数传参的闭包处理
//5 自定义handler:定义满足handler接口方法的方法
package main

import (
	"fmt"
	"net/http"
)

func testRedirc() {
	mux := http.NewServeMux()
	//建立handler
	rh := http.RedirectHandler("https://www.bilibili.com", 307)
	//建立请求模式与handler的匹配
	mux.Handle("/bili", rh)
	fmt.Println("Listening...")
	//服务器监听并提供mux服务
	http.ListenAndServe(":5000", mux)
}

//4
//方式1显示转换
/*
func closePak(word []string) http.Handler {
	case1 := func(w http.ResponseWriter, r *http.Request) {
		s := ""
		for _, i := range word {
			s += i
			s += "/"
		}
		w.Write([]byte(s))
	}
	//需要将函数转为HandlerFunc格式，该格式具有Handler接口方法
	return http.HandlerFunc(case1)
}
*/
//方式2隐式转换
func closePak(word []string) http.HandlerFunc {
	case1 := func(w http.ResponseWriter, r *http.Request) {
		s := ""
		for _, i := range word {
			s += i
			s += "/"
		}
		w.Write([]byte(s))
	}
	//需要将函数转为HandlerFunc格式，该格式具有Handler接口方法
	return case1
}

//使用默认serverMux
func testClosePak(word ...string) {
	http.Handle("/hi", closePak(word))
	http.ListenAndServe(":8080", nil)
}

//5 自定义结构体,具有ServeMux方法
type word struct {
	wd []string
}

//定义ServeHTTP方法，则结构体可隐式转换为Handler接口
func (wr *word) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := ""
	for _, i := range wr.wd {
		s += i
		s += "/"
	}
	w.Write([]byte(s))
}
func myHandler(w1, w2 []string) {
	myMux := http.NewServeMux()
	//两个handler
	myh1 := &word{w1}
	myh2 := &word{w2}
	//注册
	myMux.Handle("/", http.FileServer(http.Dir("/home/yjj/learn/go/goCode/src/net/http/res")))
	myMux.Handle("/1", myh1)
	myMux.Handle("/2", myh2)
	http.ListenAndServe(":8080", myMux)
}
func main() {
	//testRedirc()
	//testClosePak("helllo", "ok")
	var w1 []string
	w1 = append(w1, "hello")
	w1 = append(w1, "welcome to 1")
	var w2 []string
	w2 = append(w2, "hello")
	w2 = append(w2, "welcome to 1")
	myHandler(w1, w2)
}
