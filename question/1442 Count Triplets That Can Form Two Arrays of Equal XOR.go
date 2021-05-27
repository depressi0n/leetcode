package question

func CountTriplets(arr []int) int {
	// arr[i] ^ ... ^arr[j-1]^arr[j]^...^arr[k] = 0
	// prefixXOR[i..k] = prefixXOR[k]-prefixXOR[i-1]
	// 0 <= i < j <=k
	//res := 0
	//m := make(map[int][]int)
	//prefixXOR := make([]int, len(arr)+1)
	//prefixXOR[0] = 0
	//m[prefixXOR[0]] = append(m[prefixXOR[0]], 0)
	//for i := 1; i < len(arr)+1; i++ {
	//	prefixXOR[i] = prefixXOR[i-1] ^ arr[i-1]
	//	// 在prefixXOR数组中找值相等的,或者遇见为0的,作为i和k，中间的任意的j都可以作为结果
	//	if _, ok := m[prefixXOR[i]]; ok {
	//		// [index+1,...,i]可以作划分
	//		for _, index := range m[prefixXOR[i]] {
	//			res += i - index -1
	//		}
	//	}
	//	m[prefixXOR[i]] = append(m[prefixXOR[i]], i)
	//}
	//return res
	// TODO_DONE: prefixXOR 数组可以用一个变量替代掉
	res := 0
	m := make(map[int][]int)
	//prefixXOR := make([]int, len(arr)+1)
	prefixXOR := 0
	m[prefixXOR] = append(m[prefixXOR], 0)
	for i := 1; i < len(arr)+1; i++ {
		prefixXOR = prefixXOR ^ arr[i-1]
		// 在prefixXOR数组中找值相等的,或者遇见为0的,作为i和k，中间的任意的j都可以作为结果
		if _, ok := m[prefixXOR]; ok {
			// [index+1,...,i]可以作划分
			for _, index := range m[prefixXOR] {
				res += i - index - 1
			}
		}
		m[prefixXOR] = append(m[prefixXOR], i)
	}
	return res
}
