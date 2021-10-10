package question

import (
	"log"
	"sync/atomic"
	"unsafe"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func nextPermutationUnique(nums []int) []int {
	var i, j int
	j = len(nums) - 1
	for j > 0 && nums[j] <= nums[j-1] { // 找到第一个递减的
		j--
	}
	i = j
	for i < len(nums)-1 && nums[i+1] == nums[i] { //找到将被交换的位置
		i++
	}
	if j == 0 {
		for n := 0; n <= (len(nums)-1)/2; n++ {
			nums[n], nums[len(nums)-n-1] = nums[len(nums)-n-1], nums[n]
		}
		return nums
	}
	j--
	k := len(nums) - 1
	for k >= i && nums[k] <= nums[j] { // 找到第一个比nums[j]大的
		k--
	}
	nums[j], nums[k] = nums[k], nums[j] //交换
	for n := j + 1; n <= (len(nums)+j)/2; n++ {
		nums[n], nums[len(nums)+j-n] = nums[len(nums)+j-n], nums[n]
	}
	return nums
}

func checkListEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	p, q := l1, l2
	for p != nil && q != nil {
		if p.Val != q.Val {
			log.Printf("p=%d,", p.Val)
			log.Printf("q=%d", q.Val)
			log.Printf("\n")
			return false
		}
		p = p.Next
		q = q.Next
	}
	if p != nil || q != nil {
		return false
	}
	return true
}

func makeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  arr[0],
		Next: makeList(arr[1:]),
	}
	return head
}

func makeTree(arr []int, null int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	q := make([]*TreeNode, 0, len(arr))
	q = append(q, root)
	cur := 0
	for i := 1; i < len(arr); i += 2 {
		if arr[i] == null {
			q[cur].Left = nil
		} else {
			q[cur].Left = &TreeNode{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}
			q = append(q, q[cur].Left)
		}
		if i+1 < len(arr) {
			if arr[i+1] == null {
				q[cur].Right = nil
			} else {
				q[cur].Right = &TreeNode{
					Val:   arr[i+1],
					Left:  nil,
					Right: nil,
				}
				q = append(q, q[cur].Right)
			}
		}
		cur++
	}
	return root
}

// 无锁队列， TODO：正确性待验证
// lock free的正确解释为一个“锁无关”的程序能够确保执行它的所有线程中至少有一个能够继续往下执行。

type DQNode struct {
	Val        int
	Prev, Next *DQNode
}
type LockFreeQueue struct {
	head, tail *DQNode
}

func (q *LockFreeQueue) InitQueue() {
	node := &DQNode{
		Val:  0,
		Prev: nil,
		Next: nil,
	}
	q.head = node
	q.tail = node
}
func (q *LockFreeQueue) EnQueue(data int) {
	node := &DQNode{
		Val:  data,
		Prev: nil,
		Next: nil,
	}
	// 初始版本

	//p := q.tail
	//for !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(p.Next)), nil, unsafe.Pointer(node)) {
	//	p = q.tail
	//}
	// 不需要判断是否成功，因为只有成功的线程才能往下走
	//atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.tail)), unsafe.Pointer(p), unsafe.Pointer(node))

	// 但是如果成功的线程挂掉了，就会死循环，改进版
	//p := q.tail
	//oldP := p
	//for !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(p.Next)), nil, unsafe.Pointer(node)) {
	//	// 每个线程自己将指针移动到末尾，可以优化：将大家共享的q.tail移动到末尾
	//	for p.Next != nil {
	//		p = p.Next
	//	}
	//}
	//// 不需要判断是否成功，因为只有成功的线程才能往下走
	//atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.tail)), unsafe.Pointer(oldP), unsafe.Pointer(node))

	// 优化版本
	var tail *DQNode
	for {
		// 取下尾指针和next
		tail = q.tail
		next := tail.Next
		// 此时尾部指针已经移动了，则重新来过
		if tail != q.tail {
			continue
		}
		// 如果next不是nil，说明不是末尾，则将tail换成next即后移
		if next != nil {
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.tail)), unsafe.Pointer(tail), unsafe.Pointer(next))
			continue
		}
		// 如果节点加入成功，则跳出循环
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(tail.Next)), unsafe.Pointer(next), unsafe.Pointer(node)) {
			break
		}
	}
	// 不需要判断是否成功，因为只有成功的线程才能往下走
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.tail)), unsafe.Pointer(tail), unsafe.Pointer(node))
}

