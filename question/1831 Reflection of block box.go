package question

// 模拟光路会超时
//type BlackBox struct {
//	n, m int
//	// 记录每个孔的状态和位置
//	location map[struct {
//		x, y int
//	}]int
//	indexes map[int]struct {
//		x, y int
//	}
//	status                map[int]bool
//	next                  map[struct{ index, dir int }]int
//	up, down, left, right int
//}
//
//func BlackBoxConstructor(n int, m int) BlackBox {
//	b := BlackBox{
//		n:        n,
//		m:        m,
//		location: make(map[struct{ x, y int }]int, 2*(n+m)),
//		indexes:  make(map[int]struct{ x, y int }, 2*(n+m)),
//		status:   make(map[int]bool, 2*(n+m)),
//		next:     map[struct{ index, dir int }]int{},
//	}
//	index := 0
//	// 初始化每个孔均为关闭
//	for i := 0; i <= m; i++ {
//		b.location[struct {
//			x, y int
//		}{i, 0}] = index
//		b.indexes[index] = struct {
//			x, y int
//		}{i, 0}
//		b.status[index] = false
//		index++
//	}
//	b.up = index
//	for i := 1; i <= n; i++ {
//		b.location[struct {
//			x, y int
//		}{m, i}] = index
//		b.indexes[index] = struct {
//			x, y int
//		}{m, i}
//		b.status[index] = false
//		index++
//	}
//	b.right = index
//	for i := m - 1; i >= 0; i-- {
//		b.location[struct {
//			x, y int
//		}{i, n}] = index
//		b.indexes[index] = struct {
//			x, y int
//		}{i, n}
//		b.status[index] = false
//		index++
//	}
//	b.down = index
//	for i := n - 1; i > 0; i-- {
//		b.location[struct {
//			x, y int
//		}{0, i}] = index
//		b.indexes[index] = struct {
//			x, y int
//		}{0, i}
//		b.status[index] = false
//		index++
//	}
//	b.left = index
//	return b
//}
//
//func (this *BlackBox) Open(index int, direction int) int {
//	isCorner := func(loc struct{ x, y int }) bool {
//		if (loc.x == 0 && loc.y == 0) || (loc.x == this.m && loc.y == 0) || (loc.x == this.m && loc.y == this.n) || (loc.x == 0 && loc.y == this.n) {
//			return true
//		}
//		return false
//	}
//	// 打开
//	this.status[index] = true
//	var preLoc, nextLoc struct {
//		x, y int
//	}
//	var next int
//	var ok bool
//	curIndex := index
//	dir := direction
//	for {
//		// 利用缓存
//		next, ok = this.next[struct {
//			index, dir int
//		}{curIndex, dir}]
//		if ok && this.status[next] {
//			return next
//		}
//		preLoc = struct {
//			x, y int
//		}{
//			this.indexes[curIndex].x,
//			this.indexes[curIndex].y,
//		}
//		nextLoc = struct {
//			x, y int
//		}{
//			-1,
//			-1,
//		}
//		// 位置和方向检查
//		// 光线开始动
//		switch dir {
//		case -1:
//			if curIndex < this.up {
//				// (x,0)
//				if preLoc.x+this.n <= this.m {
//					// check the (x+m,m)
//					nextLoc.x = preLoc.x + this.n
//					nextLoc.y = this.n
//				} else {
//					// check the(n,n-x)
//					nextLoc.x = this.m
//					nextLoc.y = this.m - preLoc.x
//				}
//			} else if curIndex < this.right {
//				// (n,y)
//				if this.m <= preLoc.y {
//					// check the (0,y-n)
//					nextLoc.x = 0
//					nextLoc.y = preLoc.y - this.m
//				} else {
//					// check the(n-y,0)
//					nextLoc.x = this.m - preLoc.y
//					nextLoc.y = 0
//				}
//			} else if curIndex < this.down {
//				// (x,m)
//				if preLoc.x <= this.n {
//					// check the (0,m-x)
//					nextLoc.x = 0
//					nextLoc.y = this.n - preLoc.x
//				} else {
//					// check the(x-m,0)
//					nextLoc.x = preLoc.x - this.n
//					nextLoc.y = 0
//				}
//			} else {
//				// (0,y)
//				if this.m <= this.n-preLoc.y {
//					// check the (n,y+n)
//					nextLoc.x = this.m
//					nextLoc.y = preLoc.y + this.m
//				} else {
//					// check the(m-y,m)
//					nextLoc.x = this.n - preLoc.y
//					nextLoc.y = this.n
//				}
//			}
//		case 1:
//			// 四个方向
//			if curIndex < this.up {
//				// (x,0)
//				if preLoc.x <= this.n {
//					// check the (0,x)
//					nextLoc.x = 0
//					nextLoc.y = preLoc.x
//				} else {
//					// check the(x-m,m)
//					nextLoc.x = preLoc.x - this.n
//					nextLoc.y = this.n
//				}
//			} else if curIndex < this.right {
//				// (n,y)
//				if this.m <= this.n-preLoc.y {
//					// check the (0,y+n)
//					nextLoc.x = 0
//					nextLoc.y = preLoc.y + this.m
//				} else {
//					// check the(n-m+y,m)
//					nextLoc.x = this.m - this.n + preLoc.y
//					nextLoc.y = this.n
//				}
//			} else if curIndex < this.down {
//				// (x,m)
//				if this.m-preLoc.x <= this.n {
//					// check the (n,m-n+x)
//					nextLoc.x = this.m
//					nextLoc.y = this.n - this.m + preLoc.x
//				} else {
//					// check the(x+m,0)
//					nextLoc.x = preLoc.x + this.n
//					nextLoc.y = 0
//				}
//			} else {
//				// (0,y)
//				if preLoc.y <= this.m {
//					// check the (y,0)
//					nextLoc.x = preLoc.y
//					nextLoc.y = 0
//				} else {
//					// check the(n,y-n)
//					nextLoc.x = this.m
//					nextLoc.y = preLoc.y - this.m
//				}
//			}
//		}
//		next = this.location[nextLoc]
//		this.next[struct {
//			index, dir int
//		}{curIndex, dir}] = next
//
//		if this.status[next] {
//			return next
//		}
//		// 角落
//		if isCorner(this.indexes[next]) {
//			// 光路可逆
//			return index
//		}
//		preLoc = nextLoc
//		curIndex = next
//		dir = -dir
//	}
//}
//
//func (this *BlackBox) Close(index int) {
//	this.status[index] = false
//}

