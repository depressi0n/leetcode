package question

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{
				head: makeList([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: makeList([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			name: "Test2",
			args: args{
				head: makeList([]int{2, 1}),
				x:    2,
			},
			want: makeList([]int{1, 2}),
		},
		{
			name: "Test3",
			args: args{
				head: makeList([]int{}),
				x:    2,
			},
			want: makeList([]int{}),
		},
		{
			name: "Test4",
			args: args{
				head: makeList([]int{2}),
				x:    2,
			},
			want: makeList([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
