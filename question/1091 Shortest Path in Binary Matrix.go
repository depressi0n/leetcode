package question

func ShortestPathBinaryMatrix(grid [][]int) int {
	// 定义一个clear path：从top-left至bottom-right的路径，保证被访问的单元格都是0，并且路径相邻的位置是8个方向相邻的
	// 找出最短的路径
	// 思路：进行一次bfs搜索即可
	// 首先检查左上角和右下角的值是不是0，否则直接返回-1
	// 然后从右下角开始往其他方向走，为了路径最短
	if len(grid) == 1 {
		if grid[0][0] == 0 {
			return 1
		}
		return 1
	}
	if grid[0][0] != 0 || grid[len(grid)-1][len(grid)-1] != 0 {
		return -1
	}
	visit := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visit[i] = make([]bool, len(grid))
	}
	// 用dp似乎会进入一个相互依赖的死循环，或许有办法避免这种死循环？
	queue := make([]int, 0, len(grid)*len(grid))
	queue = append(queue, 0)
	visit[0][0] = true
	level := 1
	end := 1
	cur := 0
	for cur < end {
		x := queue[cur] / len(grid)
		y := queue[cur] % len(grid)
		for _, deltaX := range []int{-1, 0, 1} {
			for _, deltaY := range []int{-1, 0, 1} {
				if x+deltaX == len(grid)-1 && y+deltaY == len(grid)-1 {
					return level + 1
				}
				if 0 <= x+deltaX && x+deltaX < len(grid) &&
					0 <= y+deltaY && y+deltaY < len(grid) &&
					grid[x+deltaX][y+deltaY] == 0 &&
					!visit[x+deltaX][y+deltaY] {
					queue = append(queue, (x+deltaX)*len(grid)+y+deltaY)
					visit[x+deltaX][y+deltaY] = true
				}
			}
		}
		cur++
		if cur == end {
			level++
			end = len(queue)
		}
	}
	return -1
}
