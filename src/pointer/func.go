package main
import (
	"fmt"
	"flag"
)
func t1(a *int) int{
	*a+=1
	return *a
}
func main(){
	try1 := flag.String("c","","")
	flag.Parse()
	if *try1=="1" {
		fmt.Println(*try1)
		fmt.Println(flag.Args())
	}
	a:=10
	fmt.Println("1st a:",a)
	fmt.Println(t1(&a))
	fmt.Println("2st a:",a)
}

