package question

//TODO:不实用next permutation，而是用一个位置确定后，将是第几个全排列
func getPermutation(n int, k int) string {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = '0' + i + 1
	}
	for k > 1 {
		res = nextPermutationUnique(res)
		k--
	}
	resb := ""
	for i := 0; i < n; i++ {
		resb += string(byte(res[i]))
	}
	return resb
}
