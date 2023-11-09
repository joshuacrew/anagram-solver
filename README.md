# Anagram Solver

This program takes a file path as an argument that points to a file that contains one word per line and writes to standard output the words in the file separated by new lines, with the words that are anagrams of each other grouped together.

To run make sure you have Go 1.16 installed or higher and run
```bash
go run main.go -file_path=data/example2.txt
```
where the -file_path argument contains the path of your file to run

## Chosen Language

This program is written in Go as it is really simple to write command line programs in Go, due to the simplicity of the language and the fantastic standard library that it has.

## Big O Analysis

The algorithm determines anagrams by checking if sorted versions of words are identical.

This uses a standard library Quicksort function which has a time complexity of O(N log N), giving an overall time complexity of O(M * (N log N)), where M is the number of words and N is the number of letters of each word.

## Data Structures

A slice is used to collect the words as they are read as it is a dynamically-sized, flexible view of an array, which means it is ideal for gathering unknown length collections of words. A map is used to accumulate the anagrams as map lookups have a time complexity O(1), so it is not computationally intensive to add each anagram as a value to the sorted word key. The value of the map is a slice and appending to the slice has an amortized time complexity of O(1).  

## Future Work

* Investigate and potentially implement a more efficient algorithm that does not involve fully sorting every word.
* Thorough profiling of system resources and performance to ensure resources are used efficiently and performance is satisfactory
* Investigate parallelising the workload across all cores of the machine 
* More comprehensive unit testing
