package main

import (
	"flag"
	"fmt"
)

func main() {
	userStatus := flag.String("n", "", "")
	flag.Parse()
	//退出资源用户的用户ID
	if *userStatus == "config" {
		fmt.Println("this is ok, and exec -n=config")
	}
}
