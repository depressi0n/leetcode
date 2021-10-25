package question

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。

func getRow(rowIndex int) []int {
	return getRowCore(rowIndex)
}
func getRowCore(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	res := make([]int, rowIndex+1)
	res[0] = 1
	tmp:=1
	for i := 1; i <= rowIndex; i++ {
		for j := 1; j < i; j++ {
			tmp, res[j] = res[j], res[j]+tmp
		}
		res[i] = 1
	}
	return res
}
