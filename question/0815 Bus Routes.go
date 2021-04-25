package question

import (
	"fmt"
	"math"
)

type Node815 struct {
	Val       int
	next      []*Node815
	nextRoute []int
}

// 直接做会超时
func NumBusesToDestination815_1(routes [][]int, source int, target int) int {
	m := make(map[int]*Node815)
	// 将路线变成图
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			curNode, ok := m[routes[i][j]]
			if !ok {
				curNode = &Node815{
					Val:       routes[i][j],
					next:      make([]*Node815, 0, len(routes[i])),
					nextRoute: make([]int, 0, len(routes)),
				}
				m[curNode.Val] = curNode
			}
			if node, ok := m[routes[i][(j+1)%len(routes[i])]]; ok {
				curNode.next = append(curNode.next, node)
				curNode.nextRoute = append(curNode.nextRoute, i)
			} else {
				newNode := &Node815{
					Val:       routes[i][(j+1)%len(routes[i])],
					next:      make([]*Node815, 0, len(routes[i])),
					nextRoute: make([]int, 0, len(routes)),
				}
				curNode.next = append(curNode.next, newNode)
				curNode.nextRoute = append(curNode.nextRoute, i)
				m[newNode.Val] = newNode
			}
		}
	}
	targetNode, ok := m[target]
	if !ok {
		return -1
	}
	sourceNode, ok := m[source]
	if !ok {
		return -1
	}
	res := math.MaxInt32
	visited := make(map[int]struct{})
	route := make([]int, 0, len(routes))
	length := func(route []int) int {
		used := make(map[int]struct{})
		for i := 0; i < len(route); i++ {
			used[route[i]] = struct{}{}
		}
		return len(used)
	}
	// bfs 会超时
	var bfs func(cur *Node815)
	bfs = func(cur *Node815) {
		if cur == targetNode {
			cnt := length(route)
			if cnt < res {
				res = cnt
			}
			return
		}
		if length(route) >= res {
			return
		}
		visited[cur.Val] = struct{}{}
		for i := 0; i < len(cur.next); i++ {
			if _, ok := visited[cur.next[i].Val]; !ok {
				route = append(route, cur.nextRoute[i])
				bfs(cur.next[i])
				route = route[:len(route)-1]
			}
		}
		delete(visited, cur.Val)
	}
	bfs(sourceNode)
	if res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}

