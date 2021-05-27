package question

import (
	"math"
)

func MinChanges(nums []int, k int) int {
	// 连续k个数的异或和为0
	// 最少需要改变多少个元素才能达到
	// 基本思路是 nums[i] = nums[k+i] = ... = nums[x*k+i] 0<=i<k
	// 分组后，找到每一组的主元素
	// 用每一组的数量减去主元素的数量，和即为所求
	// 如果有一组所有的数都要修改，则选那个主元素数量最小的
	// 如果每一组都只修改部分数，那么使用动态规划
	// 存在一定问题，无法判断是否异或和为0，其次，无法找到改变最小的，所以用动态规划来处理！
	//res:=0
	//if k==1{
	//	for i := 0; i < len(nums); i++ {
	//		if nums[i]!=0{
	//			res++
	//		}
	//	}
	//	return res
	//}
	//// 按照k进行分组，为了节省空间，可以考虑进行原地置换处理?
	//help:=make([][]int,k)
	//helpM:=make([]map[int]int,k)
	//for i := 0; i < k; i++ {
	//	help[i]=make([]int,k)
	//	helpM[i]=make(map[int]int)
	//	for j := 0; j*k+i <len(nums); j++ {
	//		help[i]=append(help[i],nums[j*k+i])
	//		helpM[i][nums[j*k+i]]++
	//	}
	//}
	//// 加入动态规划进行最优求解
	//// dp[i][mask] 表示已经处理的前i组，对应的异或和的可能为mask时的最小次数
	//// 状态转移方程
	//// dp[i][mask] =size(i)+min{dp[i-1][mask^x]-count(i,x)}
	//// 其中x表示第i+1组被修改为x，那么第i组中需要改变的元素的数量为size(i)-count(i,x)
	//// 优化策略：
	//// 如果count(i,x) == 0,这意味这所有的数都会被修改，此时可以得到任意一个异或值
	//// 如果count(i,x) !=0，这说明只是修改部分数，枚举之前的状态，此时的
	//dp:=make([]map[int]int,len(help))
	//for i := 0; i < len(dp); i++ {
	//	dp[i]=make(map[int]int)
	//}
	//// 初始化
	//for v, cnt := range helpM[0] {
	//	dp[0][v]=len(help[0])-cnt
	//}
	////还要枚举其他没有达到的状态，设置为len(help[0])
	//for i := 1; i < k; i++ {
	//	for v, cnt := range helpM[i] {
	//		for xor, minCh := range dp[i-1] {
	//			if _,ok:=dp[i][xor^v];!ok {
	//				dp[i][xor^v]=minCh+len(help[i])-cnt
	//			}else{
	//				dp[i][xor^v]=min(dp[i][xor^v],minCh+len(help[i])-cnt)
	//			}
	//		}
	//		// 其他的目标值
	//	}
	//}
	//return len(nums)-res

	//TODO
	const maxX = 1 << 10          // x 的范围为 [0, 2^10)
	const inf = math.MaxInt32 / 2 // 极大值，为了防止整数溢出
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v < res {
				res = v
			}
		}
		return res
	}
	n := len(nums)
	f := make([]int, maxX)
	for i := range f {
		f[i] = inf
	}
	// 边界条件 f(-1,0)=0
	f[0] = 0

	for i := 0; i < k; i++ {
		// 第 i 个组的哈希映射
		cnt := map[int]int{}
		size := 0
		for j := i; j < n; j += k {
			cnt[nums[j]]++
			size++
		}

		// 求出 t2
		t2min := min(f...)

		g := make([]int, maxX)
		for j := range g {
			g[j] = t2min
		}
		for mask := range g {
			// t1 则需要枚举 x 才能求出
			for x, cntX := range cnt {
				g[mask] = min(g[mask], f[mask^x]-cntX)
			}
		}

		// 别忘了加上 size
		for j := range g {
			g[j] += size
		}
		f = g
	}

	return f[0]
}

// TODO best 看懂
//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
//
//func max(a, b int) int {
//	if a >= b {
//		return a
//	}
//	return b
//}
//
//func minChanges(nums []int, k int) int {
//	fs := make([]map[int]int, k)
//	for i := 0; i < k; i++ {
//		fs[i] = map[int]int{}
//	}
//	mas, mass := make([]int, k), make([]int, k)
//	for i, n := range nums {
//		fs[i%k][n]++ // 统计第i%k+1组中出现n的次数
//	}
//
//	for i, f := range fs { //对于第i+1组
//		for _, c := range f { //出现的最多的次数
//			mas[i] = max(mas[i], c) // mas[i]记录的是出现最多的数字出现的次数
//		}
//	}
//	mass[k-1] = mas[k-1] //从后往前推mass
//	for i := k - 2; i >= 0; i-- {
//		mass[i] = mass[i+1] + mas[i] // mass[i]表示后缀和
//	}
//
//	m := mas[0] //m用于记录出现mas中的最小值
//	for i := 1; i < k; i++ {
//		m = min(m, mas[i])
//	}
//	o := mass[0] - m //记录最多保留？
//
//	var cal func(int, int, int) //变成了dfs，有点像背包问题
//	cal = func(i, s, c int) {
//		if i == k {
//			if s == 0 { //目标值为0时能保留下来的数字有多少个
//				o = max(o, c)
//			}
//		} else if c+mass[i] > o {
//			for j, f := range fs[i] {
//				cal(i+1, s^j, c+f)
//			}
//		}
//	}
//
//	cal(0, 0, 0)
//	return len(nums) - o
//}
