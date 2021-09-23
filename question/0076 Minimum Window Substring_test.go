package question

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Test1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want:"BANC",
		},
		{
			name:"Test2",
			args: args{
				s: "a",
				t: "a",
			},
			want:"a",
		},
		{
			name:"Test3",
			args: args{
				s: "a",
				t: "aa",
			},
			want:"",
		},
		{
			name:"Test4",
			args: args{
				s: "aa",
				t: "aa",
			},
			want:"aa",
		},
		{
			name:"Test5",
			args: args{
				s: "bbaa",
				t: "aba",
			},
			want:"baa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
