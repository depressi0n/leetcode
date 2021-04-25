package question

func canCompleteCircuit(gas []int, cost []int) int {
	//可以先计算remain，然后求和即可
	remain := make([]int, len(cost))
	for i := 0; i < len(cost); i++ {
		remain[i] = gas[i] - cost[i]
	}
	for i := 0; i < len(remain); i++ {
		if remain[i] >= 0 {
			cur := remain[i]
			k := i + 1
			for ; k < len(remain) && cur+remain[k] >= 0; k++ {
				cur += remain[k]
			}
			if k < len(remain) {
				continue
			}
			for k = 0; k < i && cur+remain[k] >= 0; k++ {
				cur += remain[k]
			}
			if k == i {
				return i
			}
		}
	}
	return -1
	////不先计算
	//enter := func(i int) bool {
	//	remain:=gas[i] - cost[i]
	//	for k:=i+1;k<len(cost);k++{
	//		if remain+gas[k]<cost[k]{
	//			return false
	//		}
	//		remain+=gas[k]-cost[k]
	//	}
	//	for k:=0;k<i;k++{
	//		if remain+gas[k]<cost[k]{
	//			return false
	//		}
	//		remain+=gas[k]-cost[k]
	//	}
	//	return true
	//}
	//for i := 0; i < len(cost); i++ {
	//	if cost[i] <= gas[i] {
	//		if enter(i) {
	//			return i
	//		}
	//	}
	//}
	//return -1
}
