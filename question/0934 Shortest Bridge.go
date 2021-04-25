package question

//type union struct {
//	father []int
//	rank []int
//}
//
//func newUnion(size int) union {
//	u:=union{
//		father: make([]int, size),
//		rank:   make([]int, size),
//	}
//	for i := 0; i < size; i++ {
//		u.father[i]=i
//		u.rank[i]=1
//	}
//}
//func (u *union) union(x int,y int) {
//	xGrand:=u.find(x)
//	yGrand:=u.find(y)
//	if u.rank[xGrand]<=u.rank[yGrand]{
//		u.father[xGrand]=yGrand
//	}else{
//		u.father[yGrand]=xGrand
//	}
//	if u.rank[xGrand]==u.rank[yGrand] && xGrand!=yGrand{
//		u.rank[yGrand]++
//	}
//}
//
//func (u *union) find(x int) int {
//	for x!=u.father[x]{
//		u.father[x]= u.find(u.father[x])
//	}
//	return u.father[x]
//}
// 两个岛，连接起来的最短路径
func ShortestBridge(A [][]int) int {
	// 并查集肯定可以做，但是复杂度可能会上去
	help := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		help[i] = make([]int, len(A[0]))
	}
	mark := func(i, j int) int {
		level := 2
		// 用两个队列，将周围的0交替标记长度
		q1 := make([]int, 0, len(A)*len(A))
		nextQ := make([]int, 0, len(A)*len(A))
		// 当前的1入队
		help[i][j] = 1 //标记已经访问过
		q1 = append(q1, i*len(A)+j)
		cur := 0
		for cur != len(q1) {
			x := q1[cur]
			y := x % (len(A))
			x = x / (len(A))
			// 将周围没有被标记过的标记
			// up
			if x-1 >= 0 && help[x-1][y] == 0 {
				if A[x-1][y] == 1 {
					help[x-1][y] = 1
					q1 = append(q1, (x-1)*len(A)+y)
				} else {
					help[x-1][y] = level //第几层
					nextQ = append(nextQ, (x-1)*len(A)+y)
				}
			}
			// down
			if x+1 < len(A) && help[x+1][y] == 0 {
				if A[x+1][y] == 1 {
					help[x+1][y] = 1
					q1 = append(q1, (x+1)*len(A)+y)
				} else {
					help[x+1][y] = level //第几层
					nextQ = append(nextQ, (x+1)*len(A)+y)
				}
			}
			// left
			if y-1 >= 0 && help[x][y-1] == 0 {
				if A[x][y-1] == 1 {
					help[x][y-1] = 1
					q1 = append(q1, x*len(A)+y-1)
				} else {
					help[x][y-1] = level //第几层
					nextQ = append(nextQ, x*len(A)+y-1)
				}
			}
			// right
			if y+1 < len(A) && help[x][y+1] == 0 {
				if A[x][y+1] == 1 {
					help[x][y+1] = 1
					q1 = append(q1, x*len(A)+y+1)
				} else {
					help[x][y+1] = level //第几层
					nextQ = append(nextQ, x*len(A)+y+1)
				}
			}
			cur++
		}
		// 从nextQ队列开始，再次遇到1说明可以连接，最短的距离就是level当前值
		cur = 0
		level++
		start := len(nextQ) // 直接存放到后面
		for cur < start {
			x := nextQ[cur]
			y := x % (len(A))
			x = x / (len(A))
			// 将周围没有被标记过的标记
			// up
			if x-1 >= 0 && help[x-1][y] == 0 {
				if A[x-1][y] == 1 {
					return level - 2
				}
				help[x-1][y] = level //第几层
				nextQ = append(nextQ, (x-1)*len(A)+y)

			}
			// down
			if x+1 < len(A) && help[x+1][y] == 0 {
				if A[x+1][y] == 1 {
					return level - 2
				}
				help[x+1][y] = level //第几层
				nextQ = append(nextQ, (x+1)*len(A)+y)

			}
			// left
			if y-1 >= 0 && help[x][y-1] == 0 {
				if A[x][y-1] == 1 {
					return level - 2
				}
				help[x][y-1] = level //第几层
				nextQ = append(nextQ, x*len(A)+y-1)

			}
			// right
			if y+1 < len(A) && help[x][y+1] == 0 {
				if A[x][y+1] == 1 {
					return level - 2
				}
				help[x][y+1] = level //第几层
				nextQ = append(nextQ, x*len(A)+y+1)
			}
			cur++
			if cur == start {
				level++
				start = len(nextQ)
			}
		}
		return level
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			// 遍历一次
			if A[i][j] == 1 {
				return mark(i, j)
			}
		}
	}
	return 1
}
