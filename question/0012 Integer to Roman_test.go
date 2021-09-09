package question

import "testing"

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Test1",
			args: args{num: 3},
			want:"III",
		},
		{
			name:"Test2",
			args: args{num: 4},
			want:"IV",
		},
		{
			name:"Test3",
			args: args{num: 9},
			want:"IX",
		},
		{
			name:"Test4",
			args: args{num: 58},
			want:"LVIII",
		},
		{
			name:"Test5",
			args: args{num: 1994},
			want:"MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("IntToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
