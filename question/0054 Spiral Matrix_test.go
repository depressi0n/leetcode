package question

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Test2",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					//{13, 14, 15, 16},
				},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "Test3",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
					//{13, 14, 15, 16},
				},
			},
			want: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			name: "Test4",
			args: args{
				matrix: [][]int{
					{1},
					{2},
					{3},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Test5",
			args: args{
				matrix: [][]int{
					{1, 11},
					{2, 12},
					{3, 13},
					{4, 14},
					{5, 15},
					{6, 16},
					{7, 17},
					{8, 18},
					{9, 19},
					{10, 20},
				},
			},
			want: []int{1,11,12,13,14,15,16,17,18,19,20,10,9,8,7,6,5,4,3,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
