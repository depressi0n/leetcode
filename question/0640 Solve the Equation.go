package question

import "fmt"

// 考虑第一个数就是负数
// 考虑第一个数就是x项，而且是负数
// 考虑x的系数为0的情况
// 最后结果如果是小数如何处理？
func SolveEquation(equation string) string {
	leftX := 0
	leftValue := 0
	curX := 0
	curValue := 0
	flag := true // 表示正号
	for i := 0; i < len(equation); {
		switch equation[i] {
		case '+':
			// 记录flag
			flag = true
			i++
		case '-':
			// 记录flag
			flag = false
			i++
		case '=':
			// 转向右边
			leftValue = curValue
			leftX = curX
			curValue = 0
			curX = 0
			i++
		default:
			// 默认是数字和字母x
			index := i
			cur := 0
			for i < len(equation) && '0' <= equation[i] && equation[i] <= '9' {
				// 读取数字
				cur = cur*10 + int(equation[i]-'0')
				i++
			}
			if i >= len(equation) {
				if flag {
					curValue += cur
				} else {
					curValue -= cur
				}
				break
			}
			// 检查之后是不是x，结合前面的flag
			if equation[i] == 'x' {
				// 如果在这个位置没动，说明前面没有系数，默认为1
				// 否则就说明系数就是0
				if index == i {
					cur = 1
				}
				if flag {
					curX += cur
				} else {
					curX -= cur
				}
				flag = true
				i++
			} else {
				if flag {
					curValue += cur
				} else {
					curValue -= cur
				}
				flag = true
			}
		}
	}
	if leftX == curX {
		if leftValue != curValue {
			return "No solution"
		} else {
			return "Infinite solutions"
		}
	} else {
		return fmt.Sprintf("x=%d", (curValue-leftValue)/(leftX-curX))
	}
}
