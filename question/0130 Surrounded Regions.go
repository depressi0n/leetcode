package question

// TODO 用并查集
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	lengthX := len(board)
	lengthY := len(board[0])
	//var lengthM int
	//if lengthM<lengthY{
	//	lengthM=lengthX
	//}else{
	//	lengthM=lengthY
	//}
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
		//up
		if i-1 > 0 && !flag[i-1][j] && board[i-1][j] == 'O' {
			q = append(q, loc{
				X: i - 1,
				Y: j,
			})
		}
		//down
		if i+1 < lengthX-1 && !flag[i+1][j] && board[i+1][j] == 'O' {
			q = append(q, loc{
				X: i + 1,
				Y: j,
			})
		}
		//left
		if j-1 >= 0 && !flag[i][j-1] && board[i][j-1] == 'O' {
			q = append(q, loc{
				X: i,
				Y: j - 1,
			})
		}
		//right
		if j+1 < lengthY-1 && !flag[i][j+1] && board[i][j+1] == 'O' {
			q = append(q, loc{
				X: i,
				Y: j + 1,
			})
		}
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
