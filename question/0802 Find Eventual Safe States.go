package question

func Answer0802(graph [][]int) []int {
	return eventualSafeNodes2(graph)
}
func eventualSafeNodes3(graph [][]int) []int {
	//三色标记法
	//白色(0)：尚未访问，灰色(1)：在递归栈中或环中，黑色(2)：安全节点
	var res []int
	color:=make([]int,len(graph))
	var safe func(int)bool
	safe= func(i int) bool {
		if color[i]>0{
			return color[i]==2
		}
		color[i]=1
		for _, y := range graph[i] {
			if !safe(y){
				return false
			}
		}
		color[i]=2
		return true
	}
	for i := 0; i < len(graph); i++ {
		if safe(i){
			res=append(res,i)
		}
	}
	return res
}
func eventualSafeNodes2(graph [][]int) []int {
	var res = []int{} // 边反向，拓扑排序
	reverseGraph:=make([][]int,len(graph))
	inDeg:=make([]int,len(graph))
	for x, ys := range graph {
		for _, y := range ys {
			reverseGraph[y]=append(reverseGraph[y],x)
		}
		inDeg[x]=len(ys)
	}
	var queue []int
	for i, d := range inDeg {
		if d == 0{
			queue=append(queue,i)
		}
	}
	for len(queue)>0{
		y:=queue[0]
		queue=queue[1:]
		for _, x := range reverseGraph[y] {
			inDeg[x]--
			if inDeg[x]==0{
				queue=append(queue,x)
			}
		}
	}
	for i, x := range inDeg {
		if x ==0 {
			res=append(res,i)
		}
	}
	return res
}
func eventualSafeNodes1(graph [][]int) []int {
	// 三色标记法的雏形...
	// 存储已经遍历过的点，这样做超时
	safed := make(map[int]bool)
	var trace func(i int, path map[int]struct{}) bool
	trace = func(i int, path map[int]struct{}) bool {
		path[i] = struct{}{}
		cnt := 0
		for j := 0; j < len(graph[i]); j++ {
			cur := graph[i][j]
			if _, ok := path[cur]; ok {
				for index := range path {
					safed[index] = false
				}
				return false
			}

			if _, ok := safed[cur]; ok {
				if safed[cur] {
					cnt++
				}else{
					return false
				}
			}else{
				if trace(cur, path) {
					cnt++
				}
				delete(path, cur)
			}

		}
		if cnt == len(graph[i]) {
			safed[i] = true
			return true
		}
		return false
	}
	for i := 0; i < len(graph); i++ {
		if _, ok := safed[i]; !ok {
			if trace(i, map[int]struct{}{}) {
				safed[i] = true
			} else {
				safed[i] = false
			}
		}
	}
	res := make([]int, 0, len(graph))
	for i := 0; i < len(graph); i++ {
		if safed[i] {
			res = append(res, i)
		}
	}
	return res
}
