package question

func reverseString(s []byte) {
	// 逆序即可
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return
}
