package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	//	"sync"
)

//fmt.Scanf
//fmt.Fprintf
func rwfile(file3 *os.File) {
	var input []string
	for i := 0; i < 10; i++ {
		str1 := ""
		num, err := fmt.Scanf("%s", &str1)
		if str1 == "end" {
			break
		}
		input = append(input, str1)
		fmt.Printf("string number: %d, err: %#v\n", num, err)
	}
	fmt.Println(input)
	for i, str1 := range input {
		fmt.Fprintf(file3, "%d st: %s\n", i, str1)
	}
}

//rbuf.ReadSlice()(line []byte, err error)
//rbuf.ReadSring()(line string, err error)
func bufTest(file3 *os.File) {
	var f byte = byte(',')
	rbuf := bufio.NewReader(os.Stdin)
	//建立os.Writer
	//wbuf := bufio.NewWriter(os.Stdout)
	wbuf := bufio.NewWriter(file3)
	//line, err := rbuf.ReadSlice(f)
	line, err := rbuf.ReadString(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(line)
	fmt.Fprintln(file3, line)

	wbuf.WriteString("hello")
	//flush之后将对应字符串写入底层os.Writer
	wbuf.Flush()
}

//os.OPenfile(name, mode, perm)
//file.Read([]byte)
//file.ReadSrint()
func main() {
	//	file1Name := "try1.dat"
	//	file2Name := "try2.dat"
	file3Name := "try3.dat"
	//	file1, err1 := os.Create(file1Name)
	//	file2, err2 := os.Open(file2Name)
	file3, _ := os.OpenFile(file3Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file3.Close()
	defer file3.Sync()
	var test []byte = make([]byte, 4, 4)
	var offset1 int64 = 0
	var offset2 int64 = 4
	var err error
	pos1, _ := file3.Seek(offset1, os.SEEK_CUR) //current position
	pos2, _ := file3.Seek(offset2, os.SEEK_CUR) //current position + 4
	//_, err := file3.Read(test)
	file3.ReadAt(test, pos1)
	fmt.Println(string(test))
	_, err = file3.ReadAt(test, pos2)
	fmt.Println(string(test))
	_, err = file3.Read(test)
	fmt.Println(string(test))
	_, err = file3.ReadAt(test, pos1)
	fmt.Println(string(test))
	if err == io.EOF {
		fmt.Println("this is the end of file")
	}
	rwfile(file3)
	bufTest(file3)

}
