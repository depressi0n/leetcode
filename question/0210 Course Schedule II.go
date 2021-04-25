package question

func findOrder210_1(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	rm := make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		rm[i] = false
	}
	//用map实现依赖关系
	pre := make(map[int]map[int]struct{})
	for i := 0; i < len(prerequisites); i++ {
		m, ok := pre[prerequisites[i][0]]
		if ok {
			m[prerequisites[i][1]] = struct{}{}
		}
		pre[prerequisites[i][0]] = map[int]struct{}{
			prerequisites[i][1]: {},
		}
	}
	for {
		oLen := len(res)
		for k, v := range rm { // v 是一个map，记录所有的依赖
			if v {
				continue
			}
			flag := true
			//检查所有依赖
			for i, _ := range pre[k] {
				if !rm[i] {
					flag = false
					break
				}
			}
			if flag {
				res = append(res, k)
				rm[k] = true
			}
		}
		if len(res) == numCourses {
			return res
		} else if len(res) == oLen {
			return []int{}
		}
	}
}

// TODO:Khan算法
func findOrder210_2(numCourses int, prerequisites [][]int) []int {
	t := make([]int, numCourses)
	d := make([][]int, numCourses)
	for _, v := range prerequisites {
		t[v[0]]++
		d[v[1]] = append(d[v[1]], v[0])
	}
	ans := []int{}
	for i, j := range t {
		if j == 0 {
			ans = append(ans, i)
		}
	}
	for i := 0; i < len(ans); i++ {
		for _, j := range d[ans[i]] {
			t[j]--
			if t[j] == 0 {
				ans = append(ans, j)
			}
		}
	}
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}
