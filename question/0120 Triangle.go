package question

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	res := make([]int, len(triangle))
	for i := 0; i < len(res); i++ {
		res[i] = triangle[len(triangle)-1][i]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if res[j] < res[j+1] {
				res[j] = res[j] + triangle[i][j]
			} else {
				res[j] = res[j+1] + triangle[i][j]
			}
		}
	}
	return res[0]
}
