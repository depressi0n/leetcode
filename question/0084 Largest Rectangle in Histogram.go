package question

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。

func largestRectangleArea(heights []int) int {
	return largestRectangleAreaCore3(heights)
}

// 枚举高  也可以枚举宽
// 朴素思想：以每根柱子为中心往两边尽可能扩散，得到的面积与最大值相比
// 整体表现为两轮，一轮从左往右计算，一轮从右往左计算
// 超时
func largestRectangleAreaCore1(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	Max := make([]int, len(heights))
	var left, right int
	//从左往右
	for left = 0; left < len(heights); left++ {
		right = left
		for right < len(heights) && heights[right] >= heights[left] {
			right++
		}
		Max[left] = (right - left) * heights[left]
	}
	//从右往左
	for right = len(heights) - 1; right >= 0; right-- {
		left = right
		for left >= 0 && heights[left] >= heights[right] {
			left--
		}
		Max[right] += (right - left - 1) * heights[right]
	}
	res := Max[0]
	for i := 1; i < len(heights); i++ {
		if Max[i] > res {
			res = Max[i]
		}
	}
	return res
}

//单调栈的思想：存放从左往右的非递减元素，然后通过二分的思想找到两边的值
// 但是！在计算left和right的过程中，可以直接找左右边界
func largestRectangleAreaCore2(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))
	monoStack := make([]int, 0, len(heights))
	for i := 0; i < len(heights); i++ {
		// 找到左边第一个小于当前高度的值即非递减栈的维护
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		// 如果为空则加上-1作为哨兵
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	// 重置，计算右边界
	monoStack = make([]int, 0, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] > heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	var res int
	for i := 0; i < len(heights); i++ {
		res = max(res, (right[i]-left[i]-1)*heights[i])
	}
	return res
}

// TODO:常数优化，仅遍历一次
// 常数优化，当位置i被弹出栈时，说明此时遍历的位置的高度小于等于height[i]，并且在二者之间没有其他高度小于等于height[i]
// 以上说明，当前位置是弹出位置的右边界，而一根柱子左侧且最近小于其高度的柱子 => 确实无法求出正确的右边界，但对答案没有影响
// 遍历结束后，栈中仍旧会留下一些值，这些值对应对右边界就是最后一个位置len(height)，依次出栈并更新右边界
func largestRectangleAreaCore3(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		right[i] = len(heights)
	}
	monoStack := make([]int, 0, len(heights))
	for i := 0; i < len(heights); i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	res := 0
	for i := 0; i < len(heights); i++ {
		res = max(res, (right[i]-left[i]-1)*heights[i])
	}
	return res
}
