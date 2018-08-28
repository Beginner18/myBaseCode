/**************
File: testFlag
Usage:
	1 go build testFlag.go
	2 ./testFlag -s / -n true a bc def
	3 在2中/为s的flag对应值，true为n的flag对应值， a bc def为flag.Args()值
	4 Args()用法较多可以有，Os.Arg等
****************/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//此中定义方式n和sep为指针
//var n = flag.Bool("n", false, "omit trailing newline")
//var sep = flag.String("s", " ", "separator")

func main() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")
	//解析后n和sep可用,且解析后可调动-help
	flag.Parse()
	//输出用-s分割的字符串
	//flag.Args()返回[]string为flag之后所有输入参数
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//os.Args[1:]为除可执行程序名外其他参数
	//os.Args类型为[]string
	fmt.Println(os.Args)
	fmt.Print(strings.Join(os.Args[2:], *sep))
	//n为true则输出换行符
	if !*n {
		fmt.Println()
	}
}
