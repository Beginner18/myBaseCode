package main

import (
	"fmt"
	"os/exec"
)

func main() {
	exe := exec.Command("/home/yjj/learn/go/goCode/src/defer/test/temp/test", "-n", "config")
	res, err := exe.Output()
	fmt.Println("res:", string(res), "err:", err)
}
