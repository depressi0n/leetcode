package question

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			args: args{
				s:"abcabcbb",
			},
			want:3,
		},
		{
			name:"Test2",
			args: args{
				s:"bbbbb",
			},
			want:1,
		},
		{
			name:"Test3",
			args: args{
				s:"pwwkew",
			},
			want:3,
		},
		{
			name:"Test4",
			args: args{
				s:"",
			},
			want:0,
		},
		{
			name:"Test5",
			args: args{
				s:"au",
			},
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
