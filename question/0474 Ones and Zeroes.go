package question

import "strings"

func FindMaxForm(strs []string, m int, n int) int {
	// 目标：从strs中挑选string，要求个数尽可能多，但其中1的总个数不能超过m，0的总个数不能超过0
	// 基本思路：
	//（1）首先确定每个string中有多少个0和1
	//（2）利用某种算法（应该只能用dfs？）进行汇聚,但是dfs会超时
	//（3）记录一个最大值
	//contain0:=make([]int,len(strs))
	//for i := 0; i < len(strs); i++ {
	//	for j := 0; j < len(strs[i]); j++ {
	//		if strs[i][j]=='0'{
	//			contain0[i]++
	//		}
	//	}
	//}
	//res:=0
	//var dfs func(cur int,num int,cnt0 int,cnt1 int)
	//dfs= func(cur int,num int, cnt0 int,cnt1 int) {
	//	if cur == len(strs){
	//		return
	//	}
	//	// 剪枝
	//	if cnt1>n || cnt0 >m{
	//		return
	//	}
	//	if num+len(strs)-cur<res{
	//		return
	//	}
	//	//选择
	//	if cnt0+contain0[cur]<=m && cnt1+len(strs[cur])-contain0[cur]<=n{
	//		if res<num+1{
	//			res =num+1
	//		}
	//		dfs(cur+1,num+1,cnt0+contain0[cur],cnt1+len(strs[cur])-contain0[cur])
	//	}
	//	//不选择
	//	dfs(cur+1,num,cnt0,cnt1)
	//}
	//dfs(0,0,0,0)
	//return res
	// dp[i][j][k] 表示strs[0,i]中选择j个0和k个1的情况下最多可以得到的字符串数量
	// dp[0][j][k]=0
	// dp[i][j][k]= dp[i-1][j][k]  or max{dp[i-1][j][k],dp[i-1][j-contain0[i]][k-len(strs[i])+contain0[i]]+1}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zeros && k >= ones {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

//TODO
//matrix := make([][]int, m+1)
//	// 初始化记忆矩阵
//	for i, _ := range matrix {
//		matrix[i] = make([]int, n+1)
//		for j, _ := range matrix[i] {
//			matrix[i][j] = -1
//		}
//	}
//	matrix[0][0] = 0
//
//	for _, str := range strs {
//		// 计算 0 1 数
//		costOne := 0
//		costZero := 0
//		for _, c := range str {
//			if c == '1' {
//				costOne++
//			} else {
//                costZero++
//            }
//		}
//        // costZero := len(str) - costOne
//
//		// 关于符合条件的组合情况，子集数量 + 1
//		for i := m - costZero; i >= 0; i-- {
//			for j := n - costOne; j >= 0; j-- {
//				if matrix[i][j] != -1 && matrix[i][j]+1 > matrix[i+costZero][j+costOne] {
//					matrix[i+costZero][j+costOne] = matrix[i][j] + 1
//				}
//			}
//		}
//	}
//	max := 0
//	for _, a := range matrix {
//		for _, v := range a {
//			if max < v {
//				max = v
//			}
//		}
//	}
//	return max
