package question

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"Test1",
			args: args{
				s:"babad",
			},
			want:[]string{"bab","aba"},
		},
		{
			name:"Test2",
			args: args{
				s:"cbbd",
			},
			want:[]string{"bb"},
		},
		{
			name:"Test3",
			args: args{
				s:"a",
			},
			want:[]string{"a"},
		},
		{
			name:"Test4",
			args: args{
				s:"ac",
			},
			want:[]string{"a","c"},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.args.s)
			i:=0
			for ;i<len(tt.want);i++{
				if  got == tt.want[i] {
					break
				}
			}
			if i == len(tt.want){
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
