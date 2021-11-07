package question

import "testing"

func Test_singleNumber(t *testing.T) {
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
			args: args{nums: []int{2, 2, 1}},
			want: 1,
		},
		{
			name: "Test2",
			args: args{nums: []int{4, 1, 2, 1, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber0136(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber0136() = %v, want %v", got, tt.want)
			}
		})
	}
}
