package question

import "testing"

func Test_isSymmetric(t *testing.T) {
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
			args: args{root: makeTree([]int{1, 2, 2, 3, 4, 4, 3}, -1)},
			want: true,
		},
		{
			name: "Test2",
			args: args{root: makeTree([]int{1, 2, 2, -1, 3, -1, 3}, -1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
