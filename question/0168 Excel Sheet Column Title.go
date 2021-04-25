package question

func convertToTitle(n int) string {
	res := make([]byte, 0)
	for n != 0 {
		n--
		t := n % 26
		n = n / 26
		res = append([]byte{byte(t + 'A')}, res[:]...)
	}
	return string(res)
}
