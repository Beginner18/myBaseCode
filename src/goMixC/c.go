package main

//1 直接调用C代码嵌入go程序文件
/*
#include <stdio.h>
int Myadd( , )
{
}
*/
//2 C代码放入.c文件，调用同go
/*
#include "Myadd.c"
*/
//3 动态库　c文件
//编译o文件: gcc -fPIC -c Mymath.c //-fPIC与位置无关
//编译动态库lob.so: gcc -shared -o lib.so Mymath.o
//动态库与位置无关可以lib.so与编译好的go可执行程序
//一起复制到其他位置
/*
#cgo CFLAGS: -I./
cgo LDFLAGS: -L./ -lMyMath
#include "Mymath.h"
*/
import "C"

func main() {
	//直接调用
	//C.Myadd( , )
}
