package question

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	i, j, start, end := 0, 0, 0, n
	cur := 1
	for start+1 != end || start+2 != end {
		for j < end {
			res[i][j] = cur
			cur++
			j++
		}
		j--
		i++
		for i < end {
			res[i][j] = cur
			cur++
			i++
		}
		i--
		j--
		for j >= start {
			res[i][j] = cur
			cur++
			j--
		}
		j++
		i--
		for i > start {
			res[i][j] = cur
			cur++
			i--
		}
		start++
		end--
		if start >= end {
			return res
		}
		i = start
		j = start
	}
	return res
}
