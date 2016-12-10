// File Processing project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	//	DirProcess()
	//	CreateAndWriteFile()
	//	ReadFile()
	DeleteFile()
	fmt.Println("Hello World!")
}

// 目录处理
func DirProcess() {
	// 创建单级目录（如果不指定绝对路径在当前目录下创建）
	err := os.Mkdir("astaxie", 0777)
	CheckErr(err)
	// 创建多级目录
	err = os.MkdirAll("astaxie/test1/test2", 0777)
	CheckErr(err)
	// 删除单级目录，当前目录下有文件或者其它目录会报错
	//	err = os.Remove("astaxie")
	//	CheckErr(err)
	// 删除多级子目录
	err = os.RemoveAll("astaxie")
	CheckErr(err)

}

// 创建及写文件
func CreateAndWriteFile() {
	fileName := "astaxie.txt"
	fout, err := os.Create(fileName)
	CheckErr(err)
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("just a test!\r\n"))
	}
}

// 读文件
func ReadFile() {
	fileName := "astaxie.txt"
	file, err := os.Open(fileName)
	CheckErr(err)
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

// 删除文件
func DeleteFile() {
	// Go语言里面删除文件和删除文件夹是同一个函数
	fileName := "astaxie.txt"
	err := os.Remove(fileName)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
