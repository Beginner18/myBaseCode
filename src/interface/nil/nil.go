package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	fmt.Printf("the type of buf: %T\n", buf)
	fmt.Println((buf == nil))
	f(buf) // NOTE: subtly incorrect!
	//if debug {
	// ...use buf...
	//}
	fmt.Println(buf)
}
func f(out io.Writer) {
	// ...do something...
	fmt.Printf("the type of out: %T\n", out)
	fmt.Println("the value of out:", out)
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
