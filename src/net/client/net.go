//test for net package
package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func basicTest() {
	//1 返回本地系统网络接口的地址列表
	//	类似linux: ifconfig
	add, err := net.InterfaceAddrs()
	errHandle(err)
	fmt.Println("return value of",
		"InterfaceAdd:", add)
	//2 返回本地系统网络接口列表
	//interface结构体:
	//index:编号, MTU: max transport unit
	//Name: 接口名eg, en0 lo0 eth0.100
	//HardwareAddr:eg IEEE MAC-48 EUI-48
	//EUI-64格式
	//Flages: 接口属性eg up|broadcast|loopback|p2p|multicast
	inter, err := net.Interfaces()
	errHandle(err)
	fmt.Println("return value of",
		"Interface:", inter)
	//3 根据IP查找localHost
	//net.LookupHost|LookupCNAME|LookupIP|LookupMX
	//|LookupNS("www.baidu.com)
	localHost, _ := net.LookupAddr("192.168.4.125")
	fmt.Println("return value of",
		"localHost:", localHost)
	//4 conn接口: read, write, close,localAdd,remoteAdd,
	//setDeadline, setReadDeadline, setWriteDeadline接口方法
	//面向流的网络连接，多线程可以共用一个conn
	connTry, err := net.Dial("tcp", "127.0.1.1:8080")
	errHandle(err)
	//add为ip:port
	fmt.Println("conn.remoteadd: ", connTry.RemoteAddr(),
		"conn.locaAdd:", connTry.LocalAddr())
	//4.1 conn接口中写入数据
	var sendInfo []byte
	//建立Buffer结构体，具有写入字符串方法
	sendBuf := bytes.NewBuffer(sendInfo)
	sendBuf.WriteString("go")
	fmt.Println("send buff:", sendBuf.Bytes())
	resLen, _ := connTry.Write(sendBuf.Bytes())
	fmt.Printf("write to server: %s, len: %d\n",
		sendBuf.String(), resLen)
	time.Sleep(5 * time.Second)
	fmt.Println("sleep done")
	//4.2 conn接口: read write处理
	//4.2.1 从连接中读取字节
	//给字节分配容量
	content := make([]byte, 2048)
	lens, err := connTry.Read(content)
	fmt.Printf("read %d from server:%s ,err:%v\n",
		lens, string(content), err)
	//type dialer的Dial方法同net.Dial，
	//只是type dialer结构体有更多变量比如timeout等

}
func errHandle(err error) {
	if err != nil {
		fmt.Println("the err is: ", err)
	}
}
func main() {
	basicTest()
}
