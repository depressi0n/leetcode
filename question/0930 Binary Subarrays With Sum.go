package question

func Answer0930(nums []int, goal int) int {
	return numSubarraysWithSum(nums, goal)
}
func numSubarraysWithSum(nums []int, goal int) int {
	// nums[i] = 0 or 1
	// 目标：找出所有的子数组，使得和为goal
	// 思路：双指针法 TODO:把这里debug出来
	//left:=0
	//right:=1
	//// 辅助数组 prefixSum[i]=sum of num[0:i)
	//prefixSum:=make([]int,len(nums)+1)
	//prefixSum[0]=0
	//for i := 1; i <= len(nums); i++ {
	//	prefixSum[i]=prefixSum[i-1]+nums[i-1]
	//}
	//res:=0
	//if prefixSum[len(nums)]<goal{
	//	return 0
	//}
	//for left<len(nums){
	//	for right<=len(nums) && prefixSum[right]-prefixSum[left] < goal{
	//		right++
	//	}
	//	if right>len(nums) {
	//		return res
	//	}
	//	// 如果右边有0，往右走并计数
	//	rightZeroNum:=1
	//	for right+rightZeroNum-1<len(nums) && nums[right+rightZeroNum-1]==0{
	//		rightZeroNum++
	//	}
	//	// 如果左边有0，往右走并计数
	//	leftZeroNum:=1
	//	for left<=right && nums[left]==0{
	//		leftZeroNum++
	//		left++
	//	}
	//	res+=leftZeroNum*rightZeroNum
	//	if left<=right{
	//		left++
	//		right=left+1
	//	}
	//
	//}
	//return res

	// TODO：看懂下面的思想
	ans := 0
	left1, left2 := 0, 0 //对于每个right，都会有[[left1,left2],right]都满足条件
	sum1, sum2 := 0, 0
	for right, num := range nums {
		//两个left的逻辑大体是一样的，只是一个必须小于等于，一个必须小于
		sum1 += num                         //把右边加上
		for left1 <= right && sum1 > goal { //比goal大
			sum1 -= nums[left1] // 去掉左边
			left1++             // 左边右移
		}
		sum2 += num                          // 把右边加上
		for left2 <= right && sum2 >= goal { //和goal相同则必须右移，区分好
			sum2 -= nums[left2] // 去掉左边
			left2++             // 左边右移
		}
		ans += left2 - left1
	}
	return ans
}

// 0 0 0
// 0 0 0 0
// [left,right]=prefixSum[right]-prefixSum[left]
//
