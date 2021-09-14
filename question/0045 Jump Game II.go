package question

// 题目描述：给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//假设你总是可以到达数组的最后一个位置。
func jump(nums []int) int {
	return jumpCore3(nums)
}
// 从后往前，看每一个位置到最后一个位置的最小跳跃数目
func jumpCore1(nums []int) int {
	dp:=make([]int,len(nums))
	dp[len(nums)-1]=0
	for i:=len(nums)-2;i>=0;i--{
		if nums[i]+i>=len(nums)-1{
			dp[i]=1
			continue
		}
		dp[i]=len(nums)
		for j:=i+1;j<=nums[i]+i;j++{
			dp[i]=min(dp[i],dp[j]+1)
		}
	}
	return dp[0]
}

// 贪心算法: 反向查找，和上面动态规划的思想差不多
func jumpCore2(nums []int) int {
	pos:=len(nums)-1
	steps:=0
	for pos>0{
		for i:=0;i<pos;i++{
			if i+nums[i]>=pos{
				pos=i
				steps++
				break
			}
		}
	}
	return steps
}



// 贪心算法: 正向查找，记录当前能达到的最远位置及其需要的步数
func jumpCore3(nums []int) int {
	end:=0
	maxPos:=0
	steps:=0
	for i:=0;i<len(nums)-1;i++{
		maxPos=max(maxPos,i+nums[i])
		if i==end{
			end=maxPos
			steps++
		}
	}
	return steps
}