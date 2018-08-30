package main

import (
	"fmt"
	"net"
)

func locServer(netf, hostName string) {
	//get local addr
	locAdd, _ := net.LookupHost(hostName)
	//get local addr+port
	locAddP := net.JoinHostPort(locAdd[0], "8080")
	//build listen
	locLisner, _ := net.Listen(netf, locAddP)
	//关闭监听
	defer locLisner.Close()
	fmt.Println("add:", locAdd, "addP:", locAddP,
		"lis:", locLisner.Addr())
	//build conn
	conNum := 0
	for {
		connSer, err := locLisner.Accept()
		if err != nil {
			continue
		}
		//统计连接成功次数
		conNum++
		//输出连接成功次数及客户端IP地址
		fmt.Println("con:", conNum, "clientAdd:", connSer.RemoteAddr().String())
		go handleCon(connSer)
	}
}
func handleCon(conn net.Conn) {
	getRes := make([]byte, 2048)
	for {
		getLen, err := conn.Read(getRes)
		if err != nil {
			return
		}
		conn.Write([]byte("hello"))
		fmt.Println("listen string: ", getLen, "res:", string(getRes), "err:", err)
	}
}

func main() {
	locServer("tcp", "yjj-Inspiron-3476")
}
