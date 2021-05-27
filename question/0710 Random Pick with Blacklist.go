package question

import (
	"math/rand"
)

// 目标：尽可能少调用随机函数，均匀生成黑名单外的[0,n)的随机数
// 如果只有pick操作
// （1）二分查找：
//		白名单中W=N-len(blacklist)，随机选一个k，
//		对排序后的黑名单进行二分查找，找到其中比k小的数有多少个，停止时的low,high，
//		如果为0，则W[k]=k，
//		如果不为0，则k+low+1
// （2）映射
//		在[0,N-len(blacklist))中选一个数，并将黑名单中的数放在数组的后面
//		将黑名单中数分为两个部分，一部分是大于等于N-len(blocklist)的，这部分不需要被映射，一部分是小于的
//		只需要映射第二部分，将大于等于N-len(blocklist)的非黑名单的数与这部分数进行映射
//

// 但是如果还有添加黑名单的操作呢？

type Solution0710 struct {
	// 既然大家都是均匀的
	// 那么将黑名单里的数放置到数组的末尾
	validLength int
	m           map[int]int
}

func Constructor0710(n int, blacklist []int) Solution0710 {
	res := Solution0710{
		validLength: n - len(blacklist),
		m:           map[int]int{},
	}
	help := map[int]struct{}{}
	// 把所有的大于等于n-len(blacklist)的数放到help中
	for i := n - len(blacklist); i < n; i++ {
		help[i] = struct{}{}
	}
	less := make([]int, 0, len(blacklist))
	// 排除掉blacklist中的大于等于n-len(blacklist)的数
	// 剩下的数保留在less中
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] >= res.validLength {
			delete(help, blacklist[i]) //去掉在blacklist中的数
		} else {
			less = append(less, blacklist[i])
		}
	}
	i := 0
	// 完成映射关系
	for k := range help {
		res.m[less[i]] = k
		i++
	}
	return res
}

func (this *Solution0710) Pick() int {
	res := rand.Intn(this.validLength)
	if v, ok := this.m[res]; ok { //选到了一个黑名单中的数，但现在这个位置放的数是v
		return v
	} else {
		return res
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
