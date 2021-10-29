package question

import (
	"math/bits"
)

// 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
//返回所有可能的结果。答案可以按 任意顺序 返回。

func removeInvalidParentheses(s string) []string {
	return removeInvalidParenthesesCore3(s)
}

// 回溯+剪枝
func removeInvalidParenthesesCore1(s string) []string {
	isValid := func(s string) bool {
		cnt := 0
		for _, ch := range s {
			if ch == '(' {
				cnt++
			} else if ch == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}
	res := make([]string, 0, 100)
	var helper func(s string, start int, leftRemove int, rightRemove int)
	helper = func(s string, start int, leftRemove int, rightRemove int) {
		if leftRemove == 0 && rightRemove == 0 {
			if isValid(s) {
				res = append(res, s)
			}
			return
		}
		for i := start; i < len(s); i++ {
			// 去重，这里其实用的是DFS，连续相同的括号只需要搜索一次即可
			// 其余的情况都在去掉这一个括号之内的helper完成
			if i != start && s[i] == s[i-1] {
				continue
			}
			// 剪枝
			if leftRemove+rightRemove > len(s)-1 {
				return
			}
			// 尝试去掉一个括号
			if leftRemove > 0 && s[i] == '(' {
				helper(s[:i]+s[i+1:], i, leftRemove-1, rightRemove)
			}
			if rightRemove > 0 && s[i] == ')' {
				helper(s[:i]+s[i+1:], i, leftRemove, rightRemove-1)
			}
		}
	}
	leftRemove, rightRemove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			leftRemove++
		} else if ch == ')' {
			if leftRemove == 0 {
				rightRemove++
			} else {
				leftRemove--
			}
		}
	}
	helper(s, 0, leftRemove, rightRemove)
	return res
}

