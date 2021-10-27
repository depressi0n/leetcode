package question

import "testing"

func Test_isPalindrome0125(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "Test2",
			args: args{s: "race a car"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome0125(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome0125() = %v, want %v", got, tt.want)
			}
		})
	}
}
