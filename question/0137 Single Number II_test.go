package question

import "testing"

func Test_singleNumber0137(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Tset1",
			args: args{nums: []int{2,2,3,2}},
			want:3,
		},
		{
			name:"Tset2",
			args: args{nums: []int{0,1,0,1,0,1,99}},
			want:99,
		},

	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber0137(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber0137() = %v, want %v", got, tt.want)
			}
		})
	}
}
