package question

func findRepeatedDnaSequences(s string) []string {
	//找出现次数大于一次的重复子串
	res := make([]string, 0)
	m := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		if _, ok := m[s[i:i+10]]; ok {
			m[s[i:i+10]] = m[s[i:i+10]] + 1
			continue
		}
		m[s[i:i+10]] = 1
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

//TODO:有两种很棒的算法