func (q *LockFreeQueue) DeQueue() int {
	//p := q.head
	//for !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.head)), unsafe.Pointer(p), unsafe.Pointer(p.Next)) {
	//	if p.Next == nil {
	//
	//		return -1
	//	}
	//	p = q.head
	//}
	//return p.Next.Val
	// 为了避免当head和tail指向同一个节点时（此时队列为空），但在判断p.Next时，另一个入队操作做了一半，此时并不为空，但tail指针还未后移
	// 导致入队操作尚未完成就将新加入节点出队，导致队列此时为空，但head和tail并不指向同一个结点
	// 给出如下优化版本
	var head, tail, next *DQNode
	var value int
	for {
		head = q.head
		tail = q.tail
		next = head.Next
		//	如果q.head已经移动，重来
		if head != q.head {
			continue
		}
		// 如果队空
		if head == tail && next == nil {
			// ERROR_EMPTY_QUEUE
			return -1
		}
		// 如果tail指针落后了
		if head == tail && next == nil {
			atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.tail)), unsafe.Pointer(tail), unsafe.Pointer(next))
			continue
		}

		// 移动head指针成功后，取出数据
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(q.tail)), unsafe.Pointer(tail), unsafe.Pointer(next)) {
			value = next.Val
			break
		}
	}
	return value
}

// ABA问题，尤其是内存重用导致的ABA问题
// 解决ABA问题的一种方式 -- double-CAS，前半部分是值，后半部分是计数器，CAS操作时同时检查二者
// safeRead(){
// loop:
//   p=q.next
//   if p == nil {
//      return p
//   }
//   FetchAndAdd(p.refcnt,1) // 原子操作，加引用计数
//   if p == q.next {
//      return p
//   }else{
//      Release(p) // 原子操作，减引用计数
//   }
//   goto loop
//}

// 使用数组实现无锁队列
// 基本思想：环形数组，数组元素可能取值为HEAD、TAIL、EMPTY和实际数据
// 初始化时数组元素全部为EMPTY，有两个相邻元素要初始化为HEAD和TAIL，表示空队列
// 入队操作：首先定位TAIL的位置，使用double-CAS将（TAIL，EMPTY）更新为（x，TAIL），如果找不到（TAIL，EMPTY）则说明队列满
// 出队操作：首先定位HEAD的位置，将（HEAD，x）更新为（EMPTY，HEAD），并返回x，如果x为TAIL，说明队空
// 如果定位HEAD和TAIL：声明两个计数器，一个计数入队次数，一个计数出队次数，均使用FetchAndAdd进行原子累加，模运算得到位置

// 维护区间和的线段树
// 原始数组下标：默认从1开始
// 线段树下标：加入线段树中某个位置的下标，如原始数组的第一个数，一般会加入到线段树中的第二个位置
// 存储下标：元素所在叶节点编号，实际存储位置
// 一般满足：
// 线段树下标+N-1 = 存储下标
// 原始数组下标+1=线段树下标
// 也满足 => 原始数组+N=存储下标
// 其中 N 大于等于n+2的某个2的次方，n为原始数组元素个数
type rangeTree struct {
	Sum []int // 求和
	Add []int // 懒惰标记
	A   []int // 原始数组下标【1，n】
	N   int
}

func (rt *rangeTree) PushUp(index int) {
	rt.Sum[index] = rt.Sum[index<<1] + rt.Sum[index<<1|1]
}
func (rt *rangeTree) Build(l, r int, index int) {
	if l == r {
		rt.Sum[index] = rt.A[l]
		return
	}
	mid := (l + r) >> 1
	rt.Build(l, mid, index<<1)
	rt.Build(mid+1, r, index<<1|1)
	rt.PushUp(index)
}
func (rt *rangeTree) BuildNonRec(l, r int, index int) {
	rt.N = 1
	for rt.N < len(rt.A) {
		rt.N <<= 1
	}
	for i := 1; i <= len(rt.A); i++ {
		rt.Sum[rt.N+i] = rt.A[i]
	}
	for i := rt.N - 1; i > 0; i-- {
		rt.Sum[i] = rt.Sum[i<<1] + rt.Sum[i<<1|1]
		rt.Add[i] = 0
	}
}

// Update 单点修改
func (rt *rangeTree) Update(L int, C int, l, r int, index int) {
	// l,r表示当前节点区间，index表示节点编号
	if l == r {
		rt.Sum[index] += C
		return
	}
	mid := (l + r) >> 1
	if L <= mid {
		rt.Update(L, C, l, mid, index<<1)
	} else {
		rt.Update(L, C, mid+1, r, index<<1|1)
	}
	rt.PushUp(index)
}
func (rt *rangeTree) UpdateNonRec(L int, C int) {
	// l,r表示当前节点区间，index表示节点编号
	for s := rt.N + L; s != 0; s >>= 1 {
		rt.Sum[s] += C
	}
}
func (rt *rangeTree) Query(L, R int) int {
	res := 0
	s := rt.N + L - 1
	t := rt.N + R + 1
	for s^t^1 != 0 {
		if s&1 == 0 {
			res += rt.Sum[s^1]
		}
		if t&1 != 0 {
			res += rt.Sum[t^1]
		}
		s >>= 1
		t >>= 1
	}
	return res
}
func (rt *rangeTree) QueryNonRec(L, R int) int {
	res := 0
	s := rt.N + L - 1
	t := rt.N + R + 1
	for s^t^1 != 0 {
		if s&1 == 0 {
			res += rt.Sum[s^1]
		}
		if t&1 != 0 {
			res += rt.Sum[t^1]
		}
		s >>= 1
		t >>= 1
	}
	return res
}

