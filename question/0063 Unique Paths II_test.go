package question

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{obstacleGrid: [][]int{
				{0,0,0},
				{0,1,0},
				{0,0,0},
			}},
			want:2,
		},
		{
			name:"Test2",
			args: args{obstacleGrid: [][]int{
				{0,1},
				{0,0},
			}},
			want:1,
		},
		{
			name:"Test3",
			args: args{obstacleGrid: [][]int{
				{0},
			}},
			want:1,
		},
		{
			name:"Test4",
			args: args{obstacleGrid: [][]int{
				{0,1,1,1,1,0},
			}},
			want:0,
		},
		{
			name:"Test5",
			args: args{obstacleGrid: [][]int{
				{0,0},
			}},
			want:1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
