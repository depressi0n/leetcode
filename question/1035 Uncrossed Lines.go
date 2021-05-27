package question

func MaxUncrossedLines(nums1 []int, nums2 []int) int {
	// 首先设置一个矩阵，nums1[i] = nums2[j] 则matrix[i][j]=true
	// 然后通过动态规划 dp[i][j] 表示nums1[0:i+1],nums2[0:j+1]能得到的无交叉的线有多少条
	// 初始化 dp[0][0]=0 dp[0][j] dp[i][0] 如果matrix[0]j] matrix[i][0]为true，则为1否则为0
	// dp[i+1][j+1] = dp[i][j]+matrix[i][j],重新更新一下？这样的动态规划似乎不是很优秀
	//matrix := make([][]bool, len(nums1))
	//for i := 0; i < len(nums1); i++ {
	//	matrix[i] = make([]bool, len(nums2))
	//	for j := 0; j < len(nums2); j++ {
	//		if nums1[i] == nums2[j] {
	//			matrix[i][j] = true
	//		}
	//	}
	//}
	dp := make([][]int, len(nums1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2))
	}
	// 初始化
	i := 0
	for ; i < len(dp) && nums1[i] != nums2[0]; i++ {

	}
	for ; i < len(dp); i++ {
		dp[i][0] = 1
	}
	i = 0
	for ; i < len(dp[0]) && nums1[0] != nums2[i]; i++ {

	}
	for ; i < len(dp[0]); i++ {
		dp[0][i] = 1
	}
	// 转移方程 dp[i+1][j] = if matrixt[i][k] { dp[i+1][j]=max(dp[i+1][j],dp[i][j-1]+1)
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
			dp[i][j] = max(max(dp[i][j], dp[i-1][j]), dp[i][j-1])

		}
	}
	return dp[len(nums1)-1][len(nums2)-1]

	//TODO:三角形的优化策略，用一个数配上一个数组，优化空间
}
