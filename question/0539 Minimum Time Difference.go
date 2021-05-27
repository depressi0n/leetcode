package question

func FindMinDifference(timePoints []string) int {
	// 可以认为是找到两个时间点之间的最小间距 HH:MM
	// 首先最暴力的做法肯定是计算所有的差值，然后找到最小值，时间复杂度为n^2
	// 然后稍微优化一点的方式是将字符串转化成数字，然后排序【或者直接根据字符串也可以排序，指定排序规则即可】，最后计算间距找到最小值，时间复杂度为nlogn
	// 最后的想法是 分而治之
	//（1）从时间的表达上出发，一共24个小时，每个小时60分钟，最多有多少种表达呢，24*60=1440种
	//（2）字符数组长度超过这个值，则由鸽巢原理知必然有重复出现的时间点，直接返回0即可
	//（3）用桶排序，每个小时一个桶，然后计算差值时只需要计算桶内的，
	//（4）加上一个策略是，如果桶的大小超过60，或者出现了相同的时间点，直接返回0
	//（5）最后计算桶内和桶间的最小间距即可
	// 更好方式应该是基数排序? 或者直接用长度1440的布尔数组来记录，省事多了
	if len(timePoints) < 2 {
		return -1
	}
	if len(timePoints) > 24*60 {
		return 0
	}
	buckets := make(map[uint8]map[uint8]bool, 24)
	for i := 0; i < 24; i++ {
		buckets[uint8(i)] = make(map[uint8]bool, 60)
	}
	for i := 0; i < len(timePoints); i++ {
		// 转化
		if _, ok := buckets[(timePoints[i][0]-'0')*10+(timePoints[i][1]-'0')][(timePoints[i][3]-'0')*10+(timePoints[i][4]-'0')]; ok {
			return 0
		}
		buckets[(timePoints[i][0]-'0')*10+(timePoints[i][1]-'0')][(timePoints[i][3]-'0')*10+(timePoints[i][4]-'0')] = true
	}
	res := 24 * 60
	min := -1
	last := -1 // 记录上一个桶的最大值(可能中间间隔了很多bucket)，这里需要形成一个循环
	// 对桶进行处理
	for i := 0; i < 24; i++ {
		// 只需要找到桶内差比res的距离小的即可
		first := false
		for j := 0; j < 60; j++ {
			if _, ok := buckets[uint8(i)][uint8(j)]; ok {
				if !first {
					first = true
					if last != -1 && res > i*60+j-last {
						res = i*60 + j - last
					}
				}
				// 首个值需要计算桶之间的差值
				last = i*60 + j
				if min == -1 {
					min = last
				}
				// 从1开始往后查找
				delta := 1
				for delta < res && delta < 60 {
					if _, ok := buckets[uint8(i)][uint8(j+delta)]; ok {
						res = delta
						break
					} else {
						delta++
					}
				}
			}
		}
	}
	// 最后处理循环的事情
	if res > min-last+1440 {
		res = min - last + 1440
	}
	return res
}
