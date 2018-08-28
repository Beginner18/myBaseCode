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
}
