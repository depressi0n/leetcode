package question

import (
	"container/heap"
)

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
	return IsPossibleCore2(target)
}
func IsPossibleCore(target []int) bool {
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

// 核心思想：从前往后的变化可以反向推导，所以从最终状态出发
// 不妨假定为a[0] .... a[m-1]
// 其中最大值为s，s>=a[i]，是上一次的和 此时的和为sum
// 根据题设，必然用上一次的和来替换某一个值，所以s为上一次的和
// 则满足关系sum = s + s -a[k] >= s
// 即 上次被替换的值为a[k] = 2*s - sum < s
// 按照这个操作，可以一直往上替换，如果最终的结果是全部为1，那么说明成立

// 这里其实存在一个优化，如果中间有几次替换都是用和来替换最大值，那么这段将形成一个等差数列
// 所以，可以一次性直接将数组变化到那个位置，而不是循环这么多次

func IsPossibleCore2(target []int) bool {
	if len(target) < 1 {
		return false
	}
	// 长度为1，则只能为1
	if len(target) == 1 && target[0] != 1 {
		return false
	}
	// 求出当前的和
	sum := target[0]
	for i := 1; i < len(target); i++ {
		sum += target[i]
	}
	// 堆排序，找到最大的和次大的
	// 最大是h[0]，次大或h[1]或h[2]
	h := IntHeap(target)
	heap.Init(&h)
	for {
		// 判断是否结束
		if h[0] == 1 {
			return true
		}
		// 找到上一次被替换的值，这里的h[0] == sum 是防止下面除0设置的
		if 2*h[0]-sum < 1 || h[0] == sum {
			return false
		}

		// 如果每次都是替换的最大值，那么这里会形成一个等差数列，我们需要找到次大值，将最大值直接降低到等于（在次大值为1的情况下成立）或者小于次大值
		// 等差数列的公差是sum-h[0]
		second := h[1]
		if 2 < len(h) && second < h[2] {
			second = h[2]
		}
		var k int
		if second == 1 {
			// 这时候可以将h[0]降低到和second相等
			// h[0] - k(sum-h[0]) <=second  求最小整数k
			// 即 h[0] - k(sum-h[0]) <= second 如果当时的k刚好为整数，那么说明
			k = (h[0] - second) / (sum - h[0])
		} else {
			// 此时要降低到second以下
			// h[0]-k(sum-h[0])<second 求最小整数k
			k = (h[0]-second)/(sum-h[0]) + 1
		}
		//h[0], sum = 2*h[0]-sum, h[0]
		h[0], sum = h[0]-k*(sum-h[0]), sum-k*(sum-h[0])
		heap.Fix(&h, 0)
	}
}
