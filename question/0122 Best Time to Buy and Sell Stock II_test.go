package question

import "testing"

func Test_maxProfit0122(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 7,
		},
		{
			name: "Test2",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit0122(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit0122() = %v, want %v", got, tt.want)
			}
		})
	}
}
