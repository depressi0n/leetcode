package question

import "testing"

func Test_maxProduct(t *testing.T) {
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
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			name: "Test2",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
		{
			name: "Test3",
			args: args{nums: []int{2, 3, -2, 4, -2, 0, 2, 3, -2, 4, -1}},
			want: 96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
