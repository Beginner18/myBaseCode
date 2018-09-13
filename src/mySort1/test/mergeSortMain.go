//mergeSort用法示例
package main

import (
	"fmt"
	"mySort1"
)

func main() {
	//定义空接口数组
	var arr = make([]interface{}, 6)
	//待排序数组
	arr1 := []float32{4.1, 5.3, 8, 1, 3, 2}
	//空接口数组赋值
	for i, _ := range arr {
		arr[i] = arr1[i]
	}
	fmt.Println("before: ", arr)
	mySort1.MergeSort(arr, 0, len(arr)-1)
	fmt.Println("after: ", arr)
	//test quickSort: int float32 float64
	quickSortTest()
	//quickSortUnit()
}
func quickSortTest() {
	//arr1 := []int{4, 2, 1, 3, 5, 0, 8, 20, 100, 44, 50}
	arr1 := []float32{4.0, 2.1, 1.1, 3.2, 5, 0, 8, 20, 100.5, 44, 50}
	var arr = make([]interface{}, len(arr1))
	for i, _ := range arr {
		arr[i] = arr1[i]
	}
	fmt.Println("float32 before quickSort: ", arr)
	mySort1.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("float32 after quickSort: ", arr)
}
func quickSortUnit() {
	//arr1 := []int{4, 2, 1, 3, 5, 0, 8, 20, 100, 44, 50}
	arr1 := []uint16{4, 2, 1, 3, 5, 0, 8, 20, 100, 44, 50}
	var arr = make([]interface{}, len(arr1))
	for i, _ := range arr {
		arr[i] = arr1[i]
	}
	fmt.Println("uint16 before quickSort: ", arr)
	mySort1.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("uint16 after quickSort: ", arr)
}
