//比较空接口参数大小
//用于mergeSort及其他排序算法
package mySort1

import (
	"log"
)

func compare(left, right interface{}) bool {
	switch left.(type) {
	//待排序数组转为相应类型并比较
	//空接口无比较
	case int:
		if left.(int) < right.(int) {
			return true
		} else {
			return false
		}
	case float32:
		if left.(float32) < right.(float32) {
			return true
		} else {
			return false
		}
	case float64:
		if left.(float64) < right.(float64) {
			return true
		} else {
			return false
		}
	default:
		log.Panicln("ilegal type: int float32 float64")
	}
	return false
}
