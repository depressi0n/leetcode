package question

import (
	"reflect"
	"testing"
)

func makeTree116(arr []int, null int) *Node116 {
	if len(arr) == 0 {
		return nil
	}
	root := &Node116{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	q := make([]*Node116, 0, len(arr))
	q = append(q, root)
	cur := 0
	for i := 1; i < len(arr); i += 2 {
		if arr[i] == null {
			q[cur].Left = nil
		} else {
			q[cur].Left = &Node116{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}
			q = append(q, q[cur].Left)
		}
		if i+1 < len(arr) {
			if arr[i+1] == null {
				q[cur].Right = nil
			} else {
				q[cur].Right = &Node116{
					Val:   arr[i+1],
					Left:  nil,
					Right: nil,
				}
				q = append(q, q[cur].Right)
			}
		}
		cur++
	}
	return root
}
func levelTrace0116(root *Node116) []int {
	res := make([]int, 0, 100)
	q := root
	for q != nil {
		var p *Node116
		for q != nil {
			res = append(res, q.Val)
			if p == nil && q.Left != nil {
				p = q.Left
			} else if p == nil && q.Right != nil {
				p = q.Right
			}
			q = q.Next
		}
		res = append(res, -1)
		q = p
	}
	return res
}

func Test_connect(t *testing.T) {
	type args struct {
		root *Node116
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{root: makeTree116([]int{1, 2, 3, 4, 5, 6, 7}, -1)},
			want: []int{1, -1, 2, 3, -1, 4, 5, 6, 7, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect116(tt.args.root); !reflect.DeepEqual(levelTrace0116(tt.args.root), tt.want) {
				t.Errorf("connect116() = %v, want %v", got, tt.want)
			}
		})
	}
}
