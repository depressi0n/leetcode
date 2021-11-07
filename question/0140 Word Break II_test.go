package question

import (
	"testing"
)

func Test_wordBreak0140(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{
				"cats and dog",
				"cat sand dog",
			},
		},
		{
			name: "Test2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{
				"pineapple pen apple",
				"pine applepen apple",
				"pine apple pen apple",
			},
		},
		{
			name: "Test3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak0140(tt.args.s, tt.args.wordDict)
			if len(got)!=len(tt.want){
				t.Errorf("wordBreak0140() = %v, want %v", got, tt.want)
			}
			m:=make(map[string]struct{})
			for _, s := range tt.want {
				m[s]= struct{}{}
			}
			for _, s := range got {
				if _,ok:=m[s];!ok{
					t.Errorf("wordBreak0140() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
