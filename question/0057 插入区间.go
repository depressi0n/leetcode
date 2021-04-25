package question

// TODO：1. 使用二分查找来加速
// 2. 贪心算法，分情况讨论6
func insert(intervals [][]int, newInterval []int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var res [][]int
	i := 0                                                        // 找到之前不变的地方，至少i-1之前的都不会变
	for i < len(intervals) && intervals[i][0] <= newInterval[0] { //新插入的left必然在[i-1,i)
		i++
	}
	//修改的是i-1
	//特殊情况
	if i == 0 { //说明left在原来区间集合的左边
		if newInterval[1] < intervals[0][0] {
			//在最前面插入一个
			res = append(res, newInterval)
			res = append(res, intervals...)
			return res
		} else {
			//修改第一个
			res = append(res, newInterval)
			//找到right的位置
			j := 0
			for j < len(intervals) && intervals[j][1] < newInterval[1] {
				j++
			}
			if j == 0 {
				res[0][1] = intervals[0][1]
				res = append(res, intervals[1:]...)
				return res
			}
			if j >= len(intervals) || newInterval[1] < intervals[j][0] { //合并到j-1
				res[0][1] = newInterval[1]
				res = append(res, intervals[j:]...)
			} else { //合并到j
				res[0][1] = intervals[j][1]
				res = append(res, intervals[j+1:]...)
			}
			return res
		}
	}
	if newInterval[0] <= intervals[i-1][1] { //修改i-1，可能需要合并之后的窗口
		j := i - 1
		for j < len(intervals) && intervals[j][1] < newInterval[1] {
			j++
		}
		if j == i-1 { // 不用修改
			res = append(res, intervals...)
			return res
		}
		res = append(res, intervals[:i]...)
		if j >= len(intervals) {
			res[i-1][1] = newInterval[1]
			return res
		}
		if newInterval[1] < intervals[j][0] { //合并到j-1
			res[i-1][1] = newInterval[1]
			res = append(res, intervals[j:]...)
			return res
		} else { //合并到j
			res[i-1][1] = intervals[j][1]
			res = append(res, intervals[j+1:]...)
			return res
		}
	} else {
		j := i
		for j < len(intervals) && intervals[j][1] < newInterval[1] {
			j++
		}
		// 修改到是i
		res = append(res, intervals[:i]...)
		if j >= len(intervals) || newInterval[1] < intervals[j][0] { //合并到j-1
			if j == i { //新插入一个
				res = append(res, newInterval)
			} else {
				res = append(res, intervals[i]) //修改原来的
				res[i][0] = newInterval[0]
				res[i][1] = newInterval[1]
			}
			res = append(res, intervals[j:]...)
			return res
		} else { //合并到j
			res = append(res, intervals[i])
			res[i][0] = newInterval[0]
			res[i][1] = intervals[j][1]
			res = append(res, intervals[j+1:]...)
			return res
		}
	}
}
