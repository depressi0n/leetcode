package question

import "math"

func LargestNumber1449(cost []int, target int) string {
	// 目标：总费用应该等于target
	// cost 表示画一个数字的费用，没有数字0
	// 求最大的数字
	// 思路：预处理数字，然后使用可能的贪心策略，相同的代价下选择最大的数字，尽可能增长数字
	// 这种策略会超时，涉及到target对不同cost的可能划分情况
	//maxSingleCost := cost[0]
	//minSingleCost := cost[0]
	//mayCost := make(map[int]struct{})
	//m := make(map[int]int)
	//for i := 0; i < len(cost); i++ {
	//	if cost[i] > maxSingleCost {
	//		maxSingleCost = cost[i]
	//	}
	//	if cost[i] < minSingleCost {
	//		minSingleCost = cost[i]
	//	}
	//	m[cost[i]] = i + 1
	//	mayCost[cost[i]]= struct{}{}
	//}
	//var dfs func(prefix int, remain int) int
	//dfs = func(prefix int, remain int) int {
	//	if remain == 0 {
	//		return prefix
	//	}
	//	if remain < minSingleCost {
	//		return 0
	//	}
	//	ret := 0
	//	for cos, _ := range mayCost {
	//		cur := dfs(prefix*10+m[cos], remain-cos)
	//		if ret < cur {
	//			ret = cur
	//		}
	//	}
	//	return ret
	//}
	//return strconv.Itoa(dfs(0, target))
	// 优化策略：TODO（1）直接划分target （2）然后从划分的结果中选择出最大的值
	//maxSingleCost := cost[0]
	//minSingleCost := cost[0]
	//mayCost := make(map[int]struct{})
	//m := make(map[int]int)
	//for i := 0; i < len(cost); i++ {
	//	if cost[i] > maxSingleCost {
	//		maxSingleCost = cost[i]
	//	}
	//	if cost[i] < minSingleCost {
	//		minSingleCost = cost[i]
	//	}
	//	m[cost[i]] = i + 1
	//	mayCost[cost[i]] = struct{}{}
	//}
	//divideRes := make([]map[int]int, 0, target)
	//maxLength:=0
	//var divide func(cur map[int]int,length int, remain int)
	//divide = func(cur map[int]int,length int, remain int) {
	//	if remain == 0 {
	//		// 只要最长的
	//		if length  < maxLength {
	//			return
	//		}else{
	//			c:=make(map[int]int)
	//			for v, cnt := range cur {
	//				c[v]=cnt
	//			}
	//			if length==maxLength{
	//				divideRes=append(divideRes,c)
	//				return
	//			}
	//			divideRes=[]map[int]int{c}
	//			maxLength=length
	//		}
	//		return
	//	}
	//	if remain < minSingleCost {
	//		return
	//	}
	//	if length+ remain/minSingleCost<maxLength{
	//		return
	//	}
	//	for cos:=minSingleCost;cos<=maxSingleCost;cos++{
	//		if _,ok:=mayCost[cos];ok{
	//			cur[cos]++
	//			divide(cur,length+1,remain-cos)
	//			cur[cos]--
	//			if cur[cos]==0{
	//				delete(cur,cos)
	//			}
	//		}
	//	}
	//}
	//divide(map[int]int{},0,target)
	//res:=0
	//for i := 0; i < len(divideRes); i++ {
	//	tmp:=make([]int,0,len(divideRes[i]))
	//	for v, cnt := range divideRes[i] {
	//		for cnt>0{
	//			tmp=append(tmp,m[v])
	//		}
	//	}
	//	sort.Ints(tmp)
	//	t:=0
	//	for j := len(tmp)-1; j >=0 ; j-- {
	//		t=t*10+tmp[j]
	//	}
	//	if t>res{
	//		res=t
	//	}
	//}
	//return strconv.Itoa(res)

	// TODO:转换为背包问题
	// dp[i][j]  表示使用前i+1个数位且成本刚好为j的最大位数，
	// 成本无法为j则为-1，dp[0][0]=0，dp[0][j]=-1
	// dp[i+1][j] =dp[i][j] if j < cost[i]  else  max(dp[i][j],dp[i+1][j-cost[i]]+1)
	// dp[9][target]表示位数最大的值
	// 使用辅助数组from[i][j] 记录转移来源，若转移来源是dp[i+1][j-cost[i]] 则说明选取了第i个数位
	// from[i+1][j] = j if j <cost[i] OR dp[i][j]>dp[i+1][cost[i]] else j-cost[i]
	// 为了生成的整数尽可能的大，当前数位应当尽可能多选择，即若dp[i][j] == dp[i+1][j-cost[i]]+1 应当选择后者转移
	dp := make([][]int, 10)
	from := make([][]int, 10)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
		from[i] = make([]int, target+1)
	}
	dp[0][0] = 0
	for i, c := range cost {
		for j := 0; j <= target; j++ {
			if j < c {
				dp[i+1][j] = dp[i][j]
				from[i+1][j] = j
			} else {
				if dp[i][j] > dp[i+1][j-c]+1 {
					dp[i+1][j] = dp[i][j]
					from[i+1][j] = j
				} else {
					dp[i+1][j] = dp[i+1][j-c] + 1
					from[i+1][j] = j - c
				}
			}
		}
	}
	if dp[9][target] < 0 {
		return "0"
	}
	ans := make([]byte, 0, dp[9][target])
	i, j := 9, target
	for i > 0 {
		if j == from[i][j] {
			i--
		} else {
			ans = append(ans, '0'+byte(i))
			j = from[i][j]
		}
	}
	return string(ans)
	// TODO (1)滚动数组优化，dp[i+1][i]的计算只用到了dp[i][k]和dp[i+1][j-1]之前的运算
	// TODO (2)去掉from数组，直接根据dp[j]与dp[j-cost[i]]+1是否相等来判断
	//max:=func(a,b int)int{
	//        if a<b{
	//            return b
	//        }
	//        return a
	//    }
	//    dp := make([]int, target+1)
	//    for i := range dp {
	//        dp[i] = math.MinInt32
	//    }
	//    dp[0] = 0
	//    for _, c := range cost {
	//        for j := c; j <= target; j++ {
	//            dp[j] = max(dp[j], dp[j-c]+1)
	//        }
	//    }
	//    if dp[target] < 0 {
	//        return "0"
	//    }
	//    ans := make([]byte, 0, dp[target])
	//    for i, j := 8, target; i >= 0; i-- {
	//        for c := cost[i]; j >= c && dp[j] == dp[j-c]+1; j -= c {
	//            ans = append(ans, byte('1'+i))
	//        }
	//    }
	//    return string(ans)
}
