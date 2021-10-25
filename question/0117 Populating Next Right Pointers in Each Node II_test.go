package question

import (
	"reflect"
	"testing"
)

func makeTree117(arr []int, null int) *Node117 {
	if len(arr) == 0 {
		return nil
	}
	root := &Node117{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	q := make([]*Node117, 0, len(arr))
	q = append(q, root)
	cur := 0
	for i := 1; i < len(arr); i += 2 {
		if arr[i] == null {
			q[cur].Left = nil
		} else {
			q[cur].Left = &Node117{
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
				q[cur].Right = &Node117{
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
func levelTrace0117(root *Node117) []int {
	res := make([]int, 0, 100)
	q := root
	for q != nil {
		var p *Node117
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
func Test_connect117(t *testing.T) {
	type args struct {
		root *Node117
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{root: makeTree117([]int{1, 2, 3, 4, 5, -1, 7}, -1)},
			want: []int{1, -1, 2, 3, -1, 4, 5, 7, -1},
		},
		{
			name: "Test2",
			args: args{root: makeTree117([]int{3, 9, 20, -1, -1, 15, 7}, -1)},
			want: []int{3, -1, 9, 20, -1, 15, 7, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connect117(tt.args.root)
			if got := levelTrace0117(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect117() = %v, want %v", got, tt.want)
			}
		})
	}
}
