package question

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test1",
			args: args{
				root: makeTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}, -1),
				sum:  22,
			},
			want: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
		{
			name: "Test2",
			args: args{
				root: makeTree([]int{1, 2, 3}, -1),
				sum:  5,
			},
			want: [][]int{},
		},
		{
			name: "Test2",
			args: args{
				root: makeTree([]int{1, 2}, -1),
				sum:  0,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
