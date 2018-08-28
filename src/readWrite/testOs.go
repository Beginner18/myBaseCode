/***********************************
File: testOs.go
Created by yujiajia@20180703
[Description]: An example from 'Go programming Language' for arg in os package.
**********************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	var s,sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

