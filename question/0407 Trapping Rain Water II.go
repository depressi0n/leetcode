package question

import (
	"container/heap"
	"math"
)

// 给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。

func trapRainWater(heightMap [][]int) int {
	return trapRainWaterCore3(heightMap)
}

// 主要思想：一层一层看，如果当前层周边都比自己高，则能存储一个单位体积的雨水，否则不能存储
// 从最高层迭代到最底层可以一直迭代到当前层全为1为止，或者到最低底层。
// 复杂度取决于元素值的最大值...但这个复杂度应该可以接受？ => 超时了！
// 优化方向：从底层向高层迭代，如果当前层不再能保存，则上面的层也不能 => 爆栈
// 复杂度太高了
func trapRainWaterCore1(heightMap [][]int) int {
	curMap := make([][]int, len(heightMap))
	maxHeight, minHeight := 0, math.MaxInt32
	for i := 0; i < len(heightMap); i++ {
		curMap[i] = make([]int, len(heightMap[i]))
		for j := 0; j < len(heightMap[i]); j++ {
			maxHeight = max(maxHeight, heightMap[i][j])
			minHeight = min(minHeight, heightMap[i][j])
		}
	}
	type loc struct {
		x int
		y int
	}
	res := 0
	// 检查四周是否封闭，通过对边上的0进行bfs遍历，将其修改为1
	// 剩下的0就是能装水的
	handle := func(m [][]int) int {
		delta := 0
		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		q := make([]loc, 0, len(m))
		for i := 0; i < len(m); i++ {
			if m[i][0] == 0 {
				q = append(q, loc{
					x: i,
					y: 0,
				})
			}
			if m[i][len(m[0])-1] == 0 {
				q = append(q, loc{
					x: i,
					y: len(m[0]) - 1,
				})
			}
		}
		for i := 0; i < len(m[0]); i++ {
			if m[0][i] == 0 {
				q = append(q, loc{
					x: 0,
					y: i,
				})
			}
			if m[len(m)-1][i] == 0 {
				q = append(q, loc{
					x: len(m) - 1,
					y: i,
				})
			}
		}
		for len(q) != 0 {
			m[q[0].x][q[0].y] = 1
			for _, dir := range dirs {
				if 0 <= q[0].x+dir[0] && q[0].x+dir[0] < len(m) && 0 <= q[0].y+dir[1] && q[0].y+dir[1] < len(m[0]) && m[q[0].x+dir[0]][q[0].y+dir[1]] == 0 {
					q = append(q, loc{
						x: q[0].x + dir[0],
						y: q[0].y + dir[1],
					})
				}
			}
			q = q[1:]
		}
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[0]); j++ {
				if m[i][j] == 0 {
					delta++
				}
			}
		}
		return delta
	}
	for cur := minHeight + 1; cur <= maxHeight; cur++ {
		for i := 0; i < len(heightMap); i++ {
			for j := 0; j < len(heightMap[i]); j++ {
				if heightMap[i][j] >= cur {
					curMap[i][j] = 1
				} else {
					curMap[i][j] = 0
				}
			}
		}
		t := handle(curMap)
		res += t
	}
	return res
}

// 优化方向：采用类似单调栈的方式，对每行每列进行处理，各个方向上的木桶效应
// 对于每一个位置，都用单调栈进行处理，求出当前位置能储存水的多少
// water[i][j]表示接水后的高度
// water[i][j]=max(heightMap[i],min(water[i-1][j],water[i+1][j],water[i][j-1],water[i][j+1])
// 实际接水 water[i][j]-heigjtMap[i][j]
// 最外层初始化为heightMap[i][j]
// 木桶效应决定了内层的高度取决最外层的高度，所以从最外层高度的最小值出发BFS，遇到的所有位置最高都只能到这个最小值，在遇到更大值的时候停止
// 此时木桶的最外层已经更新，迭代求出所有方块的water即可求出接水容量即可
// 实际上water可不做定义，在迭代过程中有最小堆来缓存

// 建立一个最小堆
type loc struct {
	h, x, y int
}
type hp0407 []loc

func (h hp0407) Len() int {
	return len(h)
}

func (h hp0407) Less(i, j int) bool {
	return h[i].h < h[j].h
}

func (h hp0407) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp0407) Push(x interface{}) {
	*h = append(*h, x.(loc))
}

func (h *hp0407) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func trapRainWaterCore2(heightMap [][]int) int {
	if len(heightMap) <= 2 || (len(heightMap) > 0 && len(heightMap[0]) <= 2) {
		return 0
	}
	vis := make([][]bool, len(heightMap))
	for i := 0; i < len(heightMap); i++ {
		vis[i] = make([]bool, len(heightMap[0]))
	}
	// 木桶边界
	h := &hp0407{}
	for i, row := range heightMap {
		for j, v := range row {
			if i == 0 || i == len(heightMap)-1 || j == 0 || j == len(heightMap[0])-1 {
				heap.Push(h, loc{
					h: v,
					x: i,
					y: j,
				})
				vis[i][j] = true
			}
		}
	}
	res := 0
	dirs := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(loc)
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if 0 <= nx && nx < len(heightMap) && 0 <= ny && ny < len(heightMap[0]) && !vis[nx][ny] {
				// 如果当前位置可以存储水
				if heightMap[nx][ny] < cur.h {
					res += cur.h - heightMap[nx][ny]
				}
				// 标记
				vis[nx][ny] = true
				// 当前位置的water高度
				heap.Push(h, loc{
					h: max(heightMap[nx][ny], cur.h),
					x: nx,
					y: ny,
				})
			}
		}
	}
	return res
}

// BFS，一直迭代到不再变化（使用队列缓存）
// 假定每个位置接受后的高度均为maxHeight
// 首先最外层肯定不能接水，如果当前接水高度小于相邻的位置的接水高度时，要将其相邻的四个位置调整成当前高度
// 直到 所有位置的接水高度不再变化为止
func trapRainWaterCore3(heightMap [][]int) int {
	if len(heightMap) <= 2 || (len(heightMap) > 0 && len(heightMap[0]) <= 2) {
		return 0
	}
	m, n := len(heightMap), len(heightMap[0])
	maxHeight := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxHeight = max(maxHeight, heightMap[i][j])
		}
	}
	water := make([][]int, m)
	for i := 0; i < m; i++ {
		water[i] = make([]int, n)
		for j := 0; j < n; j++ {
			water[i][j] = maxHeight
		}
	}
	type pair struct {
		x, y int
	}
	q := []pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && heightMap[i][j] < water[i][j] {
				water[i][j] = heightMap[i][j]
				q = append(q, pair{i, j})
			}
		}
	}
	dirs := []int{-1, 0, 1, 0, -1}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y := p.x, p.y
		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i], y+dirs[i+1]
			// 修正
			if 0 <= nx && nx < m && 0 <= ny && ny < n && water[nx][ny] > water[x][y] && water[nx][ny] > heightMap[nx][ny] {
				water[nx][ny] = max(water[x][y], heightMap[nx][ny])
				q = append(q, pair{nx, ny})
			}
		}
	}
	res:=0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res+=water[i][j]-heightMap[i][j]
		}
	}
	return res
}
