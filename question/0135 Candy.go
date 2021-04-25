package question

func candy(ratings []int) int {
	own := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		own[i] = 1
	}
	// left to right
	for i := 0; i < len(own)-1; i++ {
		if ratings[i+1] > ratings[i] && own[i+1] <= own[i] {
			own[i+1]++
		}
	}
	// right to left
	for i := len(own) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && own[i-1] <= own[i] {
			own[i-1]++
		}
	}

	var res int
	for i := 0; i < len(own); i++ {
		res += own[i]
	}
	return res
}

//升降坡的思想
