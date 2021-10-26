package question

import "testing"

func Test_maxProfit0121(t *testing.T) {
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
			want: 5,
		},
		{
			name: "Test1",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit0121(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit0121() = %v, want %v", got, tt.want)
			}
		})
	}
}
