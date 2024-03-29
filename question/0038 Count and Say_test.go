package question

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Test1",
			args: args{n:1 },
			want:"1",
		},
		{
			name:"Test2",
			args: args{n:2 },
			want:"11",
		},
		{
			name:"Test3",
			args: args{n:4 },
			want:"1211",
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
