package main

import "fmt"

// Student has two fields: Name and Age.
type student struct {
	Name string
	//Id   int
	Age int
}

// var a map
var st map[string]student
var stu []student
var a = [5]int{1, 2, 3, 4, 5}

func main() {

	stu := []student{
		{Name: "wang", Age: 20},
		{"li", 23},
		{"zhi", 28},
	}

	for ind, val := range stu {

		//		m[val.Name] = val
		fmt.Println(ind, val)
	}

	fmt.Println("-------------------------")
	ch1, ch2 := make(chan string, 10), make(chan int, 20)
	ch3 := make(chan int, 3)

	ch1 <- "hello"
	ch2 <- 1
	ch3 <- 2

	var (
		i string
		j int
	)

	select {
	case i = <-ch1:
		fmt.Println(">>>> This is i:", i)
	case j = <-ch2:
		defer fmt.Println(">>>>>  This is j:", j)
		//		panic(i)
	case k := <-ch3:
		fmt.Println(">>> This is k", k)
	}

	var sl []int
	fmt.Println(sl)
	fmt.Println("-------------")
	sli := make([]int, 3, 5)
	fmt.Println(sli)
	//	fmt.Println(stu)
	//	fmt.Println(m)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	type stru struct {
		i int
		b bool
	}

	for key, val := range s {
		m := make(map[int]stru)

		m[key] = val
		fmt.Printf("KEY: %s \t, MAP: %s\n", key, m[key])
	}

	fmt.Println("------------------------------------")
	pow := make([]int, 11)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	//
	// m = make(map[string]student)
	// fmt.Println(m)

	// ap := make([]int, 5, 8)
	//	ap = append(ap, 1, 2, 3)
	//	fmt.Println(ap)
	//	for i := 0; i < len(a)+1; i++ {
	//		fmt.Println(a[i])
	//	}
}
