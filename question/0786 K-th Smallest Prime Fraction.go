package question

import "container/heap"

// 给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干 素数 组成，且其中所有整数互不相同。
//对于每对满足 0 < i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。
//那么第k个最小的分数是多少呢? 以长度为 2 的整数数组返回你的答案, 这里answer[0] == arr[i]且answer[1] == arr[j] 。

func kthSmallestPrimeFraction(arr []int, k int) []int {
	return kthSmallestPrimeFractionCore2(arr, k)
}

// 主要思想：素数保证了作为分式时不会约分，数组中的数保证了严格单调递增的规则，那么计算第k个小的分数
// 可以通过一个最大堆来维护一个最小的k个值，从而将答案保留在结果中
type fraction struct {
	numerator, denominator int
}
type h []fraction

func (h h) Len() int {
	return len(h)
}

func (h h) Less(i, j int) bool {
	return h[i].numerator*h[j].denominator >= h[i].denominator*h[j].numerator
}

func (h h) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *h) Push(x interface{}) {
	*h = append(*h, x.(fraction))
}

func (h *h) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallestPrimeFractionCore1(arr []int, k int) []int {
	help := h([]fraction{})
	heap.Init(&help)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if len(help) < k {
				heap.Push(&help, fraction{
					numerator:   arr[i],
					denominator: arr[j],
				})
			} else if help[0].numerator*arr[j] > help[0].denominator*arr[i] {
				heap.Pop(&help)
				heap.Push(&help, fraction{
					numerator:   arr[i],
					denominator: arr[j],
				})
			}
		}
	}
	return []int{help[0].numerator, help[0].denominator}
}

// 二分查找+双指针：借助最终的答案一定在[0，1)之间，进行二分查找！！
// 我们可以随便猜测一个实数α 如果恰好有 k 个素数分数小于 α，那么这 k 个素数分数中最大的那个就是第 k 个最小的素数分数。
//对于 α，我们如何求出有多少个小于 α 的素数分数呢？我们可以使用双指针来求出答案：
//   我们使用一个指针 j 指向分母，这个指针每次会向右移动一个位置，表示枚举不同的分母；
//   我们再使用一个指针 i 指向分子，这个指针会不断向右移动，并且保证 arr[i] / arr[j]< α 一直成立。当指针 /i 无法移动时，
//  		我们就可以知道 arr[0],⋯,arr[i] 都可以作为分子，但 arr[i+1] 及以后的元素都不可以，即分母为 arr[j] 并且小于 α 的素数分数有 i+1 个。
//	 在 j 向右移动的过程中，我们把每一个 j 对应的 i+1 都加入答案。这样在双指针的过程完成后，我们就可以得到有多少个小于 α 的素数分数了。
// 如果我们得到的答案恰好等于 k，那么我们再进行一遍上面的过程，求出所有 arr[i]/arr[j]中的最大值即为第 k 个最小的素数分数。
//
//但如果答案小于 k，这说明我们猜测的 α 太小了，我们需要增加它的值；如果答案大于 k，这说明我们猜测的 /α 太大了，我们需要减少它的值。
//这就提示我们使用二分查找来找到合适的 α。二分查找的上界为 1，下界为 0 。在二分查找的每一步中，我们取上下界区间的中点作为 α α，并计算小于
//α 的素数分数的个数，并根据这个值来调整二分查找的上界或下界。
// 最坏情况下多少步的分析：任意两个素数组成的分数之差不会小于1/(3*10^4)^2
func kthSmallestPrimeFractionCore2(arr []int, k int) []int {
	n := len(arr)
	// 必须使用浮点数
	left, right := 0., 1.
	for {
		mid := (left + right) / 2
		i, count := -1, 0
		// 记录最大的分数
		x, y := 0, 1
		for j := 1; j < n; j++ {
			for float64(arr[i+1]) / float64(arr[j]) < mid {
				i++
				if arr[i]*y > arr[j]*x {
					x, y = arr[i], arr[j]
				}
			}
			count += i + 1
		}

		if count == k {
			return []int{x, y}
		}
		if count < k {
			left = mid
		} else {
			right = mid
		}
	}
}
