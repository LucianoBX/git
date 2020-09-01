package sorts

func BubbleSort(arr []int) []int {

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				swap(arr, i, i+1)
				swapped = true
			}
		}
	}
	return arr
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
