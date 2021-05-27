package question

import "runtime/debug"

func findMaximumXOR(nums []int) int {
	// 暴力做法，可能会超时
	//res:=0
	//for i:=0;i<len(nums);i++{
	//	for j:=i+1;j<len(nums);j++{
	//		if nums[i] ^ nums[j] > res{
	//			res=nums[i]^nums[j]
	//		}
	//	}
	//}
	//return res

	//优化
	// nums[i] ^ nums[j] = x  => nums[i]=x^nums[j]
	// 从高到低依次确定x二进制表示的每一位，从而得到x值
	// 可以把每个数表示为长度为31的二进制数，高位补0
	// 从最高位开始，枚举整个数组时看x的该位是否能取到1 【贪心策略】

	// 判断方法1： pre^k(x) 表示x从最高位开始，到第k个二进制为止的数
	// pre^k(nums[j])=pre^k(x) ^ pre^k(nums[i])
	// 将所有的pre^k(nums[j])放入哈希表中，枚举i并计算pre^k(x) ^ pre^k(nums[i])
	// 如果该值出现在哈希表中，则说明第k个二进制可以取到1，否则只能取到0
	//if len(nums)==1{
	//	return nums[0]
	//}
	//if len(nums)==2{
	//	return nums[0]^nums[1]
	//}
	//var res int
	//const highBits = 30
	//for k := highBits; k >=0 ; k-- {
	//	// 记录pre^k(nums[j])
	//	seen:=map[int]bool{}
	//	for _,num:=range nums{
	//		seen[num>>k]=true
	//	}
	//	// 将结果的第k个二进制位置为1
	//	resNext:=2*res +1
	//	found:=false
	//
	//	for _,num :=range nums{
	//		if seen[num>>k ^ resNext]{
	//			found=true
	//			break
	//		}
	//	}
	//	if found{
	//		res =resNext
	//	}else{
	//		// 没有找到，则说明该位只能为0
	//		res = resNext -1
	//	}
	//}
	//return res

	// 判断方法2：将数组元素看成长度为31的01字符串，对应着字典树的查询
	// 从字典树根节点开始遍历，遍历的参照对象为nums[i]，遍历过程中，根据nums[i]的第x个二进制位是0还是1，确定走向的下一个子节点
	// 假定当前遍历到第k个二进制位
	// 如果nums[i]的第k个二进制位是0，应该往表示1的子节点走，如果不存在则只能往0的子节点走，反之亦然
	var res int
	root := &Trie0421{}
	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		res = func(a, b int) int {
			if a < b {
				return b
			}
			return a
		}(res, root.check(nums[i]))
	}
	return res
}

const highBit = 30

type Trie0421 struct {
	left  *Trie0421
	right *Trie0421
}

func (t *Trie0421) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &Trie0421{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &Trie0421{}
			}
			cur = cur.right
		}
	}
}
func (t *Trie0421) check(num int) int {
	var res int
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.right != nil {
				cur = cur.right
				res = 2*res + 1
			} else {
				cur = cur.left
				res = 2 * res
			}
		} else {
			if cur.left != nil {
				cur = cur.left
				res = 2*res + 1
			} else {
				cur = cur.right
				res = 2 * res
			}
		}
	}
	return res
}

//TODO: 看懂下面的算法
func init() { debug.SetGCPercent(-1) }

func findMaximumXOR2(a []int) (ans int) {
	mask := 0
	for i := 30; i >= 0; i-- {
		mask |= 1 << i
		try := ans | 1<<i
		leftPart := map[int]struct{}{a[0] & mask: {}}
		for _, v := range a[1:] {
			if _, has := leftPart[v&mask^try]; has {
				ans = try
				break
			}
			leftPart[v&mask] = struct{}{}
		}
	}
	return
}
