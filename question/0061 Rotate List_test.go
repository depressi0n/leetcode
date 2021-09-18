package question

import (
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{
				head: makeList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: makeList([]int{4, 5, 1, 2, 3}),
		},
		{
			name: "Test2",
			args: args{
				head: makeList([]int{0, 1, 2}),
				k:    4,
			},
			want: makeList([]int{2, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !checkListEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