// BFS，场景明确，并且要求输出所有结果，在进行BFS时每一轮删除字符串的1个括号，直到出现合法串
// 此时最少次数确定，然后将当前层遍历完毕
func removeInvalidParenthesesCore2(s string) []string {
	isValid := func(s string) bool {
		cnt := 0
		for _, ch := range s {
			if ch == '(' {
				cnt++
			} else if ch == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}
	// 保存上一轮搜索结果
	curSet := make(map[string]struct{})
	curSet[s] = struct{}{}
	res := make([]string, 0, 100)
	for {
		for str := range curSet {
			if isValid(str) {
				res = append(res, str)
			}
		}
		// 如果有结果说明结束了
		if len(res) > 0 {
			return res
		}
		// 获得下一层
		nextSet := make(map[string]struct{})
		for str := range curSet {
			for i, ch := range str {
				// 去重
				if i > 0 && byte(ch) == str[i-1] {
					continue
				}
				if ch == '(' || ch == ')' {
					nextSet[str[:i]+str[i+1:]] = struct{}{}
				}
			}
		}
		curSet = nextSet
	}
}

// 得到至少要去掉的左右括号数目，枚举状态子集【一个序列的所有子序列】
func removeInvalidParenthesesCore3(s string) []string {
	// 根据左右子序列进行判断，而不是先恢复出来再判断
	chechValid := func(s string, lmask int, rmask int, left, right []int) bool {
		cnt := 0
		pos1, pos2 := 0, 0
		for i := range s {
			if pos1 < len(left) && i == left[pos1] {
				if lmask>>pos1&1 == 0 {
					cnt++
				}
				pos1++
			} else if pos2 < len(right) && i == right[pos2] {
				if rmask>>pos2&1 == 0 {
					cnt--
					if cnt < 0 {
						return false
					}
				}
				pos2++
			}
		}
		return cnt == 0
	}
	recover := func(s string, lmask int, rmask int, left, right []int) string {
		res := []rune{}
		pos1, pos2 := 0, 0
		for i, ch := range s {
			if pos1 < len(left) && i == left[pos1] {
				if lmask>>pos1&1 == 0 {
					res = append(res, ch)
				}
				pos1++
			} else if pos2 < len(right) && i == right[pos2] {
				if rmask>>pos2&1 == 0 {
					res = append(res, ch)
				}
				pos2++
			} else {
				res = append(res, ch)
			}
		}
		return string(res)
	}
	var left, right []int
	leftRemove, rightRemove := 0, 0
	for i, ch := range s {
		if ch == '(' {
			left = append(left, i)
			leftRemove++
		} else if ch == ')' {
			right = append(right, i)
			if leftRemove == 0 {
				rightRemove++
			} else {
				leftRemove--
			}
		}
	}
	// 得到所有的子序列
	var mask1, mask2 []int
	for i := 0; i < 1<<len(left); i++ {
		if bits.OnesCount(uint(i)) == leftRemove {
			mask1 = append(mask1, i)
		}
	}
	for i := 0; i < 1<<len(right); i++ {
		if bits.OnesCount(uint(i)) == rightRemove {
			mask2 = append(mask2, i)
		}
	}
	// 组合所有的左括号和右括号子序列
	resM := make(map[string]struct{})
	for _, t1 := range mask1 {
		for _, t2 := range mask2 {
			if chechValid(s, t1, t2, left, right) {
				resM[recover(s, t1, t2, left, right)] = struct{}{}
			}
		}
	}
	res := make([]string, 0,len(resM))
	for str := range resM {
		res = append(res, str)
	}
	return res
}

// 主要思想：首先确定最小数量的无效括号是多少
// 然后通过删除来确定结果
// 第一步：去掉左边的右括号和右边的左括号，这里涉及到一些字母
// 第二步：确定要去除多少个，但是这里肯定是去除相同的时候最少
// 第三步：DFS去除，最后缓存结果
func removeInvalidParenthesesCore4(s string) []string {
	// 首先去掉最左边的非左括号和最右边的非右括号
	tmps := ""
	start := 0
	for ; start < len(s); start++ {
		if s[start] == '(' {
			tmps += s[start:]
			break
		}
		if s[start] != ')' {
			tmps = tmps + s[start:start+1]
		}
	}
	s, tmps = tmps, ""
	end := len(s) - 1
	for ; end >= 0; end-- {
		if s[end] == ')' {
			tmps = s[:end+1] + tmps
			break
		}
		if s[end] != '(' {
			tmps = s[end:end+1] + tmps
		}
	}
	s = tmps
	if len(s) == 0 {
		return []string{""}
	}
	start, end = 0, len(s)-1
	// stack[i] 在左括号代表1，右括号代表-1时，从左往右的前缀和
	// stack[i] <0 表示右括号过多，必须删除，后续无法修复
	// stack[i] =0 表示左右括号数目刚好匹配，但不一定有效
	// stack[i] >0 表示左括号过多，可以不删除，因为后续多余右括号可能帮助消除
	stack := make([]int, len(s))
	i := 0
	for ; i < len(s) && s[i] != '('; i++ {
	}
	if i == len(s) {
		return []string{s}
	}
	stack[i] = 1

	for i := start + 1; i <= end; i++ {
		if s[i] == '(' {
			stack[i-start] = stack[i-start-1] + 1
		} else if s[i] == ')' {
			stack[i-start] = stack[i-start-1] - 1
			//if stack[i-start] < 0 {
			//	// 有必须删除的右括号，但可能会有很多组合，所以尽量将后面所有的右括号都拿过到这里
			//	for ; i+1 <= end; i++ {
			//		if s[i+1] == '(' {
			//			i++
			//			stack[i] = 1
			//			break
			//		}
			//		if s[i+1] == ')' {
			//			stack[i-start+1] = stack[i-start] - 1
			//		} else {
			//			stack[i-start+1] = stack[i-start]
			//		}
			//	}
			//}
		} else {
			stack[i-start] = stack[i-start-1]
		}
	}
	resM := make(map[string]struct{}, 100)
	// 开始DFS
	var dfs func(cur int, discard int, cache string)
	dfs = func(cur int, discard int, cache string) {
		if discard == 0 {
			// 直接退出
			cache = cache + s[cur:]
			// 检查，如果合法
			left := 0
			flag := true
			for i := 0; i < len(cache); i++ {
				if cache[i] == '(' {
					left++
				} else if cache[i] == ')' {
					left--
				}
				if left < 0 {
					flag = false
					break
				}
			}
			if !flag || left != 0 {
				return
			}
			if _, ok := resM[cache]; !ok {
				resM[cache] = struct{}{}
			}
			return
		}
		if cur >= len(s) {
			return
		}
		if s[cur] == '(' && discard > 0 {
			dfs(cur+1, discard-1, cache)
		} else if s[cur] == ')' && discard < 0 {
			dfs(cur+1, discard+1, cache)
		}
		dfs(cur+1, discard, cache+s[cur:cur+1])
	}
	dfs(0, stack[len(s)-1], "")
	res := make([]string, 0, len(resM))
	for t := range resM {
		res = append(res, t)
	}
	return res
}
