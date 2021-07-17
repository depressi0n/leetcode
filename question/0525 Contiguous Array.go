package question

func FindMaxLength(nums []int) int {
	// 10 和 01 可以配成一组 题意理解错了
	// 返回这样连续的组数组成的子数组的长度
	// dp[i][j] 表示 nums[i:j]中符合上述要求的最长子数组的长度
	// dp[i][j+1] = max(dp[i][j-1]+dp[j-1][j+1],dp[i][j],dp[i][i+2]+dp[i+2][j+1])
	// 这么做是不对的，这样返回的是最多的「01」 「10」的组数
	//dp := make([][]int, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	dp[i] = make([]int, len(nums)+1)
	//	if i+1 < len(nums) && nums[i]^nums[i+1] == 1 {
	//		dp[i][i+2] = 2
	//	}
	//}
	//for length := 3; length <= len(nums); length++ {
	//	for i := 0; i+length <= len(nums); i++ {
	//		dp[i][i+length] = func(a, b, c int) int {
	//			if a < b {
	//				a = b
	//			}
	//			if a < c {
	//				a = c
	//			}
	//			return a
	//		}(dp[i][i+length-1], dp[i][i+2]+dp[i+2][i+length],dp[i][i+length-2]+dp[i+length-2][i+length])
	//	}
	//}
	//return dp[0][len(nums)]

	// 为了保证连续，引入suffix和prefix
	// prefix[i] 表示nums[i]开始的最长的子数组的长度
	// suffix[j] 表示nums[j]结束的最长的子数组的长度
	//prefix:=make([]int,len(nums))
	//// 初始化
	//if nums[len(nums)-2]^nums[len(nums)-1] == 1 {
	//	prefix[len(nums)-2]=2
	//}
	//// 从后往前更新preifx
	//for i:=len(nums)-3;i>=0;i--{
	//	if nums[i]^nums[i+1] == 1 {
	//		prefix[i]=2+prefix[i+2]
	//	}
	//}
	//suffix:=make([]int,len(nums))
	//if nums[1]^nums[0]==1{
	//	suffix[1]=2
	//}
	//for i:=2;i<len(nums);i++{
	//	if nums[i]^nums[i-1] == 1{
	//		suffix[i]=2+suffix[i-2]
	//	}
	//}
	//res:=max(prefix[0],suffix[len(nums)-1])
	//for i := 1; i+2 < len(nums); i++ {
	//	if nums[i]^nums[i+1]==1{
	//		res=max(res,suffix[i-1]+2+prefix[i+2])
	//	}
	//}
	//return res

	// 题意是要求一个子数组中具有相同数目的0和1，求最长的这样的子数组的长度
	// 把0看作-1，把1看作1，求最长的数组和为0的子数组的长度
	// 常规解法 用前缀和数组
	//prefixSum:=make([]int,len(nums)+1)
	//prefixSum[0]=0
	//for i := 1; i < len(prefixSum); i++ {
	//	prefixSum[i]=prefixSum[i-1]+nums[i-1]
	//}
	//// 结果肯定比总长度短，而且必须是偶数
	//// 考虑用二分法
	//length:=(len(nums)>>1) << 1
	//fmt.Printf("length=%d,",length)
	//for length>1{
	//	for i := 0; i+length <= len(nums); i++ {
	//		if prefixSum[i+length]-prefixSum[i]==length/2{
	//			return length
	//		}
	//	}
	//	length-=2
	//}
	//return 0

	// 优化思路：前缀和+哈希表
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxLength := 0
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return maxLength

}
