package main

import (
	//"errors"
	"fmt"
)

type try struct {
	name string
	id   []string
	age  int
}
//以下两种定义方式均可以
/***************************************
//定义结构体指针
func main() {
	var try1 = &try{
		//name: 
		"123",
		//id:   
		[]string{"1", "2"},
		//age:  
		23,
	}
	fmt.Println(*try1)
}
*****************************/
/********************
//定义结构体
func main() {
	var try1 = try{
		name: "123",
		id: []string{"1", "2"},
		age: 23,
	}
	fmt.Println(*try1)
}
***************/
func main() {
	var try1 = &try{
		//name: 
		"123",
		//id:   
		[]string{"1", "2"},
		//age:  
		23,
	}
	fmt.Println(*try1)
}