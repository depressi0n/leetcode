package question

func Answer1104(label int) []int {
	return pathInZigZagTree(label)
}
func pathInZigZagTree(label int) []int {
	//思路：首先找出这个位置上本来的值，然后通过本来的值确定本来的路径，再替换回去
	// 找本来的值
	//           1
	//    3(2)            2(3)
	//  4    5      6     7
	// 8(15) 9(14) 10(13) 11(12)  12(11) 13(10) 14(9) 15(8)
	// 首先确定这个值在奇数行还是偶数行
	rowNum := 0
	for 1<<rowNum <= label {
		rowNum++
	}
	res := make([]int, rowNum)
	if rowNum&1 == 0 {
		// 从右往左，换成原来应该的值
		res[len(res)-1] = 1<<rowNum + 1<<(rowNum-1) - 1 - label
	} else {
		res[len(res)-1] = label
	}
	for i := len(res) - 2; i >= 0; i-- {
		res[i] = res[i+1] / 2
	}
	for i := 1; i < len(res); i++ {
		if i&1 == 1 {
			res[i] = 1<<i + 1<<(i+1) - 1 - res[i]
		}
	}
	return res
}
