package question

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}
	rlen, clen := len(matrix), len(matrix[0])
	res := make([]int, rlen*clen)
	start := 0
	i, j, k := 0, 0, 0
	for {
		for ; j < clen; j++ { //向左
			res[k] = matrix[i][j]
			k++
		}
		j--
		i++
		if i >= rlen {
			break
		}
		for ; i < rlen; i++ { //跳过一个
			res[k] = matrix[i][j]
			k++
		}
		i--
		j--
		if j < start {
			break
		}
		for ; j >= start; j-- {
			res[k] = matrix[i][j]
			k++
		}
		j++
		i--
		if i < start {
			break
		}
		for ; i > start; i-- {
			res[k] = matrix[i][j]
			k++
		}

		if start == rlen-2 || start == rlen-1 {
			return res

		}
		start++
		clen--
		rlen--
		i = start
		j = start
	}
	return res
}
