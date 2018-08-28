/*****************************************
File:   getSlaveNodes.go
Created by yujiajia@2018-07-18
Description:
	Get all information of slave nodes belonged to a certain user, 
	and write the information into a file named master's port.
Usage:
*****************************************/
package main

import (
	"fmt"
	"flag"
	//"api"
	//"io/ioutil"
	//"errors"
	"os"
)
func checkErr(e error){
	if e!=nil{
		panic(e)
	}
}

func main() {
	//1 以分配给用户的资源为最终资源状态，'el_allocated_resource'
	//2 slave node不为空显示添加作业按钮，点击该按钮并完成所有配置后，
	//触发getSlaveNodes.go程序，获取用户的所有slave node信息，并保存为
	//文件名为主节点端口号的文件
	//3 服务器根据文件名将文件发送至对应主节点
	//4
	//fmt.Println("hello")
	userStatus := flag.String("n","","")
	flag.Parse()
	//退出资源用户的用户ID
	offlineID:= flag.Args()
	fmt.Println(*userStatus)
	fmt.Println(offlineID)
	tryName := "try.dat"
	try,errs := os.Open(tryName)
	checkErr(errs)
	b := make([]byte,100)
	lenb,errs := try.Read(b)
	checkErr(errs)
	fmt.Println("the length of b:",lenb,"the content is: ",string(b))

}
