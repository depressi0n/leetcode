package question

// 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。

func maxPoints0149(points [][]int) int {
	return maxPoints0149Core(points)
}

// 主要思想：求出所有的斜率和纵截距，但这里的斜率可能会出现误差，不一定准确！！！
// 这里对斜率(my/mx)进行处理，首先进行一次约分，然后根据题目给的范围，使用单个32整型变量来表示表示两个整数
// val=my+(2*10^4+1)*mx
func maxPoints0149Core(points [][]int) (ans int) {
	n := len(points)
	if n <= 2 {
		return n
	}
	for i, p := range points {
		// （1）当我们枚举到点i时，往后寻找至多能找到n-i个点共线，如果在此之前已经找到的最大值为k
		// 且k≥n−i，那么此时我们即可停止枚举，因为不可能再找到更大的答案了。
		// （2）当我们找到一条直线经过了图中超过半数的点时，我们即可以确定该直线即为经过最多点的直线；
		if ans >= n-i || ans > n/2 {
			break
		}
		cnt := map[int]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd1(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*20001]++
		}
		for _, c := range cnt {
			ans = max(ans, c+1)
		}
	}
	return
}

func gcd1(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//最大公约数
func gcd2(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return gcd2(b%a, a)
}
