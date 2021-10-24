package question

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Test1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			//want: makeTree([]int{0, -3, 9, -10, -1, 5}, -1),
			want: makeTree([]int{0, -10, 5, -1, -3, -1, 9}, -1),
		},
		{
			name: "Test2",
			args: args{nums: []int{0}},
			want: makeTree([]int{0}, -1),
		},
		{
			name: "Test3",
			args: args{nums: []int{0, 1}},
			//want: makeTree([]int{1,0}, -1),
			want: makeTree([]int{0, -1, 1}, -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
