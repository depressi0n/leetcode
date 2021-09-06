package question

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Test1",
			args:args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want:"PAHNAPLSIIGYIR",
		},
		{
			name:"Test2",
			args:args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want:"PINALSIGYAHRPI",
		},
		{
			name:"Test3",
			args:args{
				s:       "PAYPALISHIRING",
				numRows: 5,
			},
			want:"PHASIYIRPLIGAN",
		},
		{
			name:"Test4",
			args:args{
				s:       "A",
				numRows: 1,
			},
			want:"A",
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