// RangeUpdate 区间修改
func (rt *rangeTree) RangeUpdate(L, R int, C int, l, r int, index int) {
	if L <= l && r <= R {
		rt.Sum[index] += C * (r - l + 1)
		rt.Add[index] += C
		return
	}
	mid := (l + r) >> 1
	if L <= mid {
		rt.RangeUpdate(L, R, C, l, mid, index<<1)
	}
	if mid < R {
		rt.RangeUpdate(L, R, C, mid+1, r, index<<1|1)
	}
	rt.PushUp(index)
}
func (rt *rangeTree) RangeUpdateNonRec(L, R int, C int) {
	// ln: s一路走来已经包含的几个数
	// rn: t一路走来已经包含的几个数
	// x: 本层每个节点包含几个数
	s, t, ln, rn, x := rt.N+L-1, rt.N+R+1, 0, 0, 1
	for s^t^1 != 0 {
		rt.Sum[s] += C * ln
		rt.Sum[t] += C * rn
		if s&1 == 0 {
			rt.Add[s^1] += C
			rt.Sum[s^1] += C * x
			ln += x
		}
		if t&1 != 0 {
			rt.Add[t^1] += C
			rt.Sum[t^1] += C * x
			rn += x
		}
		s >>= 1
		t >>= 1
		x <<= 1
	}
	for s != 0 {
		rt.Sum[s] += C * ln
		rt.Sum[t] += C * rn
		s >>= 1
		t >>= 1
	}
}

func (rt *rangeTree) PushDown(index int, ln, rn int) {
	// ln,rn为左右子树的数字数量
	if rt.Add[index] != 0 {
		rt.Add[index<<1] += rt.Add[index]
		rt.Add[index<<1|1] += rt.Add[index]
		rt.Sum[index<<1] += rt.Add[index] * ln
		rt.Sum[index<<1+1] += rt.Add[index] * rn
		rt.Add[index] = 0
	}
}

func (rt *rangeTree) RangeQuery(L, R int, l, r int, index int) int {
	if L <= l && r <= R {
		return rt.Sum[index]
	}
	mid := (l + r) >> 1
	rt.PushDown(index, mid-l+1, r-mid)
	res := 0
	if L <= mid {
		res += rt.RangeQuery(L, R, l, mid, index<<1)
	} else {
		res += rt.RangeQuery(L, R, mid+1, r, index<<1|1)
	}
	return res
}
func (rt *rangeTree) RangeQueryNonRec(L, R int) int {
	s, t, ln, rn, x := rt.N+L-1, rt.N+R+1, 0, 0, 1
	res := 0
	for s^t^1 != 0 {
		if rt.Add[s] != 0 {
			res += rt.Add[s] * ln
		}
		if rt.Add[t] != 0 {
			res += rt.Add[t] * rn
		}
		if s&1 == 0 {
			res += rt.Sum[s^1]
			ln += x
		}
		if t&1 != 0 {
			res += rt.Sum[t^1]
			rn += x
		}
		s >>= 1
		t >>= 1
		x <<= 1
	}
	for s != 0 {
		res += rt.Add[s] * ln
		res += rt.Add[t] * rn
		s >>= 1
		t >>= 1
	}
	return res
}

// 并查集
type UnionSet struct {
	parent []int
	// 此时rank[i]定义为包含在节点i下的所有元素个数（包括本身）即树的节点数目，初始化时定义为1
	rank   []int
}

func (u *UnionSet) Init(n int) {
	u.parent = make([]int, n+1)
	for i := 1; i <= n; i++ {
		u.parent[i] = i
	}
}
func (u *UnionSet) InitWithRank(n int) {
	u.parent = make([]int, n+1)
	for i := 1; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 1 // 与上面的rank定义要保持一直
	}
}
func (u *UnionSet) Find(x int) int {
	if u.parent[x] == x {
		return x
	} else {
		return u.Find(u.parent[x])
	}
}

// 路径压缩
func (u *UnionSet) findWithCompress(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for x != root {
		x, u.parent[x] = u.parent[x], root
	}
	return root
}

func (u *UnionSet) Union(i, j int) {
	u.parent[u.Find(i)] = u.Find(j)
}
func (u *UnionSet) UnionWithRank(i, j int) {
	x := u.Find(i)
	y := u.Find(j)
	if x != y {
		// 无论将秩定义成子树高上界，还是子树节点数，此处定义为子树的节点数目
		// 按秩合并都是尝试合出最矮的树，并不保证一定最矮。
		if u.rank[x] >= u.rank[y] {
			u.parent[y] = x
			u.rank[x] += u.rank[y]
		} else {
			u.parent[x] = y
			u.rank[y] += u.rank[x]
		}
	}
}
