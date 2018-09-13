package main

import "net/http"
import "fmt"
import "runtime"
import "sync"
//import "time"


var max=11
//var maxint=11
func main() {
	runtime.GOMAXPROCS(6)
	var goGroup sync.WaitGroup
	goGroup.Add(max)
	for i:=0;i<max;i++ {
		go func (a int) {
			ex,_:=http.Get("http://192.168.2.18")
			fmt.Println(ex.Status)
			goGroup.Done()
		}(i)
		//fmt.Println(<-t1)
	}
	goGroup.Wait()
	fmt.Println("here done")
}
