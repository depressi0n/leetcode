package question

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{
				p: makeTree([]int{1, 2, 3}, -1),
				q: makeTree([]int{1, 2, 3}, -1),
			},
			want: true,
		},
		{
			name: "Test2",
			args: args{
				p: makeTree([]int{1, 2}, -1),
				q: makeTree([]int{1, -1, 2}, -1),
			},
			want: false,
		},
		{
			name: "Test3",
			args: args{
				p: makeTree([]int{1, 2, 1}, -1),
				q: makeTree([]int{1, 1, 2}, -1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
