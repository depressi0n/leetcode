package question

// 给定一个字符串 s 和一个字符串 t ，计算在s的子序列中t出现的个数。
// 字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）
// 题目数据保证答案符合 32 位带符号整数范围。

func numDistinct(s string, t string) int {
	return numDistinctCore3(s, t)
}

// 思路
// 要记录每个字符出现后，接下来的字符的出现的次数，然后加起来
// 这种方法肯定会超时，复杂度为len(s)*len(t)
func numDistinctCore(s string, t string) int {
	// 对t中的字符进行选取，记录出现的位置
	n := make([][]int, len(t))
	// 预处理，t中每个出现在s中的位置
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				if n[j] == nil {
					n[j] = make([]int, 0)
				}
				n[j] = append(n[j], i)
			}
		}
	}
	// 初始化，从最后一个字母出现的位置开始，从下往上做动态规划
	dp := make([][]int, len(t))
	dp[len(t)-1] = make([]int, len(n[len(t)-1]))
	for i := 0; i < len(dp[len(t)-1]); i++ {
		dp[len(t)-1][i] = 1
	}
	// 从转移方程可以看到，每次更新只依赖于下一行，所以可节省空间，使用两个数组进行更新
	// 但是空间复杂依旧不变，因为上面的预处理需要，但是根据二者空间的大小，发现二者可以直接结合起来！！
	for i := len(t) - 2; i >= 0; i-- {
		dp[i] = make([]int, len(n[i]))
		for j := 0; j < len(n[i]); j++ {
			for k := 0; k < len(n[i+1]); k++ {
				if n[i+1][k] > n[i][j] {
					dp[i][j] += dp[i+1][k]
				}
			}
		}
	}
	res := 0
	for i := 0; i < len(dp[0]); i++ {
		res += dp[0][i]
	}
	return res
}

// 优化：
// 将上面n和dp的使用直接结合起来，一次完成
// 原有的预处理过程分段完成
// dp[i][j]表示t[:i]出现在s[:j]中的次数
// 初始化时dp[0]均为1
// 转移方程为dp[i][j]=dp[i-1][j-1] if t[i-1] == s[j-1] else dp[i][j-1]
// 由此可见转移方程只取决与上一行的数和当前行之前的数，可优化为一维数组+一个值的形式
func numDistinctCore2(s string, t string) int {
	dp := make([]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		dp[i] = 1
	}
	tmp := 1 // 这个值可以使用dp[0]，因为每次的起始位置都没变过
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				tmp, dp[j] = dp[j], tmp+dp[j-1]
			} else {
				tmp, dp[j] = dp[j], dp[j-1]
			}
		}
		tmp = 0 // 下一轮起始时置0，因为dp[i][0]表示s为空的时候，不可能成立
	}
	return dp[len(dp)-1]

}

// 主要思想：统计每个字母后出现的下一个字母的个数，最后通过加法解决
// 本质上脱胎于上面的思路，只是这是从后往前更新，从后往前更新的理由是可以利用之前的结果直接完成累加的过程
// dp[i][j] 表示 s[:i]中出现在t[:j]的个数（注意和上面的下标是不一样的！！！！）
// 从后往前更新时，上一层的结果直接表示s[:i-1] 中出现 t[:j]的个数，所以要累加进去
// dp[i][j] = dp[i-1][j]+dp[i-1][j-1] if s[i-1]==t[j-1] else dp[i-1][j]
// 改写成一维 dp[i] = dp[i]+dp[i-1] if s[i-1]==t[j-1]
func numDistinctCore3(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1
	// 从后往前更新的目的是累加直接在此过程中一起完成
	// 这是因为：在更新i时之前，dp[i]记录的是t[j]出现在s[:i-1]中出现的次数
	// 所以要从后往前更新
	for i := 1; i <= len(s); i++ {
		// 每次都要更新一次dp，其实就是在看s中出现在t[i:]的个数是多少
		for j := len(t); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(t)]
}
// dp[i][j] 可定义为 s[j:]中出现t[i:]的组合数
// dp[i][j] = dp[i][j+1] + dp[i+1][j+1] if t[i]==s[j] else dp[i][j+1]
// 同样地，dp[i][j] 可定义为s[i:]中出现t[j:]的组合数
// dp[i][j]=dp[i+1][j] +dp[i+1][j+1] if s[i] == t[j] else dp[i+1][j]