package sectionscanner

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestSectionScanner_Scan(t *testing.T) {
	const testData = `abc
bac
hello
goodbye
`

	tests := []struct {
		name     string
		scanner  SectionScanner
		overflow string
		want     []string
		wantErr  bool
	}{
		{
			name:     "words of the same length should be returned, first word of next section added to overflowedWord",
			scanner:  *New(strings.NewReader(testData)),
			want:     []string{"abc", "bac"},
			overflow: "hello",
		},
		{
			name: "end of file - should return nil with empty overflowedWord",
			scanner: *New(strings.NewReader(`
`)),
			want:     nil,
			overflow: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.scanner.Scan()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
			if tt.scanner.overflowedWord != tt.overflow {
				t.Errorf("overflowedWord got %v, want %v", tt.scanner.overflowedWord, tt.overflow)
			}
		})
	}
}

func TestSectionScanner_scanSection(t *testing.T) {
	type scanner struct {
		scanner      *bufio.Scanner
		scannedWords []string
		overflow     string
		wordLength   int
	}
	type args struct {
		word string
	}
	tests := []struct {
		name           string
		sectionScanner scanner
		args           args
		want           []string
		sectionDone    bool
	}{
		{
			name: "should scan word input - still in same section",
			sectionScanner: scanner{
				scannedWords: []string{"abc"},
				overflow:     "",
				wordLength:   3,
			},
			args:        args{word: "bac"},
			want:        []string{"abc", "bac"},
			sectionDone: false,
		},
		{
			name: "shouldn't scan any more words - start of next section",
			sectionScanner: scanner{
				scannedWords: []string{"abc"},
				overflow:     "",
				wordLength:   3,
			},
			args:        args{word: "back"},
			want:        []string{"abc"},
			sectionDone: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SectionScanner{
				scanner:           tt.sectionScanner.scanner,
				scannedWords:      tt.sectionScanner.scannedWords,
				overflowedWord:    tt.sectionScanner.overflow,
				sectionWordLength: tt.sectionScanner.wordLength,
			}
			done := s.scanWords(tt.args.word)
			if done != tt.sectionDone {
				t.Errorf("scanWords() done = %v, want %v", done, tt.sectionDone)
			}

			if !reflect.DeepEqual(s.scannedWords, tt.want) {
				t.Errorf("scanWords() want = %v, want %v", tt.want, s.scannedWords)
			}
		})
	}
}
