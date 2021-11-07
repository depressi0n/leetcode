package question

//给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

func distributeCandies(candyType []int) int {
	return distributeCandiesCore2(candyType)
}
// 这么做会超时，复杂度太高了
// 主要思想：转成 背包问题
func distributeCandiesCore1(candyType []int) int {
	// 一人拿一半，但要求拿到的数字里面，不同的数字数量尽可能多
	// 万事都有可能
	res := 0
	cache := make(map[int]int)
	var dfs func(cur int, have int)
	dfs = func(cur int, have int) {
		if have<<1 == len(candyType) {
			res = max(res, len(cache))
			return
		}
		if cur >= len(candyType) {
			return
		}
		dfs(cur+1, have)
		if _, ok := cache[candyType[cur]]; !ok {
			cache[candyType[cur]] = 1
		} else {
			cache[candyType[cur]]++
		}
		dfs(cur+1, have+1)
		if cache[candyType[cur]] == 1 {
			delete(cache, candyType[cur])
		}else{
			cache[candyType[cur]]--
		}
		have -= 1

	}
	dfs(0, 0)
	return res
}
// 考虑种类数目的问题？
// 如果种类数目超过n/2，那么肯定可以是n/2
// 如果种类数目不超过n/2，这个值肯定是最后结果
func distributeCandiesCore2(candyType []int) int {
	m:=make(map[int]struct{})
	for i := 0; i < len(candyType); i++ {
		if _,ok:=m[candyType[i]];!ok{
			m[candyType[i]]= struct{}{}
		}
	}
	if len(m)<<1 >= len(candyType){
		return len(candyType)/2
	}else{
		return len(m)
	}
}