// 还是超时？
func NumBusesToDestination815_2(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	// 建图的时候，直接考虑起点和重点在哪些路线上，只要达到了这些路线，就不用考虑换乘，实质上是一个邻接矩阵，行是线路，列是车站
	//cnt := 0                                      //全部站点
	at := make(map[int]map[int]struct{}) // at[i][j] 表示i在j号路线上
	busBusMap := make(map[int]map[int]struct{})
	//have := make([]map[int]struct{}, len(routes)) // have[i][j]表示第i号路线上有j这个站
	for i := 0; i < len(routes); i++ {
		//have[i] = make(map[int]struct{})
		busBusMap[i] = make(map[int]struct{})
		for j := 0; j < len(routes[i]); j++ {
			//have[i][routes[i][j]] = struct{}{}
			_, ok := at[routes[i][j]]
			// 没有记录过
			if !ok {
				at[routes[i][j]] = make(map[int]struct{})
				//cnt++
			}
			at[routes[i][j]][i] = struct{}{}
		}
	}

	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			for canArrive, _ := range at[routes[i][j]] {
				busBusMap[i][canArrive] = struct{}{}
				busBusMap[canArrive][i] = struct{}{}
			}
		}
	}
	//// 在同一线路上
	//for i := 0; i < len(routes); i++ {
	//	_,ok1 := have[i][target]
	//	_,ok2 := have[i][source]
	//	if ok1 && ok2{
	//		return 1
	//	}
	//}
	// check
	sourceAt, ok := at[source]
	if !ok {
		return -1
	}
	targetAt, ok := at[target]
	if !ok {
		return -1
	}
	//for busNum, _ := range targetAt {
	//	if _, ok := sourceAt[busNum]; ok {
	//		return 0
	//	}
	//}
	visited := make(map[int]struct{})
	// 从 sourceAt的车站所在线路上所有车站开始查询，如果查询到target所在路线，则记录值
	//queue := make([]int, 0, cnt)
	queue := make([]int, 0, len(routes))

	// 将起点所在的路线上所有点入队
	for busNum, _ := range sourceAt {
		//if _, ok := have[busNum][target]; ok {
		//	return 1
		//}

		if _, ok := targetAt[busNum]; ok {
			return 1
		}
		//for i := 0; i < len(routes[busNum]); i++ {
		//	if _, ok := visited[routes[busNum][i]]; !ok {
		//		visited[routes[busNum][i]] = struct{}{}
		//		queue = append(queue, routes[busNum][i])
		//	}
		//}
		visited[busNum] = struct{}{}
		queue = append(queue, busNum)
	}
	end := len(queue)
	cur := 0
	level := 1
	for cur < end {
		//for curBusNum, _ := range at[queue[cur]] {
		//	for i := 0; i < len(routes[curBusNum]); i++ {
		//		//if _, ok := have[curBusNum][target]; ok {
		//		//	return level + 1
		//		//}
		//		if _, ok := visited[routes[curBusNum][i]]; !ok {
		//			visited[routes[curBusNum][i]] = struct{}{}
		//			queue = append(queue, routes[curBusNum][i])
		//		}
		//	}
		//}
		for canArrive, _ := range busBusMap[queue[cur]] {
			if _, ok := targetAt[canArrive]; ok {
				return level + 1
			}
			if _, ok := visited[canArrive]; !ok {
				visited[canArrive] = struct{}{}
				queue = append(queue, canArrive)
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

//TODO：看懂别人写的为什么不超时，用bus作为节点而不是站点作为节点！！
func numBusesToDestination815_3(routes [][]int, S int, T int) int {

	if S == T {
		return 0
	}

	// 包含起点的 bus 和 终点的 bus 集合
	var start, end []int

	for busNum, v1 := range routes {
		for _, v2 := range v1 {
			if v2 == S {
				start = append(start, busNum)
			}
			if v2 == T {
				end = append(end, busNum)
			}
		}
	}
	// 建一个图？用邻接表表示图
	graph := buildGraph(routes)

	// 广度优先搜索找出最短路径
	bfs := func(graph [][]int, s, e int) int {

		// 如果起始公交和结束公交是同一路，则换乘 0 趟
		if s == e {
			return 0
		}

		length := len(graph)

		// 到达每个节点的路径
		weight := make([]int, length)
		weight[s] = -1

		// 节点队列
		queue := make([]int, 1)
		queue[0] = s

		for i := 0; i < length; i++ {

			if len(queue) <= i {
				return -1
			}
			curr := queue[i]

			// 开始遍历
			for j := 0; j < len(graph[curr]); j++ {
				newVertex := graph[curr][j]
				if weight[newVertex] == 0 {
					queue = append(queue, newVertex)
					weight[newVertex] = i + 1

					if newVertex == e {
						return weight[newVertex]
					}
				}
			}

		}

		return -1
	}

	min := -1
	// 从起点所在bus到重点所在bus开始遍历
	for _, s := range start {
		for _, e := range end {
			res := bfs(graph, s, e)
			if min == -1 {
				min = res
			} else if res != -1 && res < min {
				min = res
			}
		}
	}

	if min == -1 {
		return min
	} else {
		return min + 1
	}
}

// 构建邻接表，表示某辆bus和其他的bus能连接起来
func buildGraph(routes [][]int) [][]int {

	busCount := len(routes)

	graph := make([][]int, busCount)

	for i := 0; i < busCount; i++ {
		for j := i; j < busCount; j++ {
			if i != j && edge(routes[i], routes[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	return graph
}

// 二者必须能够通过某个stop相连接起来
func edge(busA []int, busB []int) bool {

	i, j := 0, 0

	for i < len(busA) && j < len(busB) {
		if busA[i] == busB[j] {
			return true
		} else if busA[i] < busB[j] {
			i++
		} else {
			j++
		}
	}

	return false
}

func numBusesToDestination815_4(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	return bfs1(routes, S, T)
}

func bfs1(routes [][]int, S int, T int) int {
	busNum := len(routes)
	busStopMap := make(map[int]map[int]int, busNum) // busStopMap[i][j]表示第i个bus经过第j个站
	busBusMap := make(map[int]map[int]int, busNum)  //
	startBus := make(map[int]int)                   //起点所在的bus序号
	targetBus := make(map[int]int)                  // 重点所在的bus序号
	for i := 0; i < busNum; i++ {
		busStopMap[i] = make(map[int]int)
		busBusMap[i] = make(map[int]int)
		for j := 0; j < len(routes[i]); j++ {
			s := routes[i][j]
			busStopMap[i][s] = 1
			if s == S {
				startBus[i] = 1
			} else if s == T {
				targetBus[i] = 1
			}
		}
	}
	// 特例
	if len(startBus) == 0 || len(targetBus) == 0 {
		return -1
	}
	for i := 0; i < busNum; i++ {
		for j := i + 1; j < busNum; j++ {
			for v, _ := range busStopMap[i] {
				if _, ok := busStopMap[j][v]; ok { //bus和bus能连接
					busBusMap[i][j] = 1
					busBusMap[j][i] = 1
					break
				}
			}
		}
	}

	visit := make([]int, busNum)
	queue := make([]int, 0, busNum)
	for k, _ := range startBus { // 从起点bus出发
		queue = append(queue, k)
		visit[k] = 1
	}

	steps := 0
	arrive := false
	for len(queue) > 0 {
		preLen := len(queue)
		steps++
		for i := 0; i < preLen; i++ {
			bus := queue[0]
			queue = queue[1:]
			if _, ok := targetBus[bus]; ok {
				arrive = true
				return steps
			}
			for b, _ := range busBusMap[bus] {
				if visit[b] == 0 { //迭代下一次能达到的bus
					visit[b] = 1
					queue = append(queue, b)
				}
			}
		}
	}
	if !arrive {
		return -1
	}
	return steps
}

func bfs(routes [][]int, S int, T int) int {
	stopMap := make(map[int]map[int]int)
	stopCntMap := make(map[int]int)

	maxStop := 0
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			k := routes[i][j]
			if k > maxStop {
				maxStop = k
			}
			if _, ok := stopCntMap[k]; !ok {
				stopCntMap[k] = 0
			}
			stopCntMap[k]++
			for _, v := range routes[i] {
				if v == k {
					continue
				}
				if _, ok := stopMap[k]; !ok {
					stopMap[k] = make(map[int]int)
				}
				stopMap[k][v] = 1
			}
		}
	}

	queue := make([]int, 0, len(stopMap)+1)
	queue = append(queue, S)
	visit := make([]int, maxStop+1)
	steps := 0
	arrive := false

	for len(queue) > 0 {
		preLen := len(queue)
		steps++
		fmt.Println(steps, " q: ", queue)
		for i := 0; i < preLen; i++ {
			stop := queue[0]
			queue = queue[1:]
			for k, _ := range stopMap[stop] {
				if k == T {
					arrive = true
					return steps
				}
				if visit[k] != 0 || stopCntMap[k] <= 1 {
					continue
				}
				visit[k] = 1
				queue = append(queue, k)
			}
		}
	}
	if !arrive {
		return -1
	}
	return steps
}
