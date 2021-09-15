package question

import (
	"log"
	"testing"
)

func Test_rotate0048(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1",
			args: args{
				matrix:[][]int{
					{1,2,3},
					{4,5,6},
					{7,8,9},
				},
			},
		},
		{
			name: "Test2",
			args: args{
				matrix:[][]int{
					{5,1,9,11},
					{2,4,8,10},
					{13,3,6,7},
					{15,14,12,16},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.matrix)
			rotate0048(tt.args.matrix)
			log.Println(tt.args.matrix)
		})
	}
}
