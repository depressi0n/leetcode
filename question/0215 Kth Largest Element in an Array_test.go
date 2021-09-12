package question

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want:5,
		},
		{
			name:"Test2",
			args: args{
				nums: []int{3,2,3,1,2,4,5,5,6},
				k:    4,
			},
			want:4,
		},
		{
			name:"Test3",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want:1,
		},
		{
			name:"Test4",
			args: args{
				nums: []int{1,2,3,4,5,6},
				k:    1,
			},
			want:6,
		},
		{
			name:"Test5",
			args: args{
				nums: []int{6,5,4,3,2,1},
				k:    1,
			},
			want:6,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
