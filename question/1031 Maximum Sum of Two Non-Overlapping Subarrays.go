package question

import "math"

func MaxSumTwoNoOverlap(A []int, L int, M int) int {
	// f[i]表示在i之前，长度为L,M的子数组的最大和
	// g[i]表示以j之后，长度为L,M的子数组的最大和
	if L+M == len(A) {
		return func() int {
			res := 0
			for i := 0; i < len(A); i++ {
				res += A[i]
			}
			return res
		}()
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	order := func(L, M int) int {
		f := make([]int, len(A))
		sum := 0
		for i := 0; i < L; i++ {
			f[i] = -1
			sum += A[i]
		}
		f[L-1] = sum
		for i := L; i < len(A); i++ {
			sum = sum + A[i] - A[i-L]
			f[i] = max(f[i-1], sum)
		}
		sum = 0
		g := make([]int, len(A))
		for i := len(A) - 1; i >= len(A)-M; i-- {
			g[i] = -1
			sum += A[i]
		}
		g[len(A)-M] = sum
		for i := len(A) - M - 1; i >= 0; i-- {
			sum = sum - A[i+M] + A[i]
			g[i] = max(g[i+1], sum)
		}
		res := math.MinInt32
		for i := L - 1; i < len(A)-M; i++ {
			res = max(res, f[i]+g[i+1])
		}
		return res
	}
	return max(order(L, M), order(M, L))
}

// TODO 就地算法，考虑一下前缀和和滑动窗口的不同之处以及适应什么样的设置情况
// 利用前缀和可以省空间，时间可以在常数上优化
