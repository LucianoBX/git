package sorts

func InsertionSort(arr []int) []int {
	for out := 1; out < len(arr); out++ {
		tmp := arr[out]
		in := out

		for ; in > 0 && arr[in-1] > tmp; in-- { // '>=' V.S. '>'
			arr[in] = arr[in-1]
		}
		arr[in] = tmp
	}
	return arr
}
