package question

import "testing"

func Test_findMin0153(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{nums: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "Test2",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}},
			want: 0,
		},
		{
			name: "Test3",
			args: args{nums: []int{11, 13, 15, 17}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin0153(tt.args.nums); got != tt.want {
				t.Errorf("findMin0153() = %v, want %v", got, tt.want)
			}
		})
	}
}
