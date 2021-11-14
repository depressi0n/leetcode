package question

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{word: "USA"},
			want: true,
		},
		{
			name: "Test2",
			args: args{word: "FlaG"},
			want: false,
		},
		{
			name: "Test3",
			args: args{word: "fLag"},
			want: false,
		},
		{
			name: "Test4",
			args: args{word: "leetcode"},
			want: true,
		},
		{
			name: "Test5",
			args: args{word: "Google"},
			want: true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
