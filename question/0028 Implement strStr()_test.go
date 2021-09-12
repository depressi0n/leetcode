package question

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "Test3",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "Test4",
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
