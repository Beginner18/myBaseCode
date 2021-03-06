/*
//eg by Yu
package main

import (
	"fmt"
	"time"
)

func main() {
	as := make(chan int)
	//control main func
	bs := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			_, ok := <-as
			fmt.Printf("%dst: this is an output\n", i)
			if ok {
				fmt.Printf("%dst: as has a value\n", i)
			} else {
				fmt.Printf("%dst: as is empty,waitting for input\n", i)
			}
		}(i)
	}
	go func() {
		//as <- 1
		//as = make(chan int)
		fmt.Println("go routine input")
		//broadcasting the ready condition for other goroutines
		//that receive from as
		close(as)
		time.Sleep(2 * time.Second)
		fmt.Println("after close as")
		time.Sleep(1 * time.Second)
		close(bs)
	}()
	<-bs
}
*/
/*************************1 并发 不重复 无阻塞的cache, based on mutex**********
func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type entry struct {
    res   result
    ready chan struct{} // closed when res is ready
}

type Memo struct {
    f     Func
    mu    sync.Mutex // guards cache
    cache map[string]*entry
}

func New(f Func) *Memo {
    return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
    memo.mu.Lock()
    e := memo.cache[key]
    if e == nil {
        // This is the first request for this key.
        // This goroutine becomes responsible for computing
        // the value and broadcasting the ready condition.
        e = &entry{ready: make(chan struct{})}
        memo.cache[key] = e
        memo.mu.Unlock()

        e.res.value, e.res.err = memo.f(key)

        close(e.ready) // broadcast ready condition
    } else {
        // This is a repeat request for this key.
        memo.mu.Unlock()

        <-e.ready // wait for ready condition
    }
    return e.res.value, e.res.err
}

//useage memo
m := memo.New(httpGetBody)
for url := range incomingURLs() {
    start := time.Now()
    value, err := m.Get(url)
    if err != nil {
        log.Print(err)
    }
    fmt.Printf("%s, %s, %d bytes\n",
    url, time.Since(start), len(value.([]byte)))
}
***********************************************/

/*******************2 流水线循环**********
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	timeBeg := time.Now()
	a := make(chan struct{})
	b := make(chan struct{})
	//input := 1
	//i := 0
	for i := 0; i < 1000; i++ {
		//fmt.Println(i)
		wg.Add(3)
		go func(a chan<- struct{}) {
			//a<1
			//i++
			//fmt.Println(i)
			a <- struct{}{}
			wg.Done()
		}(a)
		go func(a <-chan struct{}, b chan<- struct{}) {
			//i++
			//fmt.Println(i)
			//wg.Done()
			//re2 := <-a
			//b <- re2
			b <- (<-a)
			wg.Done()
		}(a, b)
		go func(i int, b <-chan struct{}, a chan<- struct{}) {
			//i++
			fmt.Println(i)
			//wg.Done()
			//re2 := <-b
			//a <- re2
			a <- (<-b)
			wg.Done()
		}(i, b, a)
	}
	wg.Wait()
	fmt.Println("total time: ", time.Since(timeBeg))
	//fmt.Println("all done")
}
