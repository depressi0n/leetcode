package question

// 朴素方法
func largestRectangleArea_textbook(heights []int) int {
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
	max := Max[0]
	for i := 1; i < len(heights); i++ {
		if Max[i] > max {
			max = Max[i]
		}
	}
	return max
}

//单调栈
//TODO:看懂那个常数优化
func largestRectangleArea(heights []int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	left, right := make([]int, len(heights)), make([]int, len(heights))
	monoStack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = monoStack[:0]
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
