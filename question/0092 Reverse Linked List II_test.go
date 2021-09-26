package question

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
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
				m:    2,
				n:    4,
			},
			want:makeList([]int{1,4,3,2,5}),
		},
		{
			name: "Test2",
			args: args{
				head: makeList([]int{5}),
				m:    1,
				n:    1,
			},
			want:makeList([]int{5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
