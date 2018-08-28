package main

import "fmt"

func t1(i int) (j int) {
	defer func() {
		if p := recover(); p != nil {
			i = 1 - i
			j = 1
		}
	}()
	if i <= 0 {
		panic("please input a int >0")
	}
	j = i
	return j
	//return uint(i)
}
func main() {
	fmt.Println(t1(-1))
}
