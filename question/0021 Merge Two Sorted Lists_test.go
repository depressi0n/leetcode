package question

import (
	"testing"
)

func Test_mergeTwoLists0021(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{
				l1: makeList([]int{1, 2, 4}),
				l2: makeList([]int{1, 3, 4}),
			},
			want:makeList([]int{1,1,2,3,4,4}),
		},
		{
			name: "Test2",
			args: args{
				l1: makeList([]int{}),
				l2: makeList([]int{1}),
			},
			want:makeList([]int{1}),
		},
		{
			name: "Test3",
			args: args{
				l1: makeList([]int{}),
				l2: makeList([]int{}),
			},
			want:makeList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists0021(tt.args.l1, tt.args.l2); !checkListEqual(got, tt.want) {
				t.Errorf("mergeTwoLists0021() = %v, want %v", got, tt.want)
			}
		})
	}
}
