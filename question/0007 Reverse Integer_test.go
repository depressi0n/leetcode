package question

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{x: 123},
			want:321,
		},
		{
			name:"Test2",
			args: args{x: -123},
			want:-321,
		},
		{
			name:"Test3",
			args: args{x: 120},
			want:21,
		},
		{
			name:"Test4",
			args: args{x: 0},
			want:0,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
