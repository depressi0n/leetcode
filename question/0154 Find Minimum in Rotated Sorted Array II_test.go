package question

import "testing"

func Test_findMin0154(t *testing.T) {
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
			args: args{nums: []int{1,3,5}},
			want: 1,
		},
		{
			name: "Test2",
			args: args{nums: []int{2,2,2,0,1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin0154(tt.args.nums); got != tt.want {
				t.Errorf("findMin0154() = %v, want %v", got, tt.want)
			}
		})
	}
}
