package question

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test1",
			args: args{
				s: "aa",
				p: "a",
			},
			want:false,
		},
		{
			name:"Test2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want:true,
		},
		{
			name:"Test3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want:true,
		},
		{
			name:"Test4",
			args: args{
				s: "aab",
				p: "c*a*b*",
			},
			want:true,
		},
		{
			name:"Test5",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want:false,
		},
		{
			name:"Test6",
			args: args{
				s: "aasdfasdfasdfasdfas",
				p: "aasdf.*asdf.*asdf.*asdf.*s",
			},
			want:true,
		},

	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
