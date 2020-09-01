package main

import (
	"errors"
	"fmt"
	"os"
)

// 定义一个
type Queue struct {
	maxSize     int
	array       [5]int
	front, rear int
}

// 添加数据到队列
func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("queue full")
	}

	q.rear++
	q.array[q.rear] = val
	return
}

// 取出队列数据
func (q *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if q.rear == q.front {
		return -1, errors.New("queue empty")
	}

	q.front++
	val = q.array[q.front]
	return val, err
}

// 显示队列
func (q *Queue) ShowQueue() {

	if q.front == q.rear {
		errors.New("Emputy Queue")
		return
	}

	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("%v value is %v\t", i, q.array[i])
	}
	fmt.Println()

}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var (
		key string
		val int
	)

	for {
		fmt.Println("1. Add添加数据")
		fmt.Println("2. Get添加数据")
		fmt.Println("3. Show添加数据")
		fmt.Println("4. Exit添加数据")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("请输入要插入的队列")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列OK")
			}
		case "2":
			fmt.Println("GET")
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println("取出错误，err：%v", err)
			}

			fmt.Println(val)

		case "3":

			queue.ShowQueue()
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Input false")
		}
	}

}
