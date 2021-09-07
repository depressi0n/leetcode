package question

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{s: "42"},
			want:42,
		},
		{
			name:"Test2",
			args: args{s: "   -42"},
			want:-42,
		},
		{
			name:"Test3",
			args: args{s: "4193 with words"},
			want:4193,
		},
		{
			name:"Test4",
			args: args{s: "words and 987"},
			want:0,
		},
		{
			name:"Test5",
			args: args{s: "-91283472332"},
			want:-2147483648,
		},
		{
			name:"Test5",
			args: args{s: ""},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
