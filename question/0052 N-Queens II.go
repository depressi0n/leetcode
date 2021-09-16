package question

// n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

func totalNQueens(n int) int {
	return totalNQueensCore(n)
}
func totalNQueensCore(n int)int{
	res:=0
	// 用bit set来记录
	row := make([]int, n)
	used := 0
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			res++
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