/*package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {    
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    addrs := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for n, a := range addrs {
        fmt.Printf("%v: %v\n", n, a)
    }
}*/
package main
import (
	"fmt"
	"strconv"
)
//type int1 int
type int1 int
type jie1 struct {
	a1 string
	b1 string
}
//type float float64
type try1 interface{
	str()string
}
/*func (a float64)str()string{
	return string(a)
}*/
func (a int1)str()string{
	return strconv.Itoa(int(a))
	//fmt.Println(a)
}
func (a jie1)str()string{
	var a1 string = ""
	a1 =a1+ a.a1+";"
	a1 +=a.b1+";"
	return a1
	//fmt.Println(a.a1,"2",a.b1)
}
func printTry(a try1){
	fmt.Println(a.str())
}
func main(){
	//var b float64 = 3.2
	var a int1
	a=44
	b := jie1{"1","123"}
	/**********method 1
	var c try1
	c=a
	fmt.Println(c.str())
	c=b
	fmt.Println(c.str())
	*****end of method 1***/
	printTry(a)
	printTry(b)
}
