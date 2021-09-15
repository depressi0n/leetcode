package question

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams0049(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test1",
			args: args{
				strs: []string{""},
			},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Test2",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "Test3",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams0049(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams0049() = %v, want %v", got, tt.want)
			}
		})
	}
}
