package question

type binaryTreeNode1707 struct {
	// left 表示0， right 表示1
	left, right *binaryTreeNode1707
}

func (root *binaryTreeNode1707) insert(a int) {
	p := root
	for i := 31; i >= 0; i-- { // 32bit，树的高度为33
		if (a>>i)&1 == 0 {
			if p.left == nil {
				p.left = &binaryTreeNode1707{nil, nil}
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &binaryTreeNode1707{nil, nil}
			}
			p = p.right
		}
	}
}
func (root *binaryTreeNode1707) query(a int, m int) int {
	//flag := false
	// 从最高位开始往下走
	var dfs func(cur *binaryTreeNode1707, level int, flag bool, value int) int
	dfs = func(cur *binaryTreeNode1707, level int, flag bool, value int) int {
		if cur == nil {
			return -1
		}
		if level == 0 {
			return value
		}

		if flag || (m>>(level-1))&1 == 1 {
			// 可以走两边
			if (a>>(level-1))&1 == 1 {
				tmp := dfs(cur.left, level-1, true, value|(1<<(level-1)))
				if tmp == -1 {
					// 说明左边走不了，尝试走右边
					return dfs(cur.right, level-1, flag, value)
				}
				return tmp
			} else {
				// 尝试走右边
				tmp := dfs(cur.right, level-1, flag, value|(1<<(level-1)))
				if tmp == -1 {
					// 说明右边走不了，尝试走左边，这时候flag已经为true
					return dfs(cur.left, level-1, true, value)
				}
				return tmp
			}
		} else {
			// 只能走左边
			if (a>>(level-1))&1 == 1 {
				return dfs(cur.left, level-1, flag, value|(1<<(level-1)))
			} else {
				return dfs(cur.left, level-1, flag, value)
			}
		}
	}
	return dfs(root, 32, false, 0)
	//p=root
	//for i:=31;i>=0;i--{
	//	if (level[i]==p.left && (a>>i)&1 ==1) || (level[i]==p.right && (a>>i)&1 ==0) {
	//		res|=1<<i
	//	}
	//	p=level[i]
	//}
}

func MaximizeXor(nums []int, queries [][]int) []int {
	//在nums[j]<=queries[i][1]的前提下， max(nums[j],queries[i][0])
	// 如果所有的元素都大于queries[i][1]，则返回-1
	// 朴素的想法
	// 对于每一个queries[i]，都执行一次query，记录结果在res[i]中
	//res := make([]int, 0, len(queries))
	//query := func(q []int) {
	//	tmp := -1
	//	for i := 0; i < len(nums); i++ {
	//		if nums[i] <= q[1] {
	//			if tmp < nums[i]^q[0] {
	//				tmp = nums[i] ^ q[0]
	//			}
	//		}
	//	}
	//	res = append(res, tmp)
	//}
	//for i := 0; i < len(queries); i++ {
	//	query(queries[i])
	//}
	//return res

	// 可以想象，如果将nums数组中的值按照二进制做成一棵树，那么每次查询的时候都是按照树去查询
	// 对于大小限制的结果，可以在从根节点往下走的过程中设定一个值取指示当前走的结果会怎么样
	// 而且按照二进制树的表达，每次都可以直接得到一个最大的结果
	res := make([]int, 0, len(queries))
	root := &binaryTreeNode1707{
		left:  nil,
		right: nil,
	}
	for i := 0; i < len(nums); i++ {
		// 对nums[i]进行二进制分解,从低到高？ => 必须从高到低，为了与m进行比较！
		root.insert(nums[i])
	}
	for i := 0; i < len(queries); i++ {
		res = append(res, root.query(queries[i][0], queries[i][1]))
	}
	return res
}

// TODO： 看官方题解，有两条优化思路
// （1) 对queries[i][1]进行排序，以及对nums[i]进行排序，然后可以节省一部分时间和空间，因为这样做，并非每个nums[i]都放到了树中，而且每一次查询都是整棵树对查询
// （2）在树中加上一个min字段，表示以该节点为根的子树所记录的最小值是多少，此时可以根据queries[i][1]的值与指定节点的min值相比较
//     如果后者大，说明都会返回-1，进行了剪枝的操作，否则就类似于之前的那种方式，整棵树都会查询一次，循环内部额外判断相反的子节点的min是否超过queries[i][1]，不超过即可转移到该节点
// TODO 100%的题解
//// github.com/EndlessCheng/codeforces-go
//func init() { debug.SetGCPercent(-1) }
//
//type node struct {
//    son [2]*node
//    min int
//}
//type trie struct{ root *node }
//
//func (t *trie) put(v int) *node {
//    o := t.root
//    if v < o.min {
//		o.min = v
//	}
//    for i := 29; i >= 0; i-- {
//        b := v >> i & 1
//        if o.son[b] == nil {
//            o.son[b] = &node{min: math.MaxInt32}
//        }
//        o = o.son[b]
//        if v < o.min {
//            o.min = v
//        }
//    }
//    return o
//}
//
//func (t *trie) maxXorWithLimitVal(v, limit int) (ans int) {
//    o := t.root
//    if o.min > limit {
//        return -1
//    }
//    for i := 29; i >= 0; i-- {
//        b := v >> i & 1
//        if o.son[b^1] != nil && o.son[b^1].min <= limit {
//            ans |= 1 << i
//            b ^= 1
//        }
//        o = o.son[b]
//    }
//    return
//}
//
//func maximizeXor(a []int, qs [][]int) []int {
//    t := &trie{&node{min:math.MaxInt32}}
//    for _, v := range a {
//        t.put(v)
//    }
//    ans := make([]int, len(qs))
//    for i, q := range qs {
//        ans[i] = t.maxXorWithLimitVal(q[0], q[1])
//    }
//    return ans
//}
