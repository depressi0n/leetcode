package question

//你正在参与祖玛游戏的一个变种。
//在这个祖玛游戏变体中，桌面上有 一排 彩球，每个球的颜色可能是：红色 'R'、黄色 'Y'、蓝色 'B'、绿色 'G' 或白色 'W' 。你的手中也有一些彩球。
//你的目标是 清空 桌面上所有的球。每一回合：
//从你手上的彩球中选出 任意一颗 ，然后将其插入桌面上那一排球中：两球之间或这一排球的任一端。
//接着，如果有出现 三个或者三个以上 且 颜色相同 的球相连的话，就把它们移除掉。
//如果这种移除操作同样导致出现三个或者三个以上且颜色相同的球相连，则可以继续移除这些球，直到不再满足移除条件。
//如果桌面上所有球都被移除，则认为你赢得本场游戏。
//重复这个过程，直到你赢了游戏或者手中没有更多的球。
//给你一个字符串 board ，表示桌面上最开始的那排球。另给你一个字符串 hand ，表示手里的彩球。请你按上述操作步骤移除掉桌上所有球，计算并返回所需的 最少 球数。如果不能移除桌上所有的球，返回 -1 。

func findMinStep(board string, hand string) int {
	return findMinStepCore3(board, hand)
}

// 首先确认两个不同的事情：第一个结果与选择球的顺序有关系，第二个是要求选择的球数最少
// 选择深度优先遍历的方式，这里为了剪枝，记录了一个已经使用过多少个手中球的计数器，与已知最好的结果的比较
// 会超时的样子，其中一个优化的方向是因为手里的球只有有限种，所以可以用一个map来记录某种颜色的球是否还有剩余，不再是遍历used，而是遍历map，这样可以去掉了重复的选取行为
func findMinStepCore1(board string, hand string) int {
	used := make([]int, len(hand))
	cnt := 1
	res := -1
	var dfs func(cur string)
	dfs = func(cur string) {
		// 消除完成
		if len(cur) == 0 && (res == -1 || res > cnt-1) {
			res = cnt - 1
			return
		}
		// 已经用完了手里的球
		if cnt == len(hand)+1 && len(cur) != 0 {
			return
		}
		// 剪枝，当前已经选择的球数比已知最少的球数还要多
		if res != -1 && cnt-1 > res {
			return
		}
		// 选择一个球
		for i := 0; i < len(used); i++ {
			// 没有使用过
			if used[i] == 0 {
				used[i] = cnt
				cnt++
				// 选择一个位置插入
				for j := 0; j <= len(cur); j++ {
					tmp := cur[:j] + string([]byte{hand[i]}) + cur[j:]
					// 持续做消除操作
					for k := 0; k < len(tmp); k++ {
						nextDiff := k + 1
						for nextDiff < len(tmp) && tmp[nextDiff] == tmp[k] {
							nextDiff++
						}
						// 消除
						if nextDiff-k >= 3 {
							tmp = tmp[:k] + tmp[nextDiff:]
							// 消除完成后需要重新检查，因为可能造成连续消除
							k = -1
						}
					}
					dfs(tmp)
				}
				used[i] = 0
				cnt--
			}
		}
	}
	dfs(board)
	return res
}

// 依旧会超时，下一个优化方向：不再是盲目的去随机插入，而是选择插入到尽可能消除的地方
func findMinStepCore2(board string, hand string) int {
	m := make(map[rune]int)
	for _, b := range hand {
		m[b]++
	}
	cnt := 1
	res := -1
	var dfs func(cur string)
	dfs = func(cur string) {
		// 消除完成
		if len(cur) == 0 && (res == -1 || res > cnt-1) {
			res = cnt - 1
			return
		}
		// 已经用完了手里的球
		if cnt == len(hand)+1 && len(cur) != 0 {
			return
		}
		// 剪枝，当前已经选择的球数比已知最少的球数还要多
		if res != -1 && cnt-1 > res {
			return
		}
		// 选择一个球
		for r, have := range m {
			if have > 0 {
				m[r] = have - 1
				cnt++
				// 选择一个位置插入
				for j := 0; j <= len(cur); j++ {
					tmp := cur[:j] + string([]rune{r}) + cur[j:]
					// 持续做消除操作
					for k := 0; k < len(tmp); k++ {
						nextDiff := k + 1
						for nextDiff < len(tmp) && tmp[nextDiff] == tmp[k] {
							nextDiff++
						}
						// 消除
						if nextDiff-k >= 3 {
							tmp = tmp[:k] + tmp[nextDiff:]
							// 消除完成后需要重新检查，因为可能造成连续消除
							k = -1
						}
					}
					dfs(tmp)
				}
				m[r] = have
				cnt--
			}
		}
	}
	dfs(board)
	return res
}

func findMinStepCore3(board string, hand string) int {
	m := make(map[rune]int)
	for _, b := range hand {
		m[b]++
	}
	cnt := 0
	res := -1
	// 检查当前位置是否可继续删除
	var removeDup func(s string) string
	removeDup = func(s string) string {
		for i := 0; i < len(s); i++ {
			left, right := i, i
			for left-1 >= 0 && s[left-1] == s[i] {
				left--
			}
			for right+1 < len(s) && s[right+1] == s[i] {
				right++
			}
			if right-left+1 >= 3 {
				s = s[:left] + s[right+1:]
				i = -1
			}
		}
		return s
	}
	var dfs func(board string)
	dfs = func(board string) {
		// 先递归消除再检查?这是不对的，遇到"RRRRBBRR"的时候，就把前面的4个R全部删除，但是后面的2个R是可能可以跟着一起删除的如果手里有B的话
		// 应该考虑两种情况，一种是先递归删除后进入，一种是直接进入，但是这并不符合题意，玩家无法控制是否删除
		// 题设有一定问题。
		board = removeDup(board)
		// 消除完成
		if len(board) == 0 {
			if res == -1 || res > cnt {
				res = cnt
			}
			return
		}
		if res != -1 && cnt > res {
			return
		}
		for i := 0; i < len(board); i++ {
			j := i
			for i+1 < len(board) && board[i+1] == board[j] {
				i++
			}
			need := 3 - (i - j + 1)
			if need > 0 && m[rune(board[i])] >= need {
				m[rune(board[i])] -= need
				cnt += need
				dfs(board[:j] + board[i+1:])
				//if !reflect.DeepEqual(removeDup(board[:j] + board[i+1:]),board[:j] + board[i+1:]){
				//	dfs(removeDup(board[:j] + board[i+1:]))// 经过递归删除后再进入，但是这样可能会超时
				//}
				cnt -= need
				m[rune(board[i])] += need
			}
		}
	}
	dfs(board)
	return res
}
