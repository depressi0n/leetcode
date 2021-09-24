package question

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
				head: makeList([]int{1, 2, 3, 3, 4, 4, 5}),
			},
			want: makeList([]int{1, 2, 5}),
		},
		{
			name: "Test2",
			args: args{
				head: makeList([]int{1, 1, 1, 2, 3}),
			},
			want: makeList([]int{2, 3}),
		},
		{
			name: "Test3",
			args: args{
				head: makeList([]int{1, 1, 1}),
			},
			want: nil,
		},
		{
			name: "Test4",
			args: args{
				head: makeList([]int{}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !checkListEqual(tt.want, got) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
