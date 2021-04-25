package question

func kthFactor(n int, k int) int {
	// 首先找出所有的因子，只需要找到根号n之前的就可以了，因为因子是成对出现的
	// 从前开始数，发现没有则反向接着数，这样数到第k个就是答案
	// 也可以提前先判断这一段的因子个数
	mid := 1
	for {
		if mid*mid <= n {
			mid++
		} else {
			break
		}
	}
	factors := make([]int, 0, mid)
	for i := 0; i < mid; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			// 可以把下面小于的判断放在这里，不过要进行很多次判断才行
			if k == len(factors) {
				return i
			}
		}
	}
	// 考虑最后一个数的平方就是n的情况
	if k <= len(factors) {
		return factors[k-1]
	} else if k <= 2*len(factors) {
		if factors[len(factors)-1]*factors[len(factors)-1] == n {
			return -1
		} else {
			return n / factors[2*len(factors)-k]
		}
	} else {
		return -1
	}
}
