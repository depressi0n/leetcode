package question

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{root: makeTree([]int{1, -1, 2, 3}, -1)},
			want: []int{3, 2, 1},
		},
		{
			name: "Test2",
			args: args{root: makeTree([]int{}, -1)},
			want: []int{},
		},
		{
			name: "Test3",
			args: args{root: makeTree([]int{1}, -1)},
			want: []int{1},
		},
		{
			name: "Test4",
			args: args{root: makeTree([]int{1, 2}, -1)},
			want: []int{2, 1},
		},
		{
			name: "Test5",
			args: args{root: makeTree([]int{1, -1, 2}, -1)},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
