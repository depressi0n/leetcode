package question

import (
	"container/heap"
	"sort"
)

var size int

type intSlice []int

func (h *intSlice) Len() int {
	return len(*h)
}

func (h *intSlice) Less(i, j int) bool {
	return (*h)[i] <= (*h)[j]
}

func (h *intSlice) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intSlice) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intSlice) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KthLargestValue(matrix [][]int, k int) int {
	var help intSlice
	heap.Init(&help)
	size = k
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// prefixXSum[i][j] 表示第i行前j个数的异或和
	prefixXSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		prefixXSum[i] = make([]int, len(matrix[0]))
		prefixXSum[i][0] = matrix[i][0]
		for j := 1; j < len(matrix[0]); j++ {
			prefixXSum[i][j] = prefixXSum[i][j-1] ^ matrix[i][j]
		}
	}
	// prefixYSum[i][j] 表示第j列前i个数的异或和
	prefixYSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		prefixYSum[i] = make([]int, len(matrix[0]))
	}
	for j := 0; j < len(matrix[0]); j++ {
		prefixYSum[0][j] = matrix[0][j]
		for i := 1; i < len(matrix); i++ {
			prefixYSum[i][j] = prefixYSum[i-1][j] ^ matrix[i][j]
		}
	}
	pushInHeap := func(x int) {
		heap.Push(&help, x)
		if len(help) > size {
			heap.Pop(&help)
		}
	}
	sum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[0]))
	}
	sum[0][0] = matrix[0][0]
	pushInHeap(sum[0][0])
	for i := 1; i < len(matrix[0]); i++ {
		sum[0][i] = prefixXSum[0][i]
		pushInHeap(sum[0][i])
	}
	for j := 1; j < len(matrix); j++ {
		sum[j][0] = prefixYSum[j][0]
		pushInHeap(sum[j][0])
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			sum[i][j] = sum[i-1][j-1] ^ prefixXSum[i][j-1] ^ prefixYSum[i-1][j] ^ matrix[i][j]
			pushInHeap(sum[i][j])
		}
	}
	// 找到第k个
	return heap.Pop(&help).(int)
}

// 省内存和加速，只需要按照行或者按照列开始就可以，这里选择按照行开始，并且可以用两个数组替代整个矩阵，
// TODO:随后使用直接排序算法或者快排的思想找第k个
func KthLargestValue2(matrix [][]int, k int) int {
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	pushInHeap := func(x int) {
		res = append(res, x)
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	oldLine := make([]int, len(matrix[0]))
	newLine := make([]int, len(matrix[0]))
	// 初始化
	oldLine[0] = matrix[0][0]
	pushInHeap(oldLine[0])
	for i := 1; i < len(matrix[0]); i++ {
		oldLine[i] = oldLine[i-1] ^ matrix[0][i]
		pushInHeap(oldLine[i])
	}
	// 从第1行开始
	for i := 1; i < len(matrix); i++ {
		newLine[0] = matrix[i][0]
		oldLine[0] = newLine[0] ^ oldLine[0]
		pushInHeap(oldLine[0])
		for j := 1; j < len(matrix[0]); j++ {
			newLine[j] = newLine[j-1] ^ matrix[i][j]
			oldLine[j] = newLine[j] ^ oldLine[j]
			pushInHeap(oldLine[j])
		}
	}
	// 找到第k个
	// 或者是用快速排序找第k个
	//var findK func(res []int, start int, end int, k int) int
	//findK = func(res []int, start int, end int, k int) int {
	//	pivot := res[start]
	//	i := start
	//	j := end - 1
	//	for {
	//		for j > i && res[j] > pivot {
	//			j--
	//		}
	//		if j == i {
	//			break
	//		}
	//		res[i] = res[j]
	//		i++
	//		for j > i && res[i] <= pivot {
	//			i++
	//		}
	//		if i == j {
	//			break
	//		}
	//		res[j] = res[i]
	//		j--
	//	}
	//	res[i] = pivot
	//	if i-start+1 == k {
	//		return pivot
	//	} else if i-start+1 > k {
	//		return findK(res, start, i, k)
	//	} else {
	//		return findK(res, i+1, end, k+start-i-1)
	//	}
	//}
	//return findK(res, 0, len(res), len(res)-k+1)
	sort.Ints(res)
	return res[len(res)-k]
}
