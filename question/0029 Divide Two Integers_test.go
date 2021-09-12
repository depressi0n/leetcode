package question

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "Test2",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			name: "Test3",
			args: args{
				dividend: -2147483648,
				divisor:  -3,
			},
			want: 715827882,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
