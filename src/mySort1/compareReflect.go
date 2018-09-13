//比较空接口参数大小
//用于mergeSort及其他排序算法
package mySort1

import (
	"log"
	"reflect"
)

func compare(left, right interface{}) bool {
	vLef := reflect.ValueOf(left)
	vRig := reflect.ValueOf(right)
	switch vLef.Kind() {
	//待排序数组转为相应类型并比较
	//空接口无比较
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		//接口类型断言
		if vLef.Int() < vRig.Int() {
			return true
		} else {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		//接口类型断言
		if vLef.Uint() < vRig.Uint() {
			return true
		} else {
			return false
		}
	case reflect.Float32, reflect.Float64:
		if vLef.Float() < vRig.Float() {
			return true
		} else {
			return false
		}
	default:
		log.Panicln("ilegal type: int float32 float64")
	}
	return false
}
