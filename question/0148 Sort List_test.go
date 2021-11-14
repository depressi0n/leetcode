package question

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
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
			args: args{head: makeList([]int{4, 2, 1, 3})},
			want: makeList([]int{1, 2, 3, 4}),
		},
		{
			name: "Test2",
			args: args{head: makeList([]int{-1, 5, 3, 4, 0})},
			want: makeList([]int{-1, 0, 3, 4, 5}),
		},
		{
			name: "Test3",
			args: args{head: makeList([]int{-1, 5, 3, 4, 0, 8})},
			want: makeList([]int{-1, 0, 3, 4, 5, 8}),
		},
		{
			name: "Test4",
			args: args{head: makeList([]int{-1, 5, 3, 4, 0, 8, 6})},
			want: makeList([]int{-1, 0, 3, 4, 5, 6, 8}),
		},
		{
			name: "Test5",
			args: args{head: makeList([]int{-1, 5, 3, 4,9, 0, 8, 6})},
			want: makeList([]int{-1, 0, 3, 4, 5, 6, 8, 9}),
		},
		{
			name: "Test7",
			args: args{head: makeList([]int{})},
			want: makeList([]int{}),
		},
		{
			name: "Test8",
			args: args{head: makeList([]int{1})},
			want: makeList([]int{1}),
		},
		{
			name: "Test9",
			args: args{head: makeList([]int{2,1})},
			want: makeList([]int{1,2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
