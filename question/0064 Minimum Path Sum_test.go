package question

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{grid: [][]int{
				{1,3,1},
				{1,5,1},
				{4,2,1},
			}},
			want:7,
		},
		{
			name:"Test2",
			args: args{grid: [][]int{
				{1,2,3},
				{4,5,6},
			}},
			want:12,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
