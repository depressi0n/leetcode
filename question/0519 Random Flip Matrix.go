package question

import "math/rand"

//  借助约瑟夫环的思想，重新编号
type Solution struct {
	rows    int
	cols    int
	total   int
	flipped []int
}

func Constructor0519(n_rows int, n_cols int) Solution {
	res := Solution{
		rows:    n_rows,
		cols:    n_cols,
		total:   n_rows * n_cols,
		flipped: make([]int, 0, 1000),
	}
	return res
}

func (this *Solution) Flip() []int {
	remain := this.total - len(this.flipped)
	// 随机选择一个值
	loc := rand.Intn(remain)
	this.flipped = append(this.flipped, loc)
	// 返回到最顶层
	for i := len(this.flipped) - 2; i >= 0; i-- {
		loc = loc%remain + this.flipped[i] + 1
		remain++
	}
	loc = loc % remain
	return []int{loc / this.cols, loc % this.cols}
}

func (this *Solution) Reset() {
	this.total = this.rows * this.cols
	this.flipped = make([]int, 0, 1000)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
