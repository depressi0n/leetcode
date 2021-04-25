package question

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IsPossible(target []int) bool {
	if len(target) == 1 && target[0] == 1 {
		return true
	}
	if len(target) == 1 {
		return false
	}

	var sum int
	for i := 0; i < len(target); i++ {
		sum += target[i]
	}
	h := IntHeap(target)
	heap.Init(&h)
	// 每次替换用值是所有值的和，被替换的值是2*s-sum
	for {
		// 最大值是h[0]
		// 次大值是h[1]或者h[2]
		max := h[0]
		unit := sum - max
		if max == 1 {
			return true
		}
		if 2*max-sum < 1 {
			return false
		}
		secondMax := h[1]
		if 2 < len(target) && secondMax < h[2] {
			secondMax = h[2]
		}
		var k int
		if secondMax == 1 {
			k = (sum - secondMax - 1) / unit
		} else {
			k = (max-secondMax)/unit + 1
		}
		sum -= k * unit
		max -= k * unit
		if max <= 0 {
			return false
		}
		h[0] = max
		heap.Fix(&h, 0)
	}
}
