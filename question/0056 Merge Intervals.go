package question

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]。
//请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
import "sort"

type IntSliceSlice [][]int

func (p IntSliceSlice) Less(i, j int) bool {
	return p[i][0] < p[j][0] // 左端点排序
}
func (p IntSliceSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	return
}
func (p IntSliceSlice) Len() int {
	return len(p)
}
func (p IntSliceSlice) Sort() { sort.Sort(p) }
func merge(intervals [][]int) [][]int {
	return mergeCore(intervals)
}
func mergeCore(intervals [][]int) [][]int {
	//先按照left排序
	i := IntSliceSlice(intervals)
	i.Sort()
	//假设已经排好序
	var res [][]int
	res = append(res, intervals[0]) // 拿出第一个区间作为初始值
	k := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[k][1] { // 如果这个区间的左端点在范围内，说明重叠了
			// 合并区间
			if intervals[i][1] > res[k][1] {
				res[k][1] = intervals[i][1]
			}
		} else {
			// 不重叠，则作为最后一个区间，然后继续
			k++
			res = append(res, intervals[i])
		}
	}
	return res
}
