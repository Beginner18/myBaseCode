/*
1 反射:
	程序运行时更新变量检查变量值，
	调用方法支持其内在操作,不需要
	编译时直到其类型
2 类型本身作为第一类的值类型处理
3 空接口,类型断言,无法检测未知类型
4 reflect.Valueof(..)只是值的copy
reflect.Valueof(&x).Elem()可取址




*/
package main

import (
	"fmt"
	"reflect"
)

func fun1() {
	//返回值为具体类型
	t := reflect.TypeOf(interface{}(1))
	fmt.Printf("type is:%T\n ", t)
	fmt.Println(t.String())
	v := reflect.ValueOf(interface{}(1))
	fmt.Printf("value is:%v, type is %T\n ", v, v)
	fmt.Println(v)
	//int value
	fmt.Println(v.Kind())
}

func main() {
	fun1()
}
