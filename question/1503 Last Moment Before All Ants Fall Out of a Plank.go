package question

func getLastMoment(n int, left []int, right []int) int {
	res := -1
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < len(left); i++ {
		res = max(res, left[i])
	}
	for i := 0; i < len(right); i++ {
		res = max(res, n-right[i])
	}
	return res
}
