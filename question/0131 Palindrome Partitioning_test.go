package question

import (
	"reflect"
	"testing"
)

func Test_partition0131(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test1",
			args: args{s: "aab"},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "Test2",
			args: args{s: "a"},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition0131(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition0131() = %v, want %v", got, tt.want)
			}
		})
	}
}
