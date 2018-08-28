package main

import (
	"fmt"
	"log"
	"os"
)

//创建日志文件，将错误输出至日志文件中
func logErr(fileName string, info interface{}) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	logger := log.New(file, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(info)
}

func logPanic(fileName string, info interface{}) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	logger := log.New(file, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Panicln(info)
}
func logExit(fileName string, info interface{}) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	logger := log.New(file, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Fatalln(info)
}
func main() {
	fileName := "/home/yjj/goErr.log"
	err := fmt.Errorf("this is an error: %s\n", "hi")
	logErr(fileName, err)
	logPanic(fileName, "logpanic")
	logExit(fileName, "logExit")

}
