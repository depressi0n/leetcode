package question

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 本质上就是一个拓扑排序，转化为图
	m := make(map[int]map[int]struct{}) //指向当前元素
	for i := 0; i < len(prerequisites); i++ {
		if m[prerequisites[i][1]] == nil {
			m[prerequisites[i][1]] = make(map[int]struct{})
		}
		m[prerequisites[i][1]][prerequisites[i][0]] = struct{}{}
	}
	cur := make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		cur[i] = struct{}{}
	}
	for {
		pre := len(cur)
		for i := 0; i < numCourses; i++ {
			_, ok := cur[i]
			if m[i] == nil && ok {
				// 记录少了一个
				delete(cur, i)
				// 去掉与之相关的边
				for j := 0; j < numCourses; j++ {
					delete(m[j], i)
					if len(m[j]) == 0 {
						delete(m, j)
					}
				}
			}
		}
		if pre == len(cur) {
			break
		}
	}
	if len(cur) == 0 {
		return true
	} else {
		return false
	}
}

// TODO: dfs或者bfs都可以，比这种方式高效
