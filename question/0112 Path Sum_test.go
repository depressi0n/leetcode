package question

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{
				root: makeTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, -1),
				sum:  22,
			},
			want: true,
		},
		{
			name: "Test2",
			args: args{
				root: makeTree([]int{1, 2, 3}, -1),
				sum:  5,
			},
			want: false,
		},
		{
			name: "Test3",
			args: args{
				root: makeTree([]int{1, 2}, -1),
				sum:  0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
