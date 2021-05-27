package question

func StoneGame(piles []int) bool {
	// 选择最左或者最右
	// 最后得分高的人赢
	//if len(piles) <= 2 {
	//	return true
	//}
	//abs:=func(a int)int{
	//	if a<0{
	//		return -a
	//	}
	//	return a
	//}
	//max:=func(a,b int)int{
	//	if a<b{
	//		return b
	//	}
	//	return a
	//}
	// dp[i][j] 表示 piles[i:j)中Alex能获得的最大分数
	//dp := make([][]int, len(piles)+1)
	//for i := 0; i < len(piles)+1; i++ {
	//	dp[i] = make([]int, len(piles)+1)
	//	if i<len(piles){
	//		dp[i][i+1] = piles[i]
	//	}
	//	if i+1<len(piles){
	//		dp[i][i+2] = abs(piles[i]-piles[i+1])
	//	}
	//}
	//for length := 3; length <= len(piles); length++ {
	//	for i := 0; i < len(piles)+1; i++ {
	//		if i+length > len(piles) {
	//			break
	//		}
	//		dp[i][i+length]=max(piles[i]-dp[i+1][i+length],piles[i+length-1]-dp[i][i+length-1])
	//	}
	//}
	//if dp[0][len(piles)]>=0{
	//	return true
	//}else{
	//	return false
	//}

	// dp优化
	//dp:=make([]int,len(piles))
	//for i := 0; i < len(piles); i++ {
	//	dp[i]=piles[i]
	//}
	//for length := 2; length <= len(piles) ; length++ {
	//	for i := 0; i <= len(piles)-length; i++ {
	//		dp[i]=max(piles[i]-dp[i+1],piles[i+length-1]-dp[i])
	//	}
	//}
	//return dp[0]>=0

	// 数学方法
	// 根据奇偶下标将数组分为两组，偶数为第一组，奇数为第二组
	// 先手的Alex可以自由选择第一组的一堆石子或第二组的一堆石子
	// 如果Alex取走第二组的一堆石子，则剩下的部分在行开始处和结束处的石子堆都属于第一组，因此Lee只能选择第一组的石子，Lee取完后剩下的部分在行开始处和结束处的石子堆一定不同组
	// 接下来Alex依旧做上面的选择
	// 所以先手的人决定取走哪一组石子，从而具有必胜策略：每次选择数量更多的那一组的相应的一个
	return true

}
