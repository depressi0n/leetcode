package question

import "testing"

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Test1",
			args: args{
				n: 3,
				k: 3,
			},
			want:"213",
		},
		{
			name:"Test2",
			args: args{
				n: 4,
				k: 9,
			},
			want:"2314",
		},
		{
			name:"Test3",
			args: args{
				n: 3,
				k: 1,
			},
			want:"123",
		},
		{
			name:"Test4",
			args: args{
				n: 9,
				k: 7,
			},
			want:"123457689",
		},
		{
			name:"Test5",
			args: args{
				n: 1,
				k: 1,
			},
			want:"1",
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
