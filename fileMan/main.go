package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// 需要打开的文件名
	filepath := "hello"
	// 打开文件
	file, err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 要写入的语句
	str := "Hello, world!\n"

	// 使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 9; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存，因此在调用writerString方法时，其实内容先写到缓存的。

	writer.Flush()
}
