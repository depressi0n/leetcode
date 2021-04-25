package question

func MinSwaps(grid [][]int) int {
	// 矩阵的主对角线往上全部是零称之为有效
	// 每一步可以交换两个相邻的行
	// 求最小步数，如果不能则返回-1

	// 题目意思是上三角必须是全0
	// 检查每一行上最右边的1的位置，用这个位置来确定这一行的最终位置
	// 如果检查后发现不行，则直接返回-1
	location := make(map[int]int)
	mostRightOne := make([]int, len(grid))
	for i := 0; i < len(grid); i++ {
		mostRightOne[i] = len(grid[i]) - 1
		for mostRightOne[i] > 0 && grid[i][mostRightOne[i]] == 0 {
			mostRightOne[i]--
		}
		if cnt, ok := location[mostRightOne[i]]; !ok {
			location[mostRightOne[i]] = 1
		} else {
			location[mostRightOne[i]] = cnt + 1
		}
	}
	// 对mostRight检查
	cnt := 0
	for i := 0; i < len(grid); i++ {
		cnt += location[i]
		if cnt < i+1 {
			return -1
		}
	}
	// 模拟交换过程，类似于冒泡？但是指标不一样，只需要保证mostRightOne[i]<=i即可
	// 比如 0 0 0 0 不会有问题
	// 0 1 1 1 也不会有问题
	// 1 1 1 0 将返回3
	// 利用双指针完成交换的模拟
	start := 0
	step := 0
	for start < len(mostRightOne) {
		for start < len(mostRightOne) && mostRightOne[start] <= start {
			start++
		}
		if start == len(mostRightOne) {
			return step
		}
		// 必须往后交换，在后面找到第一个不大于start的值出来，往前交换
		end := start + 1
		for ; end < len(mostRightOne) && mostRightOne[end] > start; end++ {
		}
		// 以防万一，上面的检查保证了这里不可能出现这种情况
		if end == len(mostRightOne) {
			return -1
		}
		for i := end; i > start; i-- {
			if mostRightOne[i] != mostRightOne[i-1] { //相等交换是没有意义的
				mostRightOne[i], mostRightOne[i-1] = mostRightOne[i-1], mostRightOne[i]
				step++
			}
		}
		start++
	}
	return step
}
