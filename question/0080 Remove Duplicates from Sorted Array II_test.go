package question

import "testing"

func Test_removeDuplicatesII(t *testing.T) {
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
			args: args{nums: []int{1, 1, 1, 2, 2, 3}},
			want: 5,
		},
		{
			name: "Test2",
			args: args{nums: []int{10, 0, 1, 1, 1, 1, 2, 3, 3}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesII(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicatesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
