package question

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{root: makeTree([]int{3, 9, 20, -1, -1, 15, 7}, -1)},
			want: 2,
		},
		{
			name: "Test2",
			args: args{root: makeTree([]int{2, -1, 3, -1, 4, -1, 5, -1, 6}, -1)},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
