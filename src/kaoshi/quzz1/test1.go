//test1
//func main is just for test
package main

import "fmt"

//combine str1 and str2, and delete all same chars
func combine(str1, str2 string) string {
	var res string
	len1 := len(str1)
	len2 := len(str2)
	//mark the number of same chars
	var sameNum int = 0
	for len1 > 0 && sameNum < len2 {
		if str1[len1-1] == str2[sameNum] {
			len1--
			sameNum++
		} else {
			break
		}
	}
	//combine str1 and str2
	res = str1[0:len1] + str2[sameNum:len2]
	return res

}

//for test
func main() {
	var str1, str2 string
	str1 = "gl"
	str2 = "lg11"
	res := combine(str1, str2)
	fmt.Printf("str1: %s, str2: %s, res: %s\n", str1, str2, res)
}
