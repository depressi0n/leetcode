package question

import (
	"sort"
	"strconv"
)

type IntSlice []int

func (p IntSlice) Len() int { return len(p) }
func (p IntSlice) Less(i, j int) bool {
	x := strconv.Itoa(p[i])
	y := strconv.Itoa(p[j])
	x, y = x+y, y+x
	//先把短的补齐
	//if len(x) < len(y) {
	//	x += strings.Repeat(x[len(x)-1:], len(y)-len(x))
	//} else {
	//	y += strings.Repeat(y[len(y)-1:], len(x)-len(y))
	//}
	cur := 0
	for cur < len(x) && cur < len(y) {
		if x[cur] < y[cur] {
			return true
		} else if x[cur] > y[cur] {
			return false
		}
		cur++
	}
	//for cur < len(x) { //说明x比较长
	//	if x[cur] < y[len(y)-1] {
	//		return true
	//	} else if x[cur] > y[len(y)-1] {
	//		return false
	//	}
	//	cur++
	//}
	//for cur < len(y) { //说明y比较长
	//	if y[cur] < x[len(x)-1] {
	//		return false
	//	} else if y[cur] > x[len(x)-1] {
	//		return true
	//	}
	//	cur++
	//}
	return true
}

func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func largestNumber(nums []int) string {
	// 按照基数排序的规则
	s := IntSlice(nums)
	sort.Sort(s)
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += strconv.Itoa(nums[i])
	}
	//fmt.Println(res)
	//去掉前置零
	for len(res) != 0 && res[0] == '0' {
		res = res[1:]
	}
	if res == "" {
		res = "0"
	}
	return res
}

//TODO:可以简化自定义排序，只需要自定义一个Less函数即可
