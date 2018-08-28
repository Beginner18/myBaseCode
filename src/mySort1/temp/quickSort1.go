//快速排序
//1 取参考点pivot: arr[low] arr[mid] arr[high]的中位数
//2 将pivot放置合适位置
//3 二分排序
//4 待排序元素个数低于4个普通排序算法
package mySort1

import (
	"fmt"
	"log"
)

func median3(arr []int, low, high int) int {
	var pivot int
	mid := (low + high) / 2
	//arr[low]放置三者最小值
	if arr[low] > arr[mid] {
		swap(&arr[low], &arr[mid])
	}
	if arr[low] > arr[high] {
		swap(&arr[low], &arr[high])
	}
	//arr[high]放置三者最大者
	if arr[mid] > arr[high] {
		swap(&arr[mid], &arr[high])
	}
	pivot = arr[mid]
	//交换arr[high-1]为arr[mid]值
	swap(&arr[mid], &arr[high-1])
	//返回参考值pivot
	return pivot
}
func swap(left, right *int) {
	if left == nil || right == nil {
		log.Fatalln("empty input")
	}
	temp := *right
	*right = *left
	*left = temp
}
func QuickSort(arr []int, low, high int) {
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
				if arr[i] < pivot {
					i++
				} else {
					break
				}
			}
			//fmt.Println("arr[i]: ", arr[i])
			//右侧
			for {
				if arr[j] > pivot {
					j--
				} else {
					break
				}
			}
			//fmt.Println("arr[j]: ", arr[j])
			//将大数交换至pivot右侧
			//将小数交换至pivot左侧
			if i < j {
				swap(&arr[i], &arr[j])
			} else {
				break //pivot找到合适位置
			}
		}
		//pivot放置合适位置
		swap(&arr[i], &arr[high-1])
		//fmt.Println("arr: ", arr)
		//分
		QuickSort(arr, low, i-1)
		QuickSort(arr, i+1, high)
	} else {
		//待排序元素个数低于4个,simpleSort函数处理
		simpleSort(arr, low, high)
	}
}
func simpleSort(arr []int, low, high int) {
	if (high - low + 1) < 1 {
		fmt.Println("arr is empty")
	}
	if (high - low + 1) == 1 {
	}
	if (high - low + 1) == 2 {
		if arr[low] > arr[high] {
			swap(&arr[low], &arr[high])
		}
	}
	if (high - low + 1) == 3 {
		median3(arr, low, high)
	}
}
