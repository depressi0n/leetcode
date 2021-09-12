package question

import "testing"

func Test_longestValidParentheses(t *testing.T) {
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
			args: args{s: "(()"},
			want:2,
		},
		{
			name:"Test2",
			args: args{s: ")()())"},
			want:4,
		},
		{
			name:"Test3",
			args: args{s: ""},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
