package question

import (
	"log"
	"math"
	"strconv"
)

type pair0736 struct {
	left, right int
}

var pairs = map[int]pair0736{}

func Answer0736(expression string) int {
	return evaluate(expression)
}
func evaluate(expression string) int {
	// sytax
	// expression -> integer | let expression | add expression | mult expression | an assigned variable
	// integer -> -number |  number
	// (let v1 e1 ... vn en expr)
	// (add e1 e2)
	// (mult e1 e2)
	// 变量小写字母开头，不会是add、mult、let
	// 先计算括号内的值

	// "(let x 2 (add (let x 3 (let x 4 x)) x))"
	// (let x 4 x) => x=4 值为4
	// (let x 3 (let x 4 x)) => (let x 3 4) => x=3 值为4
	// (add (let x 3 (let x 4 x)) x) => (add 4 x) 找本层或者上一层的值 => 6
	// (let x 2 (add (let x 3 (let x 4 x)) x)) => (let x 2 6) => x=2 值为6
	// 思路：首先从左往右扫描括号配对情况，即分离了所有的表达式，可能遇见需要变量值的表达式，所以使用map记录当前层变量的值
	// 用空格分隔字符串， 使用栈结构去记录括号出现

	// 每次执行都清空一次pairs，其实可以用闭包解决
	pairs = map[int]pair0736{}
	stack := make([]int, len(expression))
	for i := 0; i < len(expression); i++ {
		// 对左括号压栈
		if expression[i] == '(' {
			stack = append(stack, i)
		} else if expression[i] == ')' {
			pairs[stack[len(stack)-1]] = pair0736{
				left:  stack[len(stack)-1],
				right: i,
			}
			stack = stack[:len(stack)-1]
		}
	}
	// 根据题意，必然是进入括号的处理
	return handleBracket(expression, 0, pairs[0].right, map[string]int{})
}

// 解决往下走一个，抽出来的目的是很多地方会用到
func nextSymbol(expr string, start int, boundary int) (bool, int, string, int) {
	cur := start
	//for i := 0; i < len(expr); i++ {
	//	if expr[i]==' '{
	//		fmt.Println("index",i)
	//	}else{
	//		fmt.Println(expr[i])
	//	}
	//}
	for cur < boundary {
		if expr[cur] == ' ' {
			break
		}
		cur++
	}
	if 'a' <= expr[start] && expr[start] <= 'z' {
		return false, 0, expr[start:cur], cur + 1
	} else {
		res, err := strconv.Atoi(expr[start:cur])
		if err != nil {
			log.Fatal(err)
		}
		return true, res, "", cur + 1
	}
}

func handleLet(expr string, left, right int, vars map[string]int) int {
	t := make(map[string]int, len(vars))
	for s, v := range vars {
		t[s] = v
	}
	// 期待变量和值交替，最后一个值作为返回值
	// 什么情况下遇到最后一个值呢
	// （1）当计数序号遇到左括号时，肯定是数，不可能是变量，直接返回
	// （2）解析完一个Symbol【可能是值，可能是变量】看看现在的下标，后面还有没有
	cur := left
	for cur <= right {
		var varName string
		var value int
		var flag bool
		// 对应（1）
		if expr[cur] == '(' {
			return handleBracket(expr, cur, pairs[cur].right, t)
		}
		// 解析一个Symbol
		flag, value, varName, cur = nextSymbol(expr, cur, right+1)
		// 后面没有了
		if cur > right {
			if flag {
				return value
			} else {
				return t[varName]
			}
		}
		// 后面还有，可能是括号起来的值
		if expr[cur] == '(' {
			t[varName] = handleBracket(expr, cur, pairs[cur].right, t)
			cur = pairs[cur].right + 2
		} else {
			// 可能是变量和数字
			var name string
			flag, value, name, cur = nextSymbol(expr, cur, right+1)
			if flag {
				t[varName] = value
			} else {
				t[varName] = t[name]
			}
		}
	}
	return -1
}
func handleAdd(expr string, left, right int, vars map[string]int) int {
	// 期待两个值，如果是两个变量则转换成两个值
	v1 := 0
	if expr[left] == '(' {
		v1 = handleBracket(expr, left, pairs[left].right, vars)
		left = pairs[left].right + 2
	} else {
		cur := left
		flag, v, name, cur := nextSymbol(expr, cur, right+1)
		if !flag {
			v1 = vars[name]
		} else {
			v1 = v
		}
		left = cur
	}
	v2 := 0
	if expr[left] == '(' {
		v2 = handleBracket(expr, left, pairs[left].right, vars)
		left = pairs[left].right
	} else {
		cur := left
		flag, v, name, cur := nextSymbol(expr, cur, right+1)
		if !flag {
			v2 = vars[name]
		} else {
			v2 = v
		}
		left = cur
	}
	return v1 + v2

}
func handleMult(expr string, left, right int, vars map[string]int) int {
	// 期待两个值，如果是两个变量则转换成两个值
	v1 := 0
	if expr[left] == '(' {
		v1 = handleBracket(expr, left, pairs[left].right, vars)
		left = pairs[left].right + 2
	} else {
		cur := left
		flag, v, name, cur := nextSymbol(expr, cur, right+1)
		if !flag {
			v1 = vars[name]
		} else {
			v1 = v
		}
		left = cur
	}
	v2 := 0
	if expr[left] == '(' {
		v2 = handleBracket(expr, left, pairs[left].right, vars)
	} else {
		cur := left
		flag, v, name, cur := nextSymbol(expr, cur, right+1)
		if !flag {
			v2 = vars[name]
		} else {
			v2 = v
		}
		left = cur
	}
	return v1 * v2
}
func handleBracket(expr string, left, right int, vars map[string]int) int {
	cur := left + 1
	for ; cur < right && expr[cur] != ' '; cur++ {

	}
	// 可能遇到三种情况
	switch expr[left+1 : cur] {
	case "let":
		return handleLet(expr, cur+1, right-1, vars)
	case "add":
		return handleAdd(expr, cur+1, right-1, vars)
	case "mult":
		return handleMult(expr, cur+1, right-1, vars)
	default:
		log.Fatal("error")
	}
	return math.MaxInt32
}
