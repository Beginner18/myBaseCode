package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os/exec"
)

//建立RPC方法
type MonRPC string

func (*MonRPC) Monitor(inCmd string, res *string) error {
	resValue := monitor(inCmd)
	*res = resValue
	return nil
}

/*func (*MonRPC) Try1(inCmd string, res *int) error {
		*res = 111
		return nil
}
*/
func monitor(inCmd string) string {
	//"/bin/sh"也可以
	cmd := exec.Command("/bin/bash", "-c", inCmd)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println("the cmd error: ", err)
	}
	return string(res)
}

//建立服务端
func server(hostName string) {
	mon := new(MonRPC)
	//建立新server并注册
	ser := rpc.NewServer()
	ser.Register(mon)
	//建立listener监听
	//get local addr
	locAdd, _ := net.LookupHost(hostName)
	//get local addr+port
	locAddP := net.JoinHostPort(locAdd[0], "8889")
	lis, err := net.Listen("tcp", locAddP)
	defer lis.Close()
	if err != nil {
		fmt.Printf("listener error: ", err)
	}
	//单独进程处理RPC请求
	for {
		//建立连接请求
		con, err := lis.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
		}
		go func() {
			//con, err := ser.Accept(lis)
			//并发处理RPC请求
			ser.ServeConn(con)
		}()
	}
}
func main() {
	server("yjj-Inspiron-3476")
}
