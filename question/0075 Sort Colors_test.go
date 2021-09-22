package question

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
		},
		{
			name: "Test2",
			args: args{
				nums: []int{2, 0, 1},
			},
		},
		{
			name: "Test3",
			args: args{
				nums: []int{0},
			},
		},
		{
			name: "Test4",
			args: args{
				nums: []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
