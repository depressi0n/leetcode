package question

import (
	"sort"
	"strconv"
	"unicode"
)

func Answer0726(formula string) string {
	return countOfAtoms(formula)
}
func countOfAtoms(formula string) string {
	// 化学式中原子的数量，元素第一个字母总是大写
	// 要考虑的情况有：括号，小写字母，数字，大写字母
	// 首先，所有的元素用一个map进行映射
	// 其次，遇到括号的时候，往后走找到下一个括号
	// 如果括号后面是个数字，则括号内的都要乘以这个数字
	// 结果应该是：首字母排序，第二个字母排序，然后是数字排序的结构

	// 功能十分像编译器里面的parser
	//atom := ""
	//cur := 0
	//// 表示层级,第0层即最外层
	//stack := []int{-1}
	//tmpM := []map[string]int{{}}
	//unsorted := make([]string, 0, len(formula))
	//// TODO：解决嵌套问题的方式很诡异...
	//for cur < len(formula) {
	//	if formula[cur] == '(' {
	//		stack = append(stack, cur)
	//		tmpM = append(tmpM, map[string]int{})
	//		curCntM := tmpM[len(tmpM)-1]
	//		cur++
	//		for cur < len(formula) {
	//			switch {
	//			case 'A' <= formula[cur] && formula[cur] <= 'Z':
	//				atom += formula[cur : cur+1]
	//				cur += 1
	//			case 'a' <= formula[cur] && formula[cur] <= 'z':
	//				atom += formula[cur : cur+1]
	//				cur += 1
	//			case '1' <= formula[cur] && formula[cur] <= '9':
	//				// 往后看，是不是数字，如果是数字，则往上加
	//				num := formula[cur] - '0'
	//				cur++
	//				for cur < len(formula) && '0' <= formula[cur] && formula[cur] <= '9' {
	//					num = num*10 + formula[cur] - '0'
	//					cur++
	//				}
	//				if _, ok := curCntM[atom]; ok {
	//					// 之前添加过
	//					curCntM[atom] += int(num)
	//				} else {
	//					curCntM[atom] = int(num)
	//					atom = ""
	//				}
	//			case formula[cur] == ')':
	//				// 如果后面是数字
	//				num := 0
	//				if '1' <= formula[cur] && formula[cur] <= '9' {
	//					// 往后看，是不是数字，如果是数字，则往上加
	//					num = int(formula[cur] - '0')
	//					cur++
	//					for cur < len(formula) && '0' <= formula[cur] && formula[cur] <= '9' {
	//						num = num*10 + int(formula[cur]-'0')
	//						cur++
	//					}
	//				}
	//				if num == 0 {
	//					num = 1
	//				}
	//				// 当前层所有的东西都要放到上一层去
	//				for atom, cnt := range tmpM[len(stack)-1] {
	//					if _,ok:=tmpM[len(stack)-2][atom];ok{
	//						tmpM[len(stack)-2][atom]+=cnt*num
	//					}else{
	//						unsorted=append(unsorted,atom)
	//						tmpM[len(stack)-2][atom]=cnt*num
	//					}
	//				}
	//				break
	//			}
	//		}
	//		if atom != "" {
	//			if _, ok := tmpM[0][atom]; ok {
	//				// 之前添加过
	//				tmpM[0][atom] += 1
	//			} else {
	//				tmpM[0][atom] = 1
	//				atom = ""
	//			}
	//		}
	//	}else{
	//
	//	}
	//}
	//
	//sort.Slice(unsorted, func(i, j int) bool {
	//	if strings.Compare(unsorted[i], unsorted[j]) == -1 {
	//		return true
	//	} else if strings.Compare(unsorted[i], unsorted[j]) == 1 {
	//		return false
	//	} else {
	//		log.Fatalln("error")
	//		return false
	//	}
	//})
	//res := ""
	//for i := 0; i < len(unsorted); i++ {
	//	res += unsorted[i]
	//	if tmpM[0][unsorted[i]] != 1 {
	//		res += strconv.Itoa(tmpM[0][unsorted[i]])
	//	}
	//}
	//return res

	//从左到右遍历该化学式，并使用哈希表记录当前层遍历到的原子及其数量，因此初始时需将一个空的哈希表压入栈中。对于当前遍历的字符：
	//
	//如果是左括号，将一个空的哈希表压入栈中，进入下一层。
	//如果不是括号，则读取一个原子名称，若后面还有数字，则读取一个数字，否则将该原子后面的数字视作1
	//然后将原子及数字加入栈顶的哈希表中。
	//如果是右括号，则说明遍历完了当前层，若括号右侧还有数字，则读取该数字num
	//否则将该数字视作 1。然后将栈顶的哈希表弹出，将弹出的哈希表中的原子数量与 num 相乘，加到上一层的原子数量中。
	//遍历结束后，栈顶的哈希表即为化学式中的原子及其个数。遍历哈希表，取出所有「原子-个数」对加入数组中，对数组按照原子字典序排序，然后遍历数组，按题目要求拼接成答案。
	i, n := 0, len(formula)
	// 读取一个元素
	parseAtom := func() string {
		start := i
		i++ // 扫描，跳过首字母
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++ // 扫描首字母后的小写字母
		}
		return formula[start:i]
	}
	// 读取一个数字
	parseNum := func() (num int) {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1 // 不是数字，视作 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = num*10 + int(formula[i]-'0') // 扫描数字
		}
		return
	}
	// 用栈解决嵌套问题
	stk := []map[string]int{{}}
	for i < n {
		// 遇到左括号
		if ch := formula[i]; ch == '(' {
			i++
			// 压一个空map
			stk = append(stk, map[string]int{})
		} else if ch == ')' {
			i++
			// 准备出栈，先看右边数字是多少
			num := parseNum()
			atomNum := stk[len(stk)-1]
			// 出栈
			stk = stk[:len(stk)-1]
			// 放到上一层
			for atom, v := range atomNum {
				stk[len(stk)-1][atom] += v * num // 将括号内的原子数量乘上 num，加到上一层的原子数量中
			}
		} else {
			// 正常解析元素
			atom := parseAtom()
			num := parseNum()
			stk[len(stk)-1][atom] += num // 统计原子数量
		}
	}
	// 所有的元素已经统计完
	atomNum := stk[0]
	type pair struct {
		atom string
		num  int
	}
	pairs := make([]pair, 0, len(atomNum))
	for k, v := range atomNum {
		pairs = append(pairs, pair{k, v})
	}
	//根据字符串的大小进行比较，因为首字母肯定是大写字母，所以不会出现婚论
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })
	// 用byte接收，最后转成string
	ans := []byte{}
	for _, p := range pairs {
		ans = append(ans, p.atom...)
		if p.num > 1 {
			ans = append(ans, strconv.Itoa(p.num)...)
		}
	}
	return string(ans)
}
