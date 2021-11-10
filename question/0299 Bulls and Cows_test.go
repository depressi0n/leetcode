package question

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Test1",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want:"1A3B",
		},
		{
			name:"Test2",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want:"1A1B",
		},
		{
			name:"Test3",
			args: args{
				secret: "1",
				guess:  "0",
			},
			want:"0A0B",
		},
		{
			name:"Test4",
			args: args{
				secret: "1",
				guess:  "1",
			},
			want:"1A0B",
		},
		{
			name:"Test5",
			args: args{
				secret: "11111111111112222221112233331212121213121807",
				guess:  "11111111111112222221112233331212121213127810",
			},
			want:"41A3B",
		},
		{
			name:"Test4",
			args: args{
				secret: "17",
				guess:  "10",
			},
			want:"1A0B",
		},
		{
			name:"Test4",
			args: args{
				secret: "11",
				guess:  "10",
			},
			want:"1A0B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
