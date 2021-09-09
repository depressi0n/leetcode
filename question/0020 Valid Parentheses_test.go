package question

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test1",
			args: args{s: "()"},
			want:true,
		},
		{
			name:"Test2",
			args: args{s: "()[]{}"},
			want:true,
		},
		{
			name:"Test3",
			args: args{s: "(]"},
			want:false,
		},
		{
			name:"Test4",
			args: args{s: "([)]"},
			want:false,
		},
		{
			name:"Test5",
			args: args{s: "{[]}"},
			want:true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
