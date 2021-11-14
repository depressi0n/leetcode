package question

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "Test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "Test3",
			args: args{n: 10},
			want: 16,
		},
		{
			name: "Test4",
			args: args{n: 11},
			want: 18,
		},
		{
			name: "Test5",
			args: args{n: 17},
			want: 38,
		},
		{
			name: "Test6",
			args: args{n: 18},
			want: 42,
		},
		{
			name: "Test7",
			args: args{n: 21},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
