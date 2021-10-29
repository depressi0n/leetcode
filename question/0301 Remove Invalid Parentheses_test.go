package question

import (
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test1",
			args: args{s: "()())()"},
			want: []string{
				"(())()",
				"()()()",
			},
		},
		{
			name: "Test2",
			args: args{s: "(a)())()"},
			want: []string{
				"(a())()",
				"(a)()()",
			},
		},
		{
			name: "Test3",
			args: args{s: ")("},
			//want: nil,
			want: []string{""},
			//want: []string{},
		},
		{
			name: "Test4",
			args: args{s: "(a)())()))"},
			want: []string{
				"(a(()))",
				"(a()())",
				"(a())()",
				"(a)(())",
				"(a)()()",
			},
		},
		{
			name: "Test5",
			args: args{s: "))(a)())()))(("},
			want: []string{
			"(a(()))",
			"(a()())",
			"(a())()",
			"(a)(())",
			"(a)()()",
			},
		},
		{
			name: "Test6",
			args: args{s: "n"},
			want: []string{
				"n",
			},
		},
		{
			name: "Test7",
			args: args{s: "())(((()m)("},
			want: []string{
				"()(()m)",
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeInvalidParentheses(tt.args.s)
			if len(got)!=len(tt.want){
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)

			}
			m:=make(map[string]struct{})
			for i := 0; i < len(tt.want); i++ {
				m[tt.want[i]]= struct{}{}
			}
			for i := 0; i < len(got); i++ {
				if _,ok:=m[got[i]];!ok{
					t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)

				}
			}
		})
	}
}
