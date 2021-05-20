package sectionscanner

import (
	"bufio"
	"io"
)

type SectionScanner struct {
	scanner           *bufio.Scanner
	scannedWords      []string
	overflowedWord    string
	sectionWordLength int
}

func New(r io.Reader) *SectionScanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return &SectionScanner{scanner: scanner}
}

func (s *SectionScanner) Scan() ([]string, error) {
	s.scannedWords = nil

	err := s.scanWordsInSection()
	if err != nil {
		return nil, err
	}

	return s.scannedWords, nil
}

func (s *SectionScanner) scanWordsInSection() error {
	if hasWordOverflowedFromLastSection(s.overflowedWord) {
		s.scannedWords = append(s.scannedWords, s.overflowedWord)
		s.overflowedWord = ""
	}

	for s.scanner.Scan() {
		word := s.scanner.Text()
		sectionDone := s.scanWords(word)
		if sectionDone {
			break
		}
	}
	if err := s.scanner.Err(); err != nil {
		return err
	}
	return nil
}

func hasWordOverflowedFromLastSection(overflowedWord string) bool {
	return overflowedWord != ""
}

func (s *SectionScanner) scanWords(word string) bool {
	if firstIterationOfSection(s) {
		s.sectionWordLength = len(word)
	}

	if sectionFinished(word, s.sectionWordLength) {
		s.overflowedWord = word
		s.sectionWordLength = 0
		return true
	} else {
		s.scannedWords = append(s.scannedWords, word)
	}
	return false
}

func firstIterationOfSection(s *SectionScanner) bool {
	return s.sectionWordLength == 0
}

func sectionFinished(word string, sectionWordLength int) bool {
	return len(word) > sectionWordLength
}
