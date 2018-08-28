//快速排序
//1 取参考点pivot: arr[low] arr[mid] arr[high]的中位数
//2 将pivot放置合适位置
//3 二分排序
//4 待排序元素个数低于4个普通排序算法
//5 go语言可以直接交换，不需要swap
package mySort1

import (
	"fmt"
	"log"
)

func median3(arr []interface{}, low, high int) (pivot interface{}) {
	//var pivot int
	mid := (low + high) / 2
	//arr[low]放置三者最小值
	if !compare(arr[low], arr[mid]) {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if !compare(arr[low], arr[high]) {
		arr[low], arr[high] = arr[high], arr[low]
	}
	//arr[high]放置三者最大者
	if !compare(arr[mid], arr[high]) {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	pivot = arr[mid]
	//交换arr[high-1]为arr[mid]值
	arr[mid], arr[high-1] = arr[high-1], arr[mid]
	//返回参考值pivot
	return pivot
}

func QuickSort(arr []interface{}, low, high int) {
	if low > high || high > (len(arr)-1) {
		log.Panicln("input error: low > high")
	}
	if (low + 3) <= high {
		pivot := median3(arr, low, high)
		//fmt.Println("pivot: ", pivot)
		//fmt.Println("pivot arr: ", arr)
		i, j := low+1, (high - 2)
		for {
			//左侧
			for {
				if compare(arr[i], pivot) {
					i++
				} else {
					break
				}
			}
			//fmt.Println("arr[i]: ", arr[i])
			//右侧
			for {
				if !compare(arr[j], pivot) {
					j--
				} else {
					break
				}
			}
			//fmt.Println("arr[j]: ", arr[j])
			//将大数交换至pivot右侧
			//将小数交换至pivot左侧
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				break //pivot找到合适位置
			}
		}
		//pivot放置合适位置
		arr[i], arr[high-1] = arr[high-1], arr[i]
		//fmt.Println("arr: ", arr)
		//分
		QuickSort(arr, low, i-1)
		QuickSort(arr, i+1, high)
	} else {
		//待排序元素个数低于4个,simpleSort函数处理
		simpleSort(arr, low, high)
	}
}
func simpleSort(arr []interface{}, low, high int) {
	if (high - low + 1) < 1 {
		fmt.Println("arr is empty")
	}
	if (high - low + 1) == 1 {
	}
	if (high - low + 1) == 2 {
		if !compare(arr[low], arr[high]) {
			arr[low], arr[high] = arr[high], arr[low]
		}
	}
	if (high - low + 1) == 3 {
		median3(arr, low, high)
	}
}
