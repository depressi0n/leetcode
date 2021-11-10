package question

import "testing"

func Test_reorderList(t *testing.T) {
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
			args: args{makeList([]int{1, 2, 3, 4})},
			want: makeList([]int{1, 4, 2, 3}),
		},
		{
			name: "Test2",
			args: args{makeList([]int{1, 2, 3, 4, 5})},
			want: makeList([]int{1, 5, 2, 4, 3}),
		},
		{
			name: "Test3",
			args: args{makeList([]int{1})},
			want: makeList([]int{1}),
		},
		{
			name: "Test4",
			args: args{makeList([]int{1,2})},
			want: makeList([]int{1,2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if !checkListEqual(tt.args.head, tt.want) {
				t.Errorf("reorderList() = ")
				p := tt.args.head
				for p != nil {
					t.Errorf("%d->", p.Val)
					p = p.Next
				}
				t.Errorf("\n")
				t.Errorf("want ")
				p = tt.want
				for p != nil {
					t.Errorf("%d->", p.Val)
					p = p.Next
				}
			}
		})
	}
}
