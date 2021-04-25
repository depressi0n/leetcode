package question

import "fmt"

func generate(numRows int) [][]int {
	//杨辉三角
	if numRows <= 0 {
		return nil
	}
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	for i := 2; i <= numRows; i++ {
		tmp := make([]int, i)
		tmp[0] = 1
		tmp[i-1] = 1
		for j := 1; j < i-1; j++ {
			fmt.Println(i, j)
			tmp[j] = res[i-2][j-1] + res[i-2][j]
		}
		res = append(res, tmp)
	}
	return res
}
