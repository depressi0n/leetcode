package question

// 给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。
//你可以对一个单词进行如下三种操作：
//插入一个字符
//删除一个字符
//替换一个字符

func minDistance(word1 string, word2 string) int {
	return minDistanceCore(word1, word2)
}

//主要思想：
//（1）在word1中删除一个字母和在word2中插入一个字母是等价的，反之亦然
//（2）在word1中替换一个字母和在word2中替换一个字符是等价的
// 最后的情况：
//（1）在word1中插入
//（2）在word2中插入
//（3）在word1中替换
// 并且改变顺序不影响次数，这一点才是最关键的
func minDistanceCore(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	// 处理长短不相同的情况
	var long, short string
	if len(word1) < len(word2) {
		long = word2
		short = word1
	} else {
		long = word1
		short = word2
	}
	minInThree := func(a int, b int, c int) int {
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
	// dp[i][j] 定义为long[:i]和short[:j]的编辑距离
	// dp[i][j] = 以下几种情况的最小值
	// （1）long[i-1] == short[j-1]  => min(dp[i-1][j-1],dp[i][j-1]+1,dp[i-1][j]+1)
	// （2）long[i-1] != short[j-1] => min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
	// 从上述转移方程来看，只依赖于本行之前的数和当前元素的左上元素
	// 可以降为一维处理，并用一个变量存储左上元素的值
	dp := make([]int, len(short)+1)
	// 初始化
	// 不妨假设有个第0层，初始化时就初始化这个第0层，这种哨兵挺有用的！！！！
	// 其实是对空字符串编辑距离的抽象化
	for i := 0; i < len(short)+1; i++ {
		dp[i] = i
	}
	cache := dp[0] //保存dp[i-1][j-1]
	//更新
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