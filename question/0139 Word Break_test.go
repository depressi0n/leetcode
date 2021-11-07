package question

import "testing"

func Test_wordBreak0139(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test1",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want:true,
		},
		{
			name:"Test2",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want:true,
		},
		{
			name:"Test3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want:false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak0139(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak0139() = %v, want %v", got, tt.want)
			}
		})
	}
}
