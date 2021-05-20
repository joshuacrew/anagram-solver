package formatter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestAnagramWriter_WriteFormattedAnagrams_WriteOnce(t *testing.T) {
	type args struct {
		words []string
	}
	buffer := bytes.NewBuffer([]byte{})
	tests := []struct {
		name    string
		writer  Formatter
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "write single word - should insert two new lines",
			writer: Formatter{writer: buffer},
			args:   args{words: []string{"hello"}},
			want: `hello

`,
			wantErr: false,
		},
		{
			name:   "write two words - should join with comma and insert two lines",
			writer: Formatter{writer: buffer},
			args:   args{words: []string{"hello", "olelh"}},
			want: `hello,olelh

`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.writer.Print(tt.args.words)

			wrote, _ := ioutil.ReadAll(buffer)
			if string(wrote) != tt.want {
				t.Errorf("Print() wrote = %v, want %v", wrote, tt.want)
			}
		})
	}
}

func TestAnagramWriter_WriteFormattedAnagrams_WriteTwice(t *testing.T) {
	type args struct {
		words []string
	}
	buffer := bytes.NewBuffer([]byte{})
	tests := []struct {
		name            string
		writer          Formatter
		firstWriteArgs  args
		secondWriteArgs args
		want            string
		wantErr         bool
	}{
		{
			name:            "should write two words with a new line between each",
			writer:          Formatter{writer: buffer},
			firstWriteArgs:  args{words: []string{"hello"}},
			secondWriteArgs: args{words: []string{"goodbye"}},
			want: `hello

goodbye

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.writer.Print(tt.firstWriteArgs.words)
			_ = tt.writer.Print(tt.secondWriteArgs.words)

			wrote, _ := ioutil.ReadAll(buffer)
			if string(wrote) != tt.want {
				t.Errorf("Print() wrote = %v, want %v", wrote, tt.want)
			}
		})
	}
}
