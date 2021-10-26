package question

import "testing"

func TestIsPossible(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{
				target: []int{9, 3, 5},
			},
			want: true,
		},
		{
			name: "Test2",
			args: args{
				target: []int{1, 1, 1, 2},
			},
			want: false,
		},
		{
			name: "Test3",
			args: args{
				target: []int{8, 3},
			},
			want: true,
		},
		{
			name:"Test4",
			args: args{
				target: []int{1,1000000000},
			},
			want:true,
		},
		{
			name:"Test5",
			args: args{
				target: []int{2,9000000002},
			},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPossible(tt.args.target); got != tt.want {
				t.Errorf("IsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
