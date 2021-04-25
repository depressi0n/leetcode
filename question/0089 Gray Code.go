package question

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{0, 1}
	for i := 1; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]|(1<<i))
		}
	}
	fmt.Println(res)
	return res
}
