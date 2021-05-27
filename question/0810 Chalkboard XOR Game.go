package question

func xorGame(nums []int) bool {
	// 一次擦除一个元素
	// 明显的动态规划问题 或者 说一个记忆化搜索问题（这里的状态转移方程是不好写的）
	// 但是确实不好做，从数学上考虑一下异或
	// dp[i][j] 表示 nums[i:j) 中先手的人是否能赢得游戏
	//
	// suffixSum[i] 表示 nums[:i)的XOR值
	//suffixSum := make([]int, len(nums)+1)
	//for i := 1; i < len(nums)+1; i++ {
	//	suffixSum[i] = suffixSum[i-1] ^ nums[i-1]
	//}
	//dp := make([][]bool, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	dp[i] = make([]bool, len(nums)+1)
	//	// 初始化
	//	if nums[i] == 0 {
	//		dp[i][i+1] = true
	//	}else{
	//		dp[i][i+1]=false
	//	}
	//}

	// 从数学上考虑一下异或
	// 先手的人每次选择的时候所遇见的奇偶性是完全相同的，从这个角度入手
	// 记所有的异或和为S
	// 如果全部元素的异或和不为0即S！=0，那么先手的人肯定可以选择一个如nums[i]，擦除后的结果记作S_i
	// 如果对于所有的i，S_i均为0，那么S_0 ^ S_1 ^ ... ^ S_(n-1) = 0
	// 即可以得到 当n是偶数的时候，可以推出S=0 而当n是奇数时，
	//
	if len(nums)%2 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0

}
