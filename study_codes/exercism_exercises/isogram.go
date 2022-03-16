package isogram

import "strings"

func IsIsogram(word string) bool {

	word = strings.ToLower(word)

	if word == "" {
		return true
	}

	for i := range word {
		if i != len(word) {
			if string(word[i]) == "-" || string(word[i]) == " " {
				continue
			} else {

				if strings.Count(word, string(word[i])) > 1 {
					return false
				}
			}
		}
	}

	return true
}
