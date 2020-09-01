package sorts

import (
	"fmt"
)

/*
 * 2020-07-24 by Lex.B
 */

// ShellSort, ascending
// NOTE:注：为方便记忆算法，我习惯将其记作“三层for循环+if” ------ for（for（for（if）））
func ShellSort(arr []int) []int {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		fmt.Println(">>> increment is:\t", d)
		// i即将插入元素的角标，作为每一组比较数据的最后一个元素角标
		fmt.Println("arr before this loop of sorting is: \t", arr)
		for i := d; i < len(arr); i++ {
			// j代表与i同一组元素的角标
			// arr[j-d] < arr[j]; j为即将插入的元素角标
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}
