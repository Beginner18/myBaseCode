//1 https://studygolang.com/articles/1414
//2 实现地址形式访问go包中不可访问变量，小写
//3 unsafe.Pointer类似c中的void *类型,单纯通用指针类型，用于指针转换桥梁
//利用Pointer转换指针类型再转换为uintptr进行指针运算再转换为pointer再转
//换为需求指针对象类型，再解引操作
//4 uintptr为go内置类型，底层为int类型，可用于指针运算，GC
//认为uintptr不是指针，无法持有对象，uintptr对象会被回收
//5 注意指针对齐，系统层面, Alignof,变量占用空间类型, 1的倍数,4的倍数等

package main

import (
	"fmt"
	"unsafe"
)

type test struct {
	a byte
	b int32 //64位go默认int为int64
	c int64
}

//需要知道结构体的内部结构类型
func testAlign() {
	var str1 *test = new(test)
	str1.b = 12
	str1.c = 13
	//指针占用空间，8byte
	fmt.Println("指针占用空间:", unsafe.Sizeof(str1))
	//指针指向的结构体对象占用空间16
	fmt.Println("结构体占用空间: ", unsafe.Sizeof(*str1))
	fmt.Println("结构体对齐占用空间: ", unsafe.Alignof(str1))
	fmt.Println("结构体byte占用空间: ", unsafe.Sizeof(str1.a))
	fmt.Println("结构体byte对齐占用空间: ", unsafe.Alignof(str1.a))
	fmt.Println("结构体int32占用空间: ", unsafe.Sizeof(str1.b))
	fmt.Println("结构体int32对齐占用空间: ", unsafe.Alignof(str1.b))
	fmt.Println("结构体int64占用空间: ", unsafe.Sizeof(str1.c))
	fmt.Println("结构体int64对齐占用空间: ", unsafe.Alignof(str1.c))
	fmt.Println("指针法输出b: ", *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(str1)) + unsafe.Alignof(str1.b))))
	fmt.Println("指针法输出c: ", *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(str1)) + unsafe.Alignof(str1.c))))
	//可以直接调用offsetof,给出距离结构体或切片起始地址的距离地址间隔
	fmt.Println("指针法输出c: ", *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(str1)) + unsafe.Offsetof(str1.c))))
}

func test1() {
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
func main() {
	testAlign()
}
