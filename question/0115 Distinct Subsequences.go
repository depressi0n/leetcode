package question

func numDistinct1(s string, t string) int {
	// 思路
	// 要记录每个字符出现后，接下来的字符的出现的次数，然后加起来
	// 对t中的字符进行选取，记录出现的位置
	//m:=make(map[byte]int) // t中字符出现的位子
	n := make([][]int, len(t))

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
	// 如果出现重复的怎么办？应该不会出问题，因为有序号的限制
	// 从下往上或者从上往下做一个动态规划
	dp := make([]int, len(s))
	//初始化
	for i := 0; i < len(n[0]); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(t); i-- {
		//和第i-1层比较,比其小的值有多少个
		tmp := make([]int, len(n[i]))
		for j := 0; j < len(n[i]); j++ {
			//可以用二分法优化一下
			for k := 0; k < len(n[i-1]); k++ {
				if n[i-1][k] < n[i][j] {
					tmp[j]++
				} else {
					break
				}
			}
		}
		for i := 0; i < len(dp); i++ {
			if i < len(tmp) {
				dp[i] = tmp[i]
			} else {
				dp[i] = 0
			}
		}
	}
	for i := 1; i < len(dp); i++ {
		dp[0] += dp[i]
	}
	return dp[0]
}
func numDistinct2(s string, t string) int {
	// 看了答案以后，可以优化到极致
	// 因为每个dp都只依赖于左上角的值，可以将上面n数组和下面更新合并起来
	// 选与不选的问题，从后往前选
	// 如果匹配上,可以选也可以不选
	// 如果没有匹配上，只能够不选
	dp := make([]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		dp[i] = 1
	}
	tmp := 1
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				tmp, dp[j] = dp[j], tmp+dp[j-1]
			} else {
				tmp, dp[j] = dp[j], dp[j-1]
			}
		}
		tmp = 0
	}
	return dp[len(dp)-1]

}

//TODO:从后面往前更新的理由是什么或者是状态转移方程
func numDistinct3(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		for j := len(t); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[len(t)]
}
