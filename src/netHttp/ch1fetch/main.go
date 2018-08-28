// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://")) {
			var whole string
			whole = "http://" + url
			url = whole
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//write into a file
		f, _ := os.Create("a.dat")
		//io.Copy从reader:resp.Body复制到writer:f
		//_, err1 := io.Copy(f, resp.Body)
		err1 := resp.Write(f)
		_ = f.Close()
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "copy err : %v\n", err)
			os.Exit(1)
		}
		//打印出响应状态
		fmt.Println(resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
