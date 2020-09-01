package main

import (
	"fmt"
	"github.com/goEx/sorts"
)

func main() {

	arr := []int{192, 168, 0, 1, 22, 80, 8080, 443, 22, 21, 23, 6379, 3721, 177}

	result := sorts.HeapSort(arr)

	fmt.Println("the result of sort is: ", result)
}
