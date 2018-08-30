package json

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log/mylog"
	"os"
)

//结构体需要以大写字符开头，结构体内容也是
type T struct {
	A string `json: "1st" `
	B string
}

//将读取的json格式转换为任意结构体

func ReadJson(fileName string, logFile string, resInfo interface{}) {
	resRead, err := ioutil.ReadFile(fileName)
	fmt.Println("read from file:", resRead)
	if err != io.EOF && err != nil {
		mylog.LogPanic(logFile, err)
	}
	//err = json.Unmarshal(resRead[:len(resRead)], resInfo)
	//json.Unmarshal(r1, r2)参数r2为指针
	err = json.Unmarshal(resRead, resInfo)
	if err != nil {
		mylog.LogPanic(logFile, err)
	}
	//return res

}

//将任意结构体转换为json并写入文件
func WriteJson(fileName string, logFile string, resInfo interface{}) {
	//若文件已存在会截断文件
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	//将错误信息记录至日志文件
	if err != nil {
		mylog.LogPanic(logFile, err)
	}
	fmt.Println("resInfo:", resInfo)
	//将结构体转为json　[]byte
	//resByte := make([]byte, 2064)
	resByte, err := json.Marshal(resInfo)
	if err != nil {
		mylog.LogPanic(logFile, err)
	}
	//fmt.Println(resByte)
	//将转换好的json格式文件保存至目录文件
	file.Write(resByte)
}

//错误日志记录文件，输入参数filename string, interface{}
//mylog.LogErr("/home/yjj/test.dat", "this is a test")
