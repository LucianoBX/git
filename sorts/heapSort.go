package sorts

import (
	"fmt"
)

type maxHeap struct {
	slice    []int
	heapSize int
}

// 自底向上调用maxHeap的堆调整方法
func buildMaxHeap(slice []int) maxHeap {
	h := maxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}

	fmt.Println("----build maxHeap success-----")
	fmt.Printf("the maxHeap is %v\n and size is %v\n", h, h.heapSize)
	return h
}

/*	调整以某个节点i为根的子树为大根堆
	大根堆：在大根堆中，对于以某个节点为根的子树，其各节点的值都不大于其根节点的值，即A[PARENT(i)] >= A[i]

*/
func (h maxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2

	max := i

	if l < h.size() && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.size() && h.slice[r] > h.slice[max] {
		max = r
	}

	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

func (h maxHeap) size() int { return h.heapSize } // ???

func HeapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	fmt.Println(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MaxHeapify(0)
	}

	return h.slice
}
