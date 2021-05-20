package anagrams

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "words contains anagrams - group by alphabetalised key a",
			args: args{words: []string{"from", "time", "item", "form", "toff", "test"}},
			want: map[string][]string{
				"fmor": {"from","form"},
				"eimt": {"time","item"},
				"ffot": {"toff"},
				"estt": {"test"},
			},
		},
		{
			name: "words is empty - should return empty map",
			args: args{words: []string{}},
			want: map[string][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
