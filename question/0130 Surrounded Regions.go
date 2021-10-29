package question

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
func solve0130(board [][]byte) {
	solve0130Core2(board)
}

// 核心思想：对边缘上所有的'O'进行DFS/BFS，标记沿途能走到的所有的'O'为'Q'
// 最后对board进行一次遍历，将所有的'Q'修改为'O'，所有的'O'修改为'X'
// 这里所用的标记方式是一个flag matrix，没有访问的'O'都要修改为'X'
func solve0130Core1(board [][]byte) {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	lengthX := len(board)
	lengthY := len(board[0])
	flag := make([][]bool, lengthX) //标记是否处理过
	for i := 0; i < lengthX; i++ {
		flag[i] = make([]bool, lengthY)
	}
	// 递归会爆栈，所以需要自己维护一个队列
	type loc struct {
		X int
		Y int
	}
	q := make([]loc, 0)
	var update func(int, int)
	update = func(i int, j int) {
		//不要更新最外圈
		for _, dir := range dirs {
			if 0 < i+dir[0] && i+dir[0] < lengthX-1 && 0 <= j+dir[1] && j+dir[1] < lengthY-1 && !flag[i+dir[0]][j+dir[1]] && board[i+dir[0]][j+dir[1]] == 'O' {
				q = append(q, loc{
					X: i + dir[0],
					Y: j + dir[1],
				})
			}
		}
		//if i-1 > 0 && !flag[i-1][j] && board[i-1][j] == 'O' {
		//	q = append(q, loc{
		//		X: i - 1,
		//		Y: j,
		//	})
		//}
		////down
		//if i+1 < lengthX-1 && !flag[i+1][j] && board[i+1][j] == 'O' {
		//	q = append(q, loc{
		//		X: i + 1,
		//		Y: j,
		//	})
		//}
		////left
		//if j-1 >= 0 && !flag[i][j-1] && board[i][j-1] == 'O' {
		//	q = append(q, loc{
		//		X: i,
		//		Y: j - 1,
		//	})
		//}
		////right
		//if j+1 < lengthY-1 && !flag[i][j+1] && board[i][j+1] == 'O' {
		//	q = append(q, loc{
		//		X: i,
		//		Y: j + 1,
		//	})
		//}
		return
	}

	//初始化
	for i := 0; i < lengthY; i++ {
		if board[0][i] == 'O' {
			q = append(q, loc{
				X: 0,
				Y: i,
			})
		}
		if board[lengthX-1][i] == 'O' {
			q = append(q, loc{
				X: lengthX - 1,
				Y: i,
			})
		}
	}
	for i := 0; i < lengthX; i++ {
		if board[i][0] == 'O' {
			q = append(q, loc{
				X: i,
				Y: 0,
			})
		}
		if board[i][lengthY-1] == 'O' {
			q = append(q, loc{
				X: i,
				Y: lengthY - 1,
			})
		}
	}
	//开始往里面推进，一共需要（len(board)+1）/2次
	for len(q) != 0 {
		x := q[0].X
		y := q[0].Y
		if !flag[x][y] {
			flag[x][y] = true
			update(q[0].X, q[0].Y)
		}
		q = q[1:]
	}
	for i := 0; i < lengthX; i++ {
		for j := 0; j < lengthY; j++ {
			if !flag[i][j] && board[i][j] != 'X' {
				board[i][j] = 'X'
			}
		}
	}
	//fmt.Println(board)
	return
}
func solve0130Core2(board [][]byte) {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	lengthX := len(board)
	lengthY := len(board[0])
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		for _, dir := range dirs {
			board[i][j] = 'Q'
			if 0 <= i+dir[0] && i+dir[0] < lengthX && 0 <= j+dir[1] && j+dir[1] < lengthY && board[i+dir[0]][j+dir[1]] == 'O' {
				dfs(i+dir[0], j+dir[1])
			}
		}
	}
	// 对最外圈的'O'进行DFS
	for i := 0; i < lengthX; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][lengthY-1] == 'O' {
			dfs(i, lengthY-1)
		}
	}
	for i := 0; i < lengthY; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[lengthX-1][i] == 'O' {
			dfs(lengthX-1, i)
		}
	}
	// 最后修改
	for i := 0; i < lengthX; i++ {
		for j := 0; j < lengthY; j++ {
			if board[i][j] == 'Q' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

// TODO 用并查集
