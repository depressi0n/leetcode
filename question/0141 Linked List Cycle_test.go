package question

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	t1 := makeList([]int{3, 2, 0, -4})
	tail := t1
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = t1.Next
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{head: t1},
			want: true,
		},
		{
			name: "Test2",
			args: args{head: makeList([]int{1, 2})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
			}
		})
	}
}
