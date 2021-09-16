package question

import (
	"reflect"
	"strings"
)

// n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。
//
//每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

func solveNQueens(n int) [][]string {
	return solveNQueensCore2(n)
}

// DFS + 回溯
func solveNQueensCore1(n int) [][]string {
	res := make([][]string, 0, n)
	// 用bit set来记录
	row := make([]int, n)
	used := 0
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			t := make([]string, 0, n)
			build := strings.Builder{}
			for i := 0; i < n; i++ {
				build.Reset()
				for j := 0; j < n; j++ {
					if row[i]&(1<<j)>>j == 1 {
						build.WriteByte('Q')
					} else {
						build.WriteByte('.')
					}
				}
				t = append(t, build.String())
			}
			res = append(res, t)
			return
		}
		for i := 0; i < n; i++ {
			if used&(1<<i)>>i == 1 {
				continue
			}

			// 斜线
			j := cur - 1
			for ; j >= 0 && i+j-cur >= 0 && row[j]&(1<<(i+j-cur))>>(i+j-cur) == 0; j-- {

			}
			if j != -1 && i+j-cur != -1 {
				continue
			}
			j = cur - 1
			for ; j >= 0 && i-j+cur < n && row[j]&(1<<(i-j+cur))>>(i-j+cur) == 0; j-- {

			}
			if j != -1 && (0 <= i-j+cur && i-j+cur < n) {
				continue
			}
			row[cur] |= 1 << i
			used |= 1 << i
			dfs(cur + 1)
			row[cur] ^= 1 << i
			used ^= 1 << i
		}
	}
	dfs(0)
	return res
}

// 全排列 + 判断
func solveNQueensCore2(n int) [][]string {
	nextPermutation := func(nums []int) []int {
		i := len(nums) - 2
		for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
			// pass
		}
		if i >= 0 {
			j := len(nums) - 1
			for ; j > i && nums[j] <= nums[i]; j-- {
				// pass
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		for j := i + 1; j < (i+len(nums)+1)/2; j++ {
			nums[j], nums[len(nums)+i-j] = nums[len(nums)+i-j], nums[j]
		}
		return nums
	}
	t := make([]int, n)
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = i + 1
		cur[i] = i + 1
	}
	res := make([][]string, 0, n)
	flag := false
	b := strings.Builder{}
I:
	for !flag || !reflect.DeepEqual(t, cur) {
		if !flag {
			flag = true
		}
		// 检查是否有斜线冲突
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if cur[i]-cur[j] == i-j || cur[i]-cur[j] == j-i {
					cur = nextPermutation(cur)
					continue I
				}
			}
		}
		tmp := make([]string, 0, n)
		for i := 0; i < n; i++ {
			b.Reset()
			k := 0

			for ; k < cur[i]-1; k++ {
				b.WriteByte('.')
			}
			b.WriteByte('Q')
			k++
			for ; k < n; k++ {
				b.WriteByte('.')
			}
			tmp = append(tmp, b.String())
		}
		res = append(res, tmp)
		cur = nextPermutation(cur)
	}
	return res
}
