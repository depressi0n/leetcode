package question

import (
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name:"Test1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want:[][]int{{7},{2,2,3}},
		},
		{
			name:"Test2",
			args: args{
				candidates: []int{2, 3, 5},
				target:    8,
			},
			want:[][]int{{2,2,2,2},{2,3,3},{3,5}},
		},
		{
			name:"Test3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want:[][]int{},
		},
		{
			name:"Test4",
			args: args{
				candidates: []int{1},
				target:     1,
			},
			want:[][]int{{1}},
		},
		{
			name:"Test5",
			args: args{
				candidates: []int{1},
				target:     2,
			},
			want:[][]int{{1,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); len(got)!=len(tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
