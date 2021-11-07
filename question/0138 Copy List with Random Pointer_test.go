package question

import (
	"reflect"
	"testing"
)

func makeList0138(nums [][]int) *Node138 {
	m := make(map[int]*Node138)
	for i, num := range nums {
		m[i] = &Node138{
			Val:    num[0],
			Next:   nil,
			Random: nil,
		}
	}
	for i := 0; i < len(nums); i++ {
		m[i].Next = m[i+1]
		m[i].Random = m[nums[i][1]]
	}
	return m[0]
}
func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node138
	}
	tests := []struct {
		name string
		args args
		want *Node138
	}{
		{
			name: "Test1",
			args: args{head: makeList0138([][]int{
				{7, -1},
				{13, 0},
				{11, 4},
				{10, 2},
				{1, 0},
			})},
			want: makeList0138([][]int{
				{7, -1},
				{13, 0},
				{11, 4},
				{10, 2},
				{1, 0},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
