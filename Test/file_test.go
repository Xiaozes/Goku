package test

import (
	"Goku/Gfile"
	"fmt"
	"log"
	"testing"
)

func init() {
	log.Println("run test")
}

func TestFile(t *testing.T) {
	something := Gfile.ReadFileToString("test.txt")
	Gfile.WriteString("test.txt", "这是测试")
	if Gfile.CheckExist("test.txt") {
		fmt.Println("test.txt 存在")
	} else {
		fmt.Println("文件不存在")
	}
	fmt.Println(Gfile.GetFileList("C:\\Users\\Administrator\\Desktop\\Project\\Goproject\\Goku"))
	//Gfile.UnZip("test", "test.zip")
	fmt.Printf("Readtxt something %s", something)
}
