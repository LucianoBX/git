package main

import (
	"fmt"
	"os"
)

func main() {
	// 需要打开的文件名
	var fileOpen string
	fileOpen = "test"
	// fileOpen = fmt.Scanln("")

	// 打开文件
	read, err := os.Open("fileOpen")

	fmt.Println("vim-go")
}
