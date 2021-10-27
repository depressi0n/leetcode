package question

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{x: 121},
			want: true,
		},
		{
			name: "Test2",
			args: args{x: -121},
			want: false,
		},
		{
			name: "Test3",
			args: args{x: 10},
			want: false,
		},
		{
			name: "Test4",
			args: args{x: -101},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome0009(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
