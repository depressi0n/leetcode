package question

func AllCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	// 从某个位置出发bfs过程
	res := make([][]int, 0, R*C)
	level := 0
	queue := make([]int, 0, R*C)
	end := 1
	visited := make(map[int]struct{})
	queue = append(queue, r0*(R+C)+c0)
	visited[r0*(R+C)+c0] = struct{}{}
	var bfs func(cur int)
	bfs = func(cur int) {
		for cur < end {
			x := queue[cur] / (R + C)
			y := queue[cur] % (R + C)
			res = append(res, []int{x, y})
			// up
			if _, ok := visited[(x-1)*(R+C)+y]; x-1 >= 0 && !ok {
				queue = append(queue, (x-1)*(R+C)+y)
				visited[(x-1)*(R+C)+y] = struct{}{}
			}
			// down
			if _, ok := visited[(x+1)*(R+C)+y]; x+1 < R && !ok {
				queue = append(queue, (x+1)*(R+C)+y)
				visited[(x+1)*(R+C)+y] = struct{}{}
			}
			// left
			if _, ok := visited[x*(R+C)+y-1]; y-1 >= 0 && !ok {
				queue = append(queue, x*(R+C)+y-1)
				visited[x*(R+C)+y-1] = struct{}{}
			}
			// _,ok:=visited[(x-1)*(R+C)+y];right
			if _, ok := visited[x*(R+C)+y+1]; y+1 < C && !ok {
				queue = append(queue, x*(R+C)+y+1)
				visited[x*(R+C)+y+1] = struct{}{}
			}
			cur++
			if cur == end {
				level++
				end = len(queue)
			}
		}

	}
	bfs(0)
	return res
}
