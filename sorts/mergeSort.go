package sorts

func merge(a, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	var middle = len(arr) / 2
	var a = MergeSort(arr[:middle])
	var b = MergeSort(arr[middle:])

	return merge(a, b)

}
