package question

import "strings"

// 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
// 给定 n 和 k，返回第 k 个排列。
//TODO:不实用next permutation，而是用一个位置确定后，将是第几个全排列
func getPermutation(n int, k int) string {
	return getPermutationCore2(n, k)
}
func getPermutationCore1(n int, k int) string {

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
func getPermutationCore2(n int, k int) string {
	// cnt[i] 表示i+1个数的排列有多少个
	cnt := make([]int, n)
	cnt[0] = 1
	for i := 1; i < n; i++ {
		cnt[i] = cnt[i-1] * (i + 1)
	}
	b := strings.Builder{}
	start := 1
	need := k
	m := map[int]struct{}{}
	for need!=1{
		start = 1
		for cnt[n-len(m)-2] < need {
			if _, ok := m[start]; !ok {
				need -= cnt[n-len(m)-2]
			}
			start++
		}
		for ;start<=n;start++{
			if _,ok:=m[start];!ok{
				break
			}
		}
		b.WriteByte(byte(start + '0'))
		m[start] = struct{}{}
	}
	for start = 1; start <= n; start++ {
		if _, ok := m[start]; !ok {
			b.WriteByte(byte(start + '0'))
		}
	}
	return b.String()
}
