package question

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	return combineCore3(n, k)
}

//TODO:看懂节省空间的方法，二进制next怎么转换的

func combineCore1(n int, k int) [][]int {
	if k == 0 || n < k {
		return nil
	}
	if n == k {
		res := []int{}
		for i := 0; i < n; i++ {
			res = append(res, i+1)
		}
		return [][]int{res}
	}
	var res [][]int
	var dfs_combine func(pos int, cur []int)
	dfs_combine = func(pos int, cur []int) {
		if len(cur) == k {
			curCopy := make([]int, k)
			for i := 0; i < k; i++ {
				curCopy[i] = cur[i]
			}
			res = append(res, curCopy)
			return
		}
		//剪枝，如果剩下的数不够了，就不进了
		if len(cur)+n-pos < k {
			return
		}
		for i := pos + 1; i <= n; i++ {
			cur = append(cur, i)
			dfs_combine(i, cur)
			cur = cur[:len(cur)-1]
		}
	}
	for i := 1; i <= n-k+1; i++ {
		dfs_combine(i, []int{i})
	}
	return res
}

// TODO:好好理解这个二进制到原始数组的转换过程
// 非常巧妙的构造，需要理解怎么求出Next的
// 首先n个数是否被使用了是通过n bit来表达，所以这里的Next就是怎么从上一个达到下一个
// 规则1：最低位为1，如果末尾有t个连续的1，直接将倒数第t位上的1和倒数第t+1位的0交换即可恶【本质上是规则2的特殊情况】
// 规则2：最低位为0，如果末尾有t个连续的0，而这t个连续的0之前有m个连续的1，可以将倒数第t+m位置上的1与倒数第t+m+1位上的0交换，然后把倒数第t+1位到倒数t+m-1位上的1移动到最低位
// 	如 0110 -> 1001, 1010 -> 1100 , 1011100 -> 1100011
// 至此，得到一个朴素的二进制构造方案
// 为了在原始的数组上直接变换而不是依赖二进制，使用下标来表示
// 从低位向高位找到第一个j是的a[j]+1 != a[j+1] 【因为这个意味着a[j]和a[j+1]对应的二进制位中间有0即这两个1不连续】，将a[j]对应的1推送到高位即a[j]=a[j]+1
func combineCore2(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := make([]int, k+1)
	for i := 0; i < k; i++ {
		temp[i] = i + 1
	}
	temp = append(temp, n+1) //哨兵，保证到第k个就一定会停止下面第二层循环

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}

// 本质上还是一个DFS，通过构造进行了剪枝+一次遍历完成
func combineCore3(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 这种构造挺巧妙的...
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
