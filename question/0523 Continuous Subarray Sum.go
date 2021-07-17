package question

func CheckSubarraySum(nums []int, k int) bool {
	//mod k 为0的连续数组

	// O(n^2)
	// 预处理
	//for i := 0; i < len(nums); i++ {
	//	nums[i]%=k
	//}
	//prefixSum:=make([]int,len(nums)+1)
	//prefixSum[0]=0
	//for i := 1; i <= len(nums); i++ {
	//	prefixSum[i]=prefixSum[i-1]+nums[i-1]
	//}
	//for length:=2;length<len(prefixSum);length++{
	//	for i := 0; i < len(nums); i++ {
	//		if i+length >len(nums){
	//			break
	//		}
	//		if (prefixSum[i+length]-prefixSum[i])%k== 0{
	//			return true
	//		}
	//	}
	//}
	//return false

	// 利用模的意义进行处理
	// 如果有一段的模余为0，则必然有两个前缀和相等的情况出现
	// map中存放nums[0:i]和位置i
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % k
		if pos, ok := m[sum]; ok {
			if i-pos >= 2 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
