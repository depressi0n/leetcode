package question

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{root: makeTree([]int{1,2,3},-1)},
			want:6,
		},
		{
			name:"Test2",
			args: args{root: makeTree([]int{-10,9,20,-1,-1,15,7},-1)},
			want:42,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
