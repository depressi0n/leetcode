package question

import "testing"

func Test_isMatch0044(t *testing.T) {
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
			name: "Test1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "Test2",
			args: args{
				s: "aa",
				p: "*",
			},
			want: true,
		},
		{
			name: "Test3",
			args: args{
				s: "cb",
				p: "?a",
			},
			want: false,
		},
		{
			name: "Test4",
			args: args{
				s: "adceb",
				p: "*a*b",
			},
			want: true,
		},
		{
			name: "Test5",
			args: args{
				s: "acdcb",
				p: "a*c?b",
			},
			want: false,
		},
		{
			name: "Test6",
			args: args{
				s: "",
				p: "",
			},
			want: true,
		},
		{
			name: "Test7",
			args: args{
				s: "abcabczzzde",
				p: "*abc???de*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch0044(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch0044() = %v, want %v", got, tt.want)
			}
		})
	}
}
