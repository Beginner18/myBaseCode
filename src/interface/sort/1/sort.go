/*
package sort

type Interface interface {
	Len() int
	Less(i, j int) bool // i, j are indices of sequence elements
	Swap(i, j int)
}
*/
package main

import (
	"fmt"
	"sort"
)

//无法直接用[]string，方法内必须是type类型
type StringSlice []string

func (a StringSlice) Len() int {
	return len(a)
}
func (a StringSlice) Less(i, j int) bool {
	return a[i] < a[j]
}
func (a StringSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	//var a sort.Interface
	a := []string{"a", "z", "f", "b"}
	fmt.Printf("before sort: %s\n", a)
	sort.Sort(StringSlice(a))
	fmt.Printf("after sort: %s\n", a)
}
