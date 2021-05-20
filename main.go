package main

import (
	"flag"
	"log"
	"os"

	"github.com/joshuacrew/anagram-solver/formatter"

	"github.com/joshuacrew/anagram-solver/sectionscanner"

	"github.com/joshuacrew/anagram-solver/anagrams"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file_path", "", "file path of words to find anagrams for")
	flag.Parse()

	file := openFile(filePath)
	defer file.Close()

	scanner := sectionscanner.New(file)
	formatter := formatter.New(os.Stdout)

	for {
		wordsOfSameLength, err := scanner.Scan()
		if err != nil {
			log.Fatalf("failed to read file %s: %v", filePath, err)
		}

		if isEndOfFile(wordsOfSameLength) {
			break
		}

		anagramSet := anagrams.Find(wordsOfSameLength)

		for _, words := range anagramSet {
			err := formatter.Print(words)
			if err != nil {
				log.Fatalf("failed to write words to output %s: %v", words, err)
			}
		}
	}
}

func openFile(filePath string) *os.File {
	if filePath == "" {
		log.Fatal("file path is empty")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", filePath, err)
	}
	return file
}

func isEndOfFile(wordsOfSameLength []string) bool {
	return len(wordsOfSameLength) == 0
}
