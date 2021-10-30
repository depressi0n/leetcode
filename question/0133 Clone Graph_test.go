package question

import (
	"testing"
)

func makeGraphFromAdjList(adjList [][]int) *Node133 {
	if len(adjList) == 0 {
		return nil
	}
	m := make(map[int]*Node133)
	for i := 0; i < len(adjList); i++ {
		if _, ok := m[i+1]; !ok {
			m[i+1] = &Node133{
				Val:       i + 1,
				Neighbors: nil,
			}
		}
		m[i+1].Neighbors = make([]*Node133, 0, len(adjList[i]))
		for j := 0; j < len(adjList[i]); j++ {
			if _, ok := m[adjList[i][j]]; !ok {
				m[adjList[i][j]] = &Node133{
					Val:       adjList[i][j],
					Neighbors: nil,
				}
			}
			m[i+1].Neighbors = append(m[i+1].Neighbors, m[adjList[i][j]])
		}
	}
	return m[1]
}
func graphDeepEqual(node1 *Node133, node2 *Node133) bool {
	// 重新生成AdjList
	help := func(node *Node133) map[int][]int {
		q := make([]*Node133, 0, 100)
		q = append(q, node)
		cache := make(map[int][]int)
		for len(q) > 0 {
			if _, ok := cache[q[0].Val]; !ok {
				cache[q[0].Val] = make([]int, 0, len(q[0].Neighbors))
			}
			for i := 0; i < len(q[0].Neighbors); i++ {
				cache[q[0].Val] = append(cache[q[0].Val], q[0].Neighbors[i].Val)
				if _, ok := cache[q[0].Neighbors[i].Val]; !ok {
					q = append(q, q[0].Neighbors[i])
				}
			}
			q = q[1:]
		}
		return cache
	}
	if node1==nil && node2==nil{
		return true
	}
	if node1==nil || node2==nil{
		return false
	}
	cache1 := help(node1)
	cache2 := help(node2)
	if len(cache1) != len(cache2) {
		return false
	}
	for v, nodeinone := range cache1 {
		nodeintwo, ok := cache2[v]
		if !ok || len(nodeinone) != len(nodeintwo) {
			return false
		}
		t := map[int]struct{}{}
		for i := 0; i < len(nodeinone); i++ {
			t[nodeinone[i]] = struct{}{}
		}
		for i := 0; i < len(nodeintwo); i++ {
			if _, exist := t[nodeintwo[i]]; !exist {
				return false
			}
		}
	}
	return true
}
func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node133
	}
	tests := []struct {
		name string
		args args
		want *Node133
	}{
		{
			name: "Test1",
			args: args{
				node: makeGraphFromAdjList([][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				}),
			},
			want: makeGraphFromAdjList([][]int{
				{2, 4},
				{1, 3},
				{2, 4},
				{1, 3},
			}),
		},
		{
			name: "Test2",
			args: args{
				node: makeGraphFromAdjList([][]int{
					{},
				}),
			},
			want: makeGraphFromAdjList([][]int{
				{},
			}),
		},
		{
			name: "Test3",
			args: args{
				node: makeGraphFromAdjList([][]int{}),
			},
			want: makeGraphFromAdjList([][]int{}),
		},
		{
			name: "Test4",
			args: args{
				node: makeGraphFromAdjList([][]int{
					{2},
					{1},
				}),
			},
			want: makeGraphFromAdjList([][]int{
				{2},
				{1},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !graphDeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