/**
 * Your BlackBox object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Open(index,direction);
 * obj.Close(index);
 */

// TODO：思考初始状态和最终状态，并考虑用什么数据结构去包装数据，思考状态转移（不一定是动态规划）
// 考虑index之间的关系以及光线：
// （1）在第i个孔上以y=x的角度上射入的，下一个应该到达第2*(m+n)+2-i个孔
// （2）在第i个孔上以y=-x的角度上射入的，下一个应该达到4*m+2*n+2-i或者2*m+2-i个孔，前者发生在y=-x的左下方，后者发生在右上方
// 其次的目的就是找出所有可能的循环，而这个循环肯定是有规律可以寻找的，最终只会有4*(m+n)-4种状态，多个状态会在同一个循环中出现

//type BlackBox struct {
//	circles      [][]int
//	index2Circle map[struct{ index, direction int }]int
//	index2Pos    map[struct{ index, direction int }]int
//	status       map[int]bool
//}
//
//func BlackBoxConstructor(n int, m int) BlackBox {
//	b:=BlackBox{
//		circles:      [][]int{},
//		index2Circle: make(map[struct{ index, direction int }]int),
//		index2Pos:    make(map[struct{ index, direction int }]int),
//		status:       make(map[int]bool),
//	}
//	createCircle := func(index int, direction int) []int {
//		res := []int{}
//		switch direction {
//		case 1:
//			res = append(res, 2*(m+n)-index)
//		case -1:
//
//		}
//	}
//}
//
//func (this *BlackBox) Open(index int, direction int) int {
//
//}
//
//func (this *BlackBox) Close(index int) {
//
//}

/**
 * Your BlackBox object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Open(index,direction);
 * obj.Close(index);
 */

// TODO:
// type BlackBox struct {
//	Hole []bool
//	P    []int
//	N    []int
//	T      int
//	L      int
//	M      int
//	Num    int
//}
//
//func Constructor(n int, m int) BlackBox {
//	Box := new(BlackBox)
//	Box.M = m
//	Box.Num = n
//	l := 2*m + 2*n
//	Box.L = l
//	Box.Hole = make([]bool, Box.L)
//	Box.T = 0
//	Box.P = make([]int, Box.L) // pre
//	Box.N = make([]int, Box.L) // next
//
//	for i := 0; i < l; i++ {
//		Box.P[i] = (l - i) % l //从i以y=x射向 (l - i) % l点
//		Box.N[i] = (m * 2 - i + l) % l //从i以y=x射向(m * 2 - i + l) % l
//	}
//
//	return *Box
//}
//
//func (this *BlackBox) Open(index int, direction int) int {
//	if !this.Hole[index] {
//		this.T ++
//		this.Hole[index] = true //表示打开
//	}
//	if this.T == 1 { //只打开来当前的则只可能从这里出去
//		return index
//	} else {
//		for {
//			if direction == 1 {
//				index = this.P[index]
//			} else {
//				index = this.N[index]
//			}
//			direction = -direction
//			if this.Hole[index] {
//				break;
//			}
//		}
//	}
//
//	return index
//}
//
//func (this *BlackBox) Close(index int) {
//	this.Hole[index] = false
//	this.T--
//}
