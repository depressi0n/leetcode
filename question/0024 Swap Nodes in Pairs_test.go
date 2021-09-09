package question

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{
				head: makeList([]int{1,2,3,4}),
			},
			want: makeList([]int{2, 1, 4, 3}),
		},
		{
			name: "Test2",
			args: args{
				head: makeList([]int{}),
			},
			want: nil,
		},
		{
			name: "Test3",
			args: args{
				head: makeList([]int{1}),
			},
			want: makeList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !checkListEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
