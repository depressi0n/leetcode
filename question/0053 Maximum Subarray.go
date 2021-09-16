package question

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 注意是连续子数组
func maxSubArray(nums []int) int {
	return maxSubArrayCore3(nums)
}

// 首先是 通过 动态规划的思路
func maxSubArrayCore1(nums []int) int {
	// dp[i] 表示以nums[i]结尾的最大的连续子数组和是多少
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	res := dp[0]
	for i := 1; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

// 根据上面的动态规划的情况，可以把辅助数组去掉，因为只取决于前一个值
func maxSubArrayCore2(nums []int) int {
	res := nums[0]
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(nums[i], nums[i]+cur)
		res = max(res, cur)
	}
	return res
}

// 由这道题引出"线段树"的概念，除了维护[0,n)外，还可以维护任意的[l,r)的问题，时间复杂度为对数级别，支持修改、查询等功能
// 定义一个操作get(nums,l,r)，表示查询nums中[l,r)的最大子段和
// 通过分治策略，将[l,r)分解为[l,(l+r)/2)),[(l+r)/2,r)
// 两个关键问题：
// （1）维护区间的什么信息
// （2）如何合并维护的信息
// 对于（1）中[l,r)，维护四个信息：
// 	lSum以l为左端点的最大子段和：要么等于左区间的lSum,要么等于左区间的iSum加上右区间的lSum
//	rSum以r-1为右端点的最大子段和：要么等于右区间的rSum，要么等于右区间的iSum加上左区间的rSum
//	mSum[l,r)间的最大子段和：要么是左区间的mSum，要么是右区间的mSum，要么是左区间的rSum加上右区间的lSum
//	iSum[l,r)区间和:[l,(l+r)/2)),[(l+r)/2,r)的区间和的和

type Status struct {
	lSum, rSum, mSum, iSum int
}

func pushUp(l, r Status) Status {
	return Status{
		lSum: max(l.lSum, l.iSum+r.lSum),
		rSum: max(r.rSum, r.iSum+l.rSum),
		mSum: max(max(l.mSum, r.mSum), l.rSum+r.lSum),
		iSum: l.iSum + r.iSum,
	}
}
func get(nums []int, l int, r int) Status {
	if l == r-1 {
		return Status{
			lSum: nums[l],
			rSum: nums[l],
			mSum: nums[l],
			iSum: nums[l],
		}
	}
	m := (l + r - 1) >> 1
	lSub := get(nums, l, m+1)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}
func maxSubArrayCore3(nums []int) int {
	return get(nums, 0, len(nums)).mSum
}
