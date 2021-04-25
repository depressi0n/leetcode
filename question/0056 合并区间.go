package question

import "sort"

type IntSliceSlice [][]int

func (p IntSliceSlice) Less(i, j int) bool {
	return p[i][0] < p[j][0]
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
	//先按照left排序
	i := IntSliceSlice(intervals)
	i.Sort()
	//假设已经排好序
	var res [][]int
	res = append(res, intervals[0])
	k := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[k][1] {
			if intervals[i][1] > res[k][1] {
				res[k][1] = intervals[i][1]
			}
		} else {
			k++
			res = append(res, intervals[i])
		}
	}
	return res
}
