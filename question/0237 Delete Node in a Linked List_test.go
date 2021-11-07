package question

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		origin *ListNode
		node *ListNode
	}
	t1 := makeList([]int{4, 1, 5, 9, 2})
	t2 := makeList([]int{4, 1, 5, 9, 2})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test1",
			args: args{
				origin: t1,
				node: t1.Next,
			},
			want: makeList([]int{4, 5, 9, 2}),
		},
		{
			name: "Test2",
			args: args{
				origin: t2,
				node: t2.Next.Next,
			},
			want: makeList([]int{4, 1, 9, 2}),
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if !checkListEqual(tt.args.origin, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", tt.args.node, tt.want)
			}
		})
	}
}
