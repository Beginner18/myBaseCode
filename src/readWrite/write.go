/*****************************************
File:   getSubordinateNodes.go
Created by yujiajia@2018-07-18
Description:
	Get all information of subordinate nodes belonged to a certain user,
	and write the information into a file named main's port.
Usage:
*****************************************/
package main

import (
	"flag"
	"fmt"
	//"api"
	//"io/ioutil"
	//"errors"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//1 以分配给用户的资源为最终资源状态，'el_allocated_resource'
	//2 subordinate node不为空显示添加作业按钮，点击该按钮并完成所有配置后，
	//触发getSubordinateNodes.go程序，获取用户的所有subordinate node信息，并保存为
	//文件名为主节点端口号的文件
	//3 服务器根据文件名将文件发送至对应主节点
	//4
	//fmt.Println("hello")
	var mainPort string = "10001"
	userStatus := flag.String("n", "", "")
	flag.Parse()
	//退出资源用户的用户ID
	offlineID := flag.Args()
	fmt.Println(*userStatus)
	fmt.Println(offlineID)
	//创建文件，若文件已存在则覆盖为空文件
	subordinateFile, errs := os.Create(mainPort)
	//check创建文件错误
	checkErr(errs)
	//写入字符串
	_, errs = subordinateFile.WriteString("10002 2 4\n")
	_, errs = subordinateFile.WriteString("10002 2 4\n")
	_, errs = subordinateFile.WriteString("10002 2 4\n")
	//check写入文件错误
	checkErr(errs)
	//关闭文件
	errs = subordinateFile.Close()
	checkErr(errs)
}
