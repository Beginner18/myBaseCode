//if k<=len(a)
//the kth largest value of L[] is,
//the (k+1)th largest value of a[];
//借助quicksort 思想，获得res在a[]，
//下标为i，满足len(a)-i=(k+1);
//则a[i]即为待求值，数组下标从0开始
package main

import "fmt"

func quickSort(arrA []int, low int, high int) {
	key := arrA[low] //取出第一项
	p := low
	i, j := low, high
	for i <= j {
		//由后开始向前搜索(j--)，找到第一个小于key的值arrA[j]
		for j >= p && arrA[j] >= key {
			j--
		}
		//第一个小于key的值 赋给 arrA[p]
		if j >= p {
			arrA[p] = arrA[j]
			p = j
		}

		if arrA[i] <= key && i <= p {
			i++
		}
		if i < p {
			arrA[p] = arrA[i]
			p = i
		}
		arrA[p] = key
		if p-low > 1 {
			quickSort(arrA, low, p-1)
		}
		if high-p > 1 {
			quickSort(arrA, p+1, high)
		}
	}

}
func findKst(k int, arrA []int) int {
	if k > len(arrA) {
		panic("k > len(arrA)")
	}
	pos := len(arrA) - (k + 1)
	res := arrA[pos]
	return res
}

func main() {
	arrA := []int{5, 4, 3, 2, 1}
	fmt.Println(arrA)
	quickSort(arrA, 0, len(arrA)-1)
	fmt.Println(arrA)
	k := 2
	res := findKst(k, arrA)
	fmt.Println(res)
}
