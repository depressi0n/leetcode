package question

// 题目描述：给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

func trap(height []int) int {
	return trapCore4(height)
}

// 暴力：求出两边的最大值
func trapCore1(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		leftMax := i
		left := i - 1
		for ; left >= 0; left-- {
			if height[left] > height[leftMax] {
				leftMax = left
			}
		}
		rightMax := i
		right := i + 1
		for ; right < len(height); right++ {
			if height[right] > height[rightMax] {
				rightMax = right
			}
		}
		res += min(height[leftMax], height[rightMax]) - height[i]
	}
	return res
}

// 动态规划：left[i]表示height[:i+1]的最大值，right[i]表示右侧的最大值
func trapCore2(height []int) int {
	res := 0
	left := make([]int, len(height))
	left[0] = 0
	for i := 1; i < len(height); i++ {
		if height[i] > height[left[i-1]] {
			left[i] = i
		} else {
			left[i] = left[i-1]
		}
	}
	right := make([]int, len(height))
	right[len(height)-1] = len(height) - 1
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > height[right[i+1]] {
			right[i] = i
		} else {
			right[i] = right[i+1]
		}
	}
	for i := 1; i < len(height)-1; i++ {
		res += min(height[left[i]], height[right[i]]) - height[i]
	}
	return res
}

// 单调栈的思想：从左往右扫描，非递增维护高度值，如果当前值比栈顶小则插入，如果比当前值大，则出栈，积累雨水
// 扫描完成后，结束
func trapCore3(height []int) int {
	res := 0
	s := make([]int, 0, len(height))
	for i := 0; i < len(height); i++ {
		if len(s) == 0 || height[i] <= height[s[len(s)-1]] {
			s = append(s, i)
		} else {
			// 出栈
			for len(s) > 0 && height[s[len(s)-1]] < height[i] {
				if len(s) > 1 {
					res += (min(height[s[len(s)-2]], height[i]) - height[s[len(s)-1]]) * (i - s[len(s)-2] - 1)
				}
				s = s[:len(s)-1]
			}
			s = append(s, i)
		}
	}
	return res
}

// 终极方案：双指针,left和right分别维护当前左边和右边遇见的最大高度
func trapCore4(height []int) int {
	//left := 0
	//for left+1 < len(height) && height[left] <= height[left+1] {
	//	left++
	//}
	//if left == len(height) {
	//	return 0
	//}
	//right := len(height) - 1
	//for right-1 >= 0 && height[right] <= height[right-1] {
	//	right--
	//}
	//if right == 0 {
	//	return 0
	//}
	res := 0
	// 此时 left和right 分别指向一个高度
	left, right := 0, len(height)-1
	i, j := left+1, right-1
	for i <= j {
		if height[left] < height[right] {
			if height[i] < height[left] {
				res += min(height[left], height[right]) - height[i]
			} else {
				left = i
			}
			i++
		} else {
			if height[j] < height[right] {
				res += min(height[left], height[right]) - height[j]
			} else {
				right = j
			}
			j--
		}
	}
	return res
}
