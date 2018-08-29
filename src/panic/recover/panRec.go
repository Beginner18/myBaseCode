/*1 defer函数调用recover, defer程序正常退出或panic
	执行: defer func(){ if p:= recover(); p != nil{};}()
	1.1 recover使程序从panic恢复，并返回panic value
	1.2 导致panic异常的函数不会继续运行，但正常返回
	1.3 未发生panic时调用recover，recover会返回nil
2 对panic value做特殊标记，当recover收到的p值为标记
	值则处理，其他情况继续异常:
	defer func(){
		switch p := recover(); p{
		//无panic异常
		case nil:
		//panic value为标记值，执行recover恢复+处理语句
		case bailout{}:
				err := fmt.Errorf("the error is ...")
		//panic异常且不是标记值继续panic
		default:
				panic(p)
		}
	}()





*/
package main

import "fmt"

func t1(i int) (j int) {
	defer func() {
		if p := recover(); p != nil {
			i = 1 - i
			j = 1
		}
	}()
	if i <= 0 {
		panic("please input a int >0")
	}
	j = i
	return j
	//return uint(i)
}
func main() {
	fmt.Println(t1(-1))
}
