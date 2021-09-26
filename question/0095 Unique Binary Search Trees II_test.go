package question

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "Test1",
			args: args{n: 3},
			want: []*TreeNode{
				makeTree([]int{1, -1, 2, -1, 3}, -1),
				makeTree([]int{1, -1, 3, 2}, -1),
				makeTree([]int{2, 1, 3}, -1),
				makeTree([]int{3, 1, -1, -1, 2}, -1),
				makeTree([]int{3, 2, -1, 1}, -1),
			},
		},
		{
			name: "Test2",
			args: args{n: 1},
			want: []*TreeNode{
				makeTree([]int{1}, -1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
