package question

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test1",
			args: args{root: makeTree([]int{2,1,3},-1)},
			want:true,
		},
		{
			name:"Test2",
			args: args{root: makeTree([]int{5,1,4,-1,-1,3,6},-1)},
			want:false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
