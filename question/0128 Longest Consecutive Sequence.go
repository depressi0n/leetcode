package question

//TODO:看懂并查集
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	max := 1
	for k, _ := range m {
		start := k - 1
		end := k + 1
		for {
			if _, ok := m[start]; ok {
				delete(m, start)
				start--
			} else {
				break
			}
		}

		for {
			if _, ok := m[end]; ok {
				delete(m, end)
				end++
			} else {
				break
			}
		}
		delete(m, k)
		if end-start-1 > max {
			max = end - start - 1
		}
	}
	return max
}
