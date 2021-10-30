package question

// 老师想给孩子们分发糖果，有 N个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
//那么这样下来，老师至少需要准备多少颗糖果呢？
func candy(ratings []int) int {
	return candyCore4(ratings)
}

// 主要思想：将题设分成从左往右和从右往左两个规则
// 从左往右，如果右侧排级比左侧高，则增加一颗，否则，将其变为1
// 然后从右往左，如果左侧排级比右侧高，则增加一颗
// 最后累加起来
func candyCore1(ratings []int) int {
	// 先每个人分一颗
	own := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		own[i] = 1
	}
	// left to right
	for i := 0; i < len(own)-1; i++ {
		// 如果右侧孩子排级高，那么就应该增加糖果
		for ratings[i+1] > ratings[i] && own[i+1] <= own[i] {
			own[i+1]++
		}
	}
	// right to left
	for i := len(own) - 1; i > 0; i-- {
		for ratings[i-1] > ratings[i] && own[i-1] <= own[i] {
			own[i-1]++
		}
	}

	var res int
	for i := 0; i < len(own); i++ {
		res += own[i]
	}
	return res
}

// 上面思想的另一种实现：
// 先从左往右，所有孩子的糖果满足左侧规则
// 后从右往左，所有孩子的糖果满足右侧规则
// 最后综合两者，孩子能分到的糖果必须满足二者的最大值
func candyCore2(ratings []int) int {
	left := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right, res := 0, 0
	for i := len(ratings) - 1; i >= 0; i++ {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return res
}

// 升降坡的思想：
// 对于左规则而言：当左侧孩子的排序与右侧孩子的排序相同，则只给1个糖果，如果较大，则多给1个，否则 不知道应该怎么判断，因为后面可能跟了一个很长的递减序列，只满足一侧是不行？？？
// 但是如果我们看递减序列，知道递减序列的长度，就能判断最少的糖果数目
// 也就是保证，两侧递减序列的较大值就是给孩子的糖果数目
func candyCore3(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	// 从i向左看，左侧递减序列的长度
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	// 从i向右看，右侧递减序列的长度，数组可以被优化为一个值，看上面的结果。
	right := make([]int, len(ratings))
	right[len(ratings)-1] = 1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}
	res := 0
	for i := 0; i < len(ratings); i++ {
		res += max(left[i], right[i])
	}
	return res
}

// 常数空间的遍历：糖果要尽可能的少给，还是升降坡的思想
// 一个特殊情况是峰值可以作为递增序列的最后一个，也可以作为递减序列的第一个
// 记录前一个同学的糖果为pre：
// 如果当前同学比上一个同学评分高或相等，则说明现在处于递增序列中，可以给该同学分配pre+1个糖果
// 反之，则在一个递减序列中，那么分配给该同学1个糖果，并将该同学所在递减序列中所有同学都再多分配一个糖果
// 在此过程中，我们维护当前递减序列的长度，同时如果当前递减序列的长度和上一个递增序列的长度相等时，则最近的递增序列的最后一个同学要并入递减序列中
// 总而言之，我们维护当前递减序列的长度dec，最近的递增学列的长度inc，前一个同学分配的糖果数目pre
func candyCore4(ratings []int) int {
	res, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			// 当前处于递增序列中
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			res += pre
			inc = pre
		} else {
			// 当前处于递减序列
			dec++
			if dec == inc {
				dec++
			}
			res += dec
			pre = 1
		}
	}
	return res
}
