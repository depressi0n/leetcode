package question

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	// t1
	t1 := makeList([]int{3, 2, 0, -4})
	tail := t1
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = t1.Next
	// t2
	t2 := makeList([]int{1, 2})
	tail = t2
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = t2

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{head: t1},
			want: t1.Next,
		},
		{
			name: "Test2",
			args: args{head: t2},
			want: t2,
		},
		{
			name: "Test3",
			args: args{head: makeList([]int{1, 2})},
			want: nil,
		},
		{
			name: "Test4",
			args: args{head: makeList([]int{1})},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
