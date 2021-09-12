package question

import (
	"fmt"
	"math/bits"
)

// 题目描述：解数独
// 规则:
// 	数字 1-9 在每一行只能出现一次。
//	数字 1-9 在每一列只能出现一次。
//	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

func solveSudoku(board [][]byte) {
	solveSudokuCore(board)
	for t1 := 0; t1 < 9; t1++ {
		for t2 := 0; t2 < 9; t2++ {
			fmt.Printf("%c ", board[t1][t2])
		}
		fmt.Println()
	}
	fmt.Println()
}

func solveSudokuCore(board [][]byte) {
	// bit set节省空间
	row, col, submatrix := [9]int16{}, [9]int16{}, [9]int16{}
	// 初始化
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			row[i] |= 1 << (board[i][j] - '1')
			col[j] |= 1 << (board[i][j] - '1')
			submatrix[i/3*3+j/3] |= 1 << (board[i][j] - '1')
		}
	}
	var dfs func() bool
	dfs = func() bool {
		// 打印进来的情况
		for t1 := 0; t1 < 9; t1++ {
			for t2 := 0; t2 < 9; t2++ {
				fmt.Printf("%c ", board[t1][t2])
			}
			fmt.Println()
		}
		fmt.Println()
		// 寻找一个位置，尝试填入
		x, y := 0, 0
		for ; x < 9; x++ {
			flag := false
			y = 0
			for ; y < 9; y++ {
				if board[x][y] == '.' {
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
		if x == 9 && y == 9 {
			//if check() {
			for t1 := 0; t1 < 9; t1++ {
				for t2 := 0; t2 < 9; t2++ {
					fmt.Printf("%c ", board[t1][t2])
				}
				fmt.Println()
			}
			fmt.Println()
			return true
			//}
			//return false
		}
		for i := 0; i < 9; i++ {
			if row[x]&(1<<i) == 0 && col[y]&(1<<i) == 0 && submatrix[x/3*3+y/3]&(1<<i) == 0 {
				// 从当前可选值里面选择一个填入
				board[x][y] = byte('1' + i)
				row[x] |= 1 << i
				col[y] |= 1 << i
				submatrix[x/3*3+y/3] |= 1 << i
				// 进入一次
				if dfs() {
					return true
				}
				board[x][y] = '.'
				row[x] ^= 1 << i
				col[y] ^= 1 << i
				submatrix[x/3*3+y/3] ^= 1 << i
			}
		}
		return false
	}
	dfs()
}

func solveSudokuCore2(board [][]byte) {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int
	// 反转
	flip := func(i, j int, digit byte) {
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}
	// 标记已经填入的
	for i, row := range board {
		for j, b := range row {
			if b != '.' {
				digit := b - '1'
				flip(i, j, digit)
			}
		}
	}
	// 将当前能填入的位置上是唯一值的都填入
	for {
		modified := false
		for i, row := range board {
			for j, b := range row {
				if b != '.' {
					continue
				}
				mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3])
				if mask&(mask-1) == 0 { // mask 的二进制表示仅有一个 1
					digit := byte(bits.TrailingZeros(mask))
					flip(i, j, digit)
					board[i][j] = digit + '1'
					modified = true
				}
			}
		}
		// 如果没有了则退出去
		if !modified {
			break
		}
	}
	// 查看当前还有多少需要进行深度优先遍历的
	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}
	// 进行深度优先遍历，相当于穷举，比上一个方法的好处是穷举次数少了很多
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3]) // 0x1ff 即二进制的 9 个 1
		for ; mask > 0; mask &= mask - 1 {                       // 最右侧的 1 置为 0
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
}
