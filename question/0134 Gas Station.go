package question

// 在一条环路上有N个加油站，其中第i个加油站有汽油gas[i]升。
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。你从其中的一个加油站出发，开始时油箱为空。
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
//说明:
//如果题目有解，该答案即为唯一答案。
//输入数组均为非空数组，且长度相同。
//输入数组中的元素均为非负数。
func canCompleteCircuit(gas []int, cost []int) int {
	return canCompleteCircuitCore1(gas, cost)
}

// 主要思想，因为容量无限，所以可以将所有的油全部加入油箱
// 然后从任何一个站点出发，只要能在此回到这个地点
// 则返回这个站点的编号
// 但这样会超时？？？  => 因为每个站点都要检查，可以通过减小被检查的加油站数目来降低总时间复杂度
func canCompleteCircuitCore1(gas []int, cost []int) int {
	// 可以先计算remain，然后求和即可
	// remain的含义是汽车此站点达到下一个站点至少有remain的油量
	remain := make([]int, len(cost))
	for i := 0; i < len(cost); i++ {
		remain[i] = gas[i] - cost[i]
	}
	// 从任何一个站点出发，计算
	for i := 0; i < len(remain);i++{
		// 说明可以从这个站点出发达到下一个站点
		if remain[i] >= 0 {
			cur, have := (i+1)%len(cost), remain[i]
			for ; cur != i && have >= 0; cur = (cur + 1) % len(cost) {
				have += remain[cur]
			}
			if cur == i && have>=0 {
				return i
			}
		}
	}
	return -1
}
// 主要思想：因为每个站点都要检查，可以通过减小被检查的加油站数目来降低总时间复杂度
// 从x站点出发，每经过一个加油站就要加一次油，最后一个可以达到的加油站是y
// 那么至少 满足sum(gas,x,y) < sum(cost,x,y) && sum(gas,x,y)>=sum(cost,x,i) for all i -> [x.y)
// 对于 z -> [x,y]，从该加油站出发，是否可以达到y+1即 sum(gas,z,y) 与 sum(cost,z,y)
// sum(gas,z,y) = sum(gas,z,y) - sum(gas,z,y-1) < sum(cost,x,y) - sum(gas,x,z-1) < sum(cost,x,y) - sum(cost,x.z-1) = sum(cost,z,y)
// 这就意味着 z ->[x,y] 都无法达到y的下一个加油站，这也就意味着中间的加油站可以不用查了，直接查下一个
func canCompleteCircuitCore2(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			// 中间的一部分都可以直接跳过，不用再重复
			i += cnt + 1
		}
	}
	return -1
}
