package question

import (
	"fmt"
	"log"
	"strconv"
)

// 和之前题目的思路很像，dfs从前往后遍历一到三个值，遇到的情况有：
//（1）这个值是一个合法值，那么可以选择，然后处理下一个值
//（2）这个值不是一个合法值，直接返回
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	isValid := func(index, length int) (int, bool) {
		if index+length > len(s) {
			return -1, false
		}
		if length == 1 { //长度为1都是可以的
			return int(s[index] - '0'), true
		}
		if s[index] == '0' {
			return -1, false
		}
		value, err := strconv.Atoi(s[index : index+length])
		if err != nil {
			log.Fatal(err)
		}
		if 0 <= value && value < 256 {
			return value, true
		}
		return -1, false
	}
	var dfsIn93 func(int, int, []int)
	dfsIn93 = func(index int, length int, cur []int) {
		value, ok := isValid(index, length)
		if !ok {
			return
		}
		cur = append(cur, value)
		if len(cur) == 4 && index+length == len(s) {
			res = append(res, fmt.Sprintf("%d.%d.%d.%d", cur[0], cur[1], cur[2], cur[3]))
			return
		}

		index += length
		for i := 1; i <= 3; i++ {
			if index+i <= len(s) {
				dfsIn93(index, i, cur)
			}
		}
	}
	dfsIn93(0, 1, []int{})
	dfsIn93(0, 2, []int{})
	dfsIn93(0, 3, []int{})
	return res
}
