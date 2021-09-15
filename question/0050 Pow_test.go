package question

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name:"Test1",
			args: args{
				x: 2.00000,
				n: 10,
			},
			want:1024.00000,
		},
		{
			name:"Test2",
			args: args{
				x: 2.10000,
				n: 3,
			},
			want:9.26100,
		},{
			name:"Test3",
			args: args{
				x: 2.00000,
				n: -2,
			},
			want:0.25000,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); math.Abs(got - tt.want)>=0.00001 {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
