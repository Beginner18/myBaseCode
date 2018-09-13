package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

//取地址, *sync.WaitGroup，函数间传递
func clientRPC(rpcName, serAdd, cmd string, wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", serAdd)
	if err != nil {
		fmt.Println("client dial error: ", err)
	}
	res := "the res is: "
	errRes := client.Call(rpcName, cmd, &res)
	if errRes != nil {
		fmt.Println("client call error: ", errRes)
	}
	if err == nil {
		fmt.Println("the remote process call res is: ", res)
	}
	wg.Done()
	return
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			clientRPC("MonRPC.Monitor", "127.0.1.1:8889", "ls", wg)
		}(&wg)
	}
	wg.Wait()
}
