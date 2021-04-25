package question

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// No two characters may map to the same character but a character may map to itself.
	m1 := make(map[byte]byte) //t[i]->s[i]
	m2 := make(map[byte]byte) //s[i]->t[i]
	for i := 0; i < len(s); i++ {
		// 先检查是否存在之前的映射s[i]->t[i],如果存在，则相应的转化即可
		if _, ok := m2[s[i]]; ok {
			if t[i] != m2[s[i]] { //同一个字母被映射为不同字母
				return false
			}
		}
		// 如果不存在之前的映射，则建立新的映射，但要检查是否是单射
		m2[s[i]] = t[i]
		if _, ok := m1[t[i]]; ok { //不同字母被映射为同一个字母
			if m1[t[i]] != s[i] {
				return false
			}
		}
		m1[t[i]] = s[i]
	}
	return true
}
