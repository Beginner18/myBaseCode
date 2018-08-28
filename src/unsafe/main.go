//https://studygolang.com/articles/1414
//实现地址形式访问go包中不可访问变量，小写
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]int32, 3)
	for i, _ := range a {
		a[i] = int32(i + 1)
	}
	b := (*int32)(unsafe.Pointer(&a[0]))
	//指针运算需要先将指针转换为uintptr类型
	//c := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(unsafe.Sizeof(int32(*b)))))
	c := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + unsafe.Sizeof(int32(*b))))
	fmt.Printf("b type %T\n", b)
	fmt.Printf("c type:%T, c vale %d\n", c, *c)
	fmt.Println(a)
}
