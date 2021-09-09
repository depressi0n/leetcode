package question

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{
				lists: []*ListNode{
					makeList([]int{1, 4, 5}),
					makeList([]int{1, 3, 4}),
					makeList([]int{2,6}),
				},
			},
			want: makeList([]int{1,1,2,3,4,4,5,6}),
		},
		{
			name: "Test2",
			args: args{
				lists: []*ListNode{
					makeList([]int{}),
				},
			},
			want: nil,
		},
		{
			name: "Test3",
			args: args{
				lists: []*ListNode{
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !checkListEqual(got, tt.want) {
				t.Errorf("mergeKListsCore() = %v, want %v", got, tt.want)
			}
		})
	}
}
