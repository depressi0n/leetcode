package question

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}
	s, e := 0, len(numbers)-1
	for s < e {
		tmp := numbers[s] + numbers[e]
		if tmp < target {
			s++
		} else if tmp > target {
			e--
		} else {
			return []int{s + 1, e + 1}
		}
	}
	return []int{}
}
