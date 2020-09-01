package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	// 获得文件路径
	//filepath := fmt.Scanf("")
	filepath := "test.txt"

	file, err := os.OpenFile(filepath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 要写入的字符串
	fmt.Println("输入要写入的字符")
	str := ""
	_, err = fmt.Scanln(&str)
	if err != nil {
		str = "没有读取到可用数据"
	}

	str = fmt.Sprintf("%s\r\n", str)
	// 打开输入接口
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//
	writer.Flush()
}
