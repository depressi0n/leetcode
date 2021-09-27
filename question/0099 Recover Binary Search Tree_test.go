package question

import (
	"reflect"
	"testing"
)

func Test_recoverTree(t *testing.T) {
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
			args: args{root: makeTree([]int{1, 3, -1, -1, 2}, -1)},
			want: []int{1, 2, 3},
		},
		{
			name: "Test2",
			args: args{root: makeTree([]int{3, 1, 4, -1, -1, 2}, -1)},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
			got := inorderTraversal(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
