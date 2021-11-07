package question

import "testing"

func Test_trapRainWater(t *testing.T) {
	type args struct {
		heightMap [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				heightMap: [][]int{
					{1, 4, 3, 1, 3, 2},
					{3, 2, 1, 3, 2, 4},
					{2, 3, 3, 2, 3, 1},
				},
			},
			want: 4,
		},
		{
			name: "Test2",
			args: args{
				heightMap: [][]int{
					{3, 3, 3, 3, 3},
					{3, 2, 2, 2, 3},
					{3, 2, 1, 2, 3},
					{3, 2, 2, 2, 3},
					{3, 3, 3, 3, 3},
				},
			},
			want: 10,
		},
		{
			name: "Test3",
			args: args{
				heightMap: [][]int{
					{1, 0, 1, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "Test4",
			args: args{
				heightMap: [][]int{
					{111, 0, 1, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "Test5",
			args: args{
				heightMap: [][]int{
				{12, 13, 1, 12},
				{13, 4, 13, 12},
				{13, 8, 10, 12},
				{12, 13, 12, 12},
				{13, 13, 13, 13},
				},
			},
			want: 14,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRainWater(tt.args.heightMap); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
