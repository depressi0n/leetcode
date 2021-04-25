package question

func minDistance(word1 string, word2 string) int {
	//主要思想：
	//（1）在word1中删除一个字母和在word2中插入一个字母是等价的，反之亦然
	//（2）在word1中替换一个字母和在word2中替换一个字符是等价的
	// 最后的情况：
	//（1）在word1中插入
	//（2）在word2中插入
	//（3）在word1中替换
	// 并且改变顺序不影响次数，这一点才是最关键的
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	var long, short string
	if len(word1) < len(word2) {
		long = word2
		short = word1
	} else {
		long = word1
		short = word2
	}
	dp := make([]int, len(short)+1)
	//初始化
	//不妨假设有个第0层，初始化时就初始化这个第0层，这种哨兵挺有用的！！！！
	for i := 0; i < len(short)+1; i++ {
		dp[i] = i
	}
	cache := dp[0] //保存dp[i-1][j-1]
	//更新
	// TODO: 这里的更新策略可以学习一下
	for i := 0; i < len(long); i++ {
		cache = dp[0]
		dp[0] = i + 1
		for j := 0; j < len(short); j++ {
			if long[i] == short[j] {
				dp[j+1], cache = 1+minInThree(dp[j], cache-1, dp[j+1]), dp[j+1]
			} else {
				dp[j+1], cache = 1+minInThree(dp[j], cache, dp[j+1]), dp[j+1]
			}
		}
	}
	//返回
	return dp[len(dp)-1]
}

func minInThree(a int, b int, c int) int {
	//两次比较即可
	//假设a是最小的
	if b < a {
		a = b
	}
	if c < a {
		return c
	} else {
		return a
	}
}
