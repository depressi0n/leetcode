package question

import (
	"fmt"
)

func BrokenCalc(X int, Y int) int {
	// 经过t次操作后，范围是[x-t,x*2^t]
	// 所有y如果要在t次能到达的范围内，至少是需要t次
	// 那个就要看是x和y的相对大小了
	// 如果y比x小，那么只能通过减法达到
	// 如果y比x大，那么可能就要经过各种方式
	// 应该用一种dp的方式，将y去减半之类的，涉及到奇数或者偶数，而且只能向上取整，如7只能做加法到8而不能做减法到6
	cnt := 0
	tmp := Y
	for tmp != X {
		if tmp < X {
			break
		}
		fmt.Println(tmp)
		if tmp&1 == 1 {
			tmp++
			cnt++
		} else {
			// 比较此时的状态
			// [x ... tmp]
			tmp /= 2
			cnt++
		}
	}
	return cnt + (X - tmp) - 1
}
