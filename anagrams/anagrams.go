package anagrams

func Find(words []string) map[string][]string {
	if len(words) == 0 {
		return map[string][]string{}
	}

	wordGroups := make(map[string][]string)
	for _, word := range words {
		sortedWord := SortString(word)

		wordGroups[sortedWord] = append(wordGroups[sortedWord], word)
	}

	return wordGroups
}
