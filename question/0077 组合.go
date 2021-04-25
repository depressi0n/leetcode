package question

//TODO:看懂节省空间的方法，二进制next怎么转换的
func combine(n int, k int) [][]int {
	if k == 0 || n < k {
		return nil
	}
	if n == k {
		res := []int{}
		for i := 0; i < n; i++ {
			res = append(res, i+1)
		}
		return [][]int{res}
	}
	var res [][]int
	for i := 1; i <= n-k+1; i++ {
		dfs_combine(n, k, i, []int{i}, &res)
	}
	return res
}
func dfs_combine(n int, k int, pos int, cur []int, res *[][]int) {
	if len(cur) == k {
		curCopy := make([]int, k)
		for i := 0; i < k; i++ {
			curCopy[i] = cur[i]
		}
		*res = append(*res, curCopy)
		return
	}
	//剪枝，如果剩下的数不够了，就不进了
	if len(cur)+n-pos < k {
		return
	}
	for i := pos + 1; i <= n; i++ {
		cur = append(cur, i)
		dfs_combine(n, k, i, cur, res)
		cur = cur[:len(cur)-1]
	}
}

func combine1(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1) //哨兵，保证到第k个就一定会停止下面第二层循环

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
func combine2(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 这种构造挺巧妙的...
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
