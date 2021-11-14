package question

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test1",
			args: args{s: "the sky is blue"},
			want: "blue is sky the",
		},
		{
			name: "Test2",
			args: args{s: "  hello world  "},
			want: "world hello",
		},
		{
			name: "Test3",
			args: args{s: "a good   example"},
			want: "example good a",
		},
		{
			name: "Test4",
			args: args{s: "  Bob    Loves  Alice   "},
			want: "Alice Loves Bob",
		},
		{
			name: "Test5",
			args: args{s: "Alice does not even like bob"},
			want: "bob like even not does Alice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
