//merge sort
//1 divide: 中分，仅存在一个变量时不分
//2 merge: 合并子列，若一个子列为空则
//直接复制另外一个子列
//3 需要与待排序数组等大小数组
package mySort

import (
	//"fmt"
	"log"
)

//利用空接口实现任意类型，空接口类型不存在比较
//比较需利用a.(int)将空接口转换为其他类型
//通过a.(type)可以判断空接口的实际类型
func MergeSort(arr []interface{}, low, high int) {
	desArr := make([]interface{}, high+1)
	if len(desArr) < 1 {
		log.Panicln(" short of memory")
	}
	mergeSort(arr, desArr, low, high)
}
func mergeSort(arr, desArr []interface{}, low, high int) {
	if len(arr) < 1 || len(desArr) < 1 || high > (len(arr)-1) || len(desArr) < len(arr) || (low > high) {
		log.Fatalf("func mergeSort: input error\n")
	}
	i, j := low, high
	if low < high {
		mid := (i + j) / 2
		mergeSort(arr, desArr, i, mid)
		mergeSort(arr, desArr, mid+1, j)
		merge(arr, desArr, i, mid, mid+1, j)

	}

}
func merge(arr, desArr []interface{}, lowleft, highleft, lowright, highright int) {
	i1, j1, i2, j2 := lowleft, highleft, lowright, highright
	if len(arr) < 1 || len(desArr) < 1 || highright > (len(arr)-1) ||
		len(desArr) < len(arr) || (highleft > lowright) ||
		(lowleft > highleft) || (lowright > highright) {
		log.Fatalf("func merge: input error \n")
	}
	len := (j1 - i1) + (j2 - i2) + 2
	var num int = 0
	//merge
	for {
		if i1 > j1 || i2 > j2 {
			break
		}
		//判定待排序数组类型
		switch arr[lowleft].(type) {
		//待排序数组转为相应类型并比较
		//空接口无比较
		case int:
			if arr[i1].(int) < arr[i2].(int) {
				desArr[num] = arr[i1]
				i1++
				num++
			} else {
				desArr[num] = arr[i2]
				i2++
				num++
			}
		case float32:
			if arr[i1].(float32) < arr[i2].(float32) {
				desArr[num] = arr[i1]
				i1++
				num++
			} else {
				desArr[num] = arr[i2]
				i2++
				num++
			}
		case float64:
			if arr[i1].(float64) < arr[i2].(float64) {
				desArr[num] = arr[i1]
				i1++
				num++
			} else {
				desArr[num] = arr[i2]
				i2++
				num++
			}
		default:
			log.Panicln("ilegal type")
		}
	}
	//copy
	if i1 <= j1 {
		copy(desArr[num:len], arr[i1:(j1+1)])
	}
	if i2 <= j2 {
		copy(desArr[num:len], arr[i2:(j2+1)])
	}
	//copy destArr to arr
	copy(arr[lowleft:(highright+1)], desArr[0:len])
}
