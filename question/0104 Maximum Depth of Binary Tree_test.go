package question

import "testing"

func Test_maxDepth(t *testing.T) {
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
			args: args{makeTree([]int{3, 9, 20, -1, -1, 15, 7}, -1)},
			want: 3,
		},
		{
			name: "Test2",
			args: args{makeTree([]int{}, -1)},
			want: 0,
		},
		{
			name: "Test3",
			args: args{makeTree([]int{3}, -1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
