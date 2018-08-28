package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//读入字符串以换行符分隔
	//遇到"exit"推出
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "exit" {
			break
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
