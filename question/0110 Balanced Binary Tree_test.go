package question

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{root: makeTree([]int{3, 9, 20, -1, -1, 15, 7}, -1)},
			want: true,
		},
		{
			name: "Test2",
			args: args{root: makeTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, -1)},
			want: false,
		},
		{
			name: "Test2",
			args: args{root: makeTree([]int{}, -1)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
