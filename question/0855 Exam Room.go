package question

import "container/list"

type ExamRoom struct {
	capacity int
	done     []byte
}

func Constructor0855(n int) ExamRoom {
	return ExamRoom{
		capacity: n,
		done:     make([]byte, (n+7)/8),
	}
}

func (this *ExamRoom) Seat() int {
	res := this.findNextLoc()
	this.done[res/8] |= 1 << (res % 8)
	return res
}

func (this *ExamRoom) Leave(p int) {
	this.done[p/8] ^= 1 << (p % 8)
}

func (this *ExamRoom) findNextLoc() int {
	//distancs := make([]int, this.capacity)
	//min := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//// 计算距离
	//i := 0
	//for ; i < this.capacity && !this.done[i]; i++ {
	//}
	//
	//if i != 0 {
	//	for j := 0; j < i; j++ {
	//		distancs[j] = i - j
	//	}
	//}
	//if i<this.capacity{
	//	distancs[i] = 0
	//}
	//for j := i + 1; j < this.capacity; j++ {
	//	if this.done[j] {
	//		for k := i + 1; k < j; k++ {
	//			distancs[k] = min(k-i, j-k)
	//		}
	//		i = j
	//	}
	//}
	//// 最后一个
	//for k := i + 1; k < this.capacity; k++ {
	//	distancs[k] = k - i
	//}
	//res := -1
	//for k := 0; k < this.capacity; k++ {
	//	if res == -1 && distancs[k] != 0 {
	//		res = k
	//		continue
	//	}
	//	if res != -1 && distancs[k] > distancs[res] {
	//		res = k
	//	}
	//
	//}
	//return res

	// 实际上不需要维护distance数组
	// ___1___1___1___

	max := -1
	res := -1
	// 前部位置
	i := 0
	for ; i < this.capacity && (this.done[i/8]&(1<<(i%8)) == 0); i++ {
	}
	if 0 < i && i < this.capacity {
		max = i
		res = 0
	}
	if i == this.capacity {
		return 0
	}
	// 中间位置
	j := i + 1
	for j < this.capacity {
		for ; j < this.capacity && (this.done[j/8]&(1<<(j%8)) == 0); j++ {

		}
		if j == this.capacity {
			break
		}
		// 取中间
		if (j-i)/2 > max {
			max = (j - i) / 2
			res = i + (j-i)/2
		}
		i = j
		j += 1
	}
	// 后部位置
	if i < this.capacity {
		if max < this.capacity-i-1 {
			max = this.capacity - 1 - i
			res = this.capacity - 1

		}
	}
	return res
}

// TODO，超时，看懂下面的
type ExamRoom1 struct {
	seat *list.List // 表示坐着的同学的位置
	n    int
}

func Constructor08552(N int) ExamRoom1 {
	return ExamRoom1{
		seat: list.New(),
		n:    N - 1,
	}
}

func (this *ExamRoom1) Seat() int {
	// 还没有人入座，直接将0插入
	if this.seat.Len() == 0 {
		this.seat.PushFront(0)
		return 0
	}
	e := this.seat.Front()
	pre := e.Value.(int)
	max := pre // 头部需要特殊判断
	addVal := 0
	addFront := e
	e = e.Next()
	for ; e != nil; e = e.Next() {
		val := e.Value.(int)
		distant := (val - pre) / 2 // 两点之间的最远距离
		if distant > max {
			max = distant
			addFront = e           // 需要插入的点的后一个元素。方便找到后直接插入
			addVal = pre + distant // 需要插入的位置
		}
		pre = val
	}
	distant := this.n - pre // 尾部特殊判断
	if distant > max {
		this.seat.PushBack(this.n) // 直接插入到链表尾部
		return this.n
	}
	this.seat.InsertBefore(addVal, addFront) // 插入
	return addVal
}

// 遍历知道对应的位置删除
func (this *ExamRoom1) Leave(p int) {
	for e := this.seat.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			this.seat.Remove(e)
			return
		}
	}
	return
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
// 离最近的人最远
// 如果有多个 则选择序号最小的
// 第一个进来的选择0
