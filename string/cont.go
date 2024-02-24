package string_test

import "strings"

func CountSubstringOccurrences(text, substring string) int {
	return strings.Count(text, substring)
}
