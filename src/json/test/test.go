package main

import (
	"fmt"
	"json"
)

//结构体需要以大写字符开头，结构体内容也是
type T struct {
	A string `json: "1st" `
	B string
}

func main() {
	fileName, logFile := "/home/yjj/json.dat", "/home/yjj/log.dat"
	//test := com.JobInfo{"12", "23", "1", "t", "wait"}
	test := []T{
		{A: "h",
			B: "l"},
		{A: "m",
			B: "n"}}
	//test.A = "hello"
	//test.B = "hi"
	var res []T
	json.WriteJson(fileName, logFile, test)
	json.ReadJson(fileName, logFile, &res)
	fmt.Println("the read res json:", res)
}
