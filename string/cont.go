package string_test

import "strings"

//WE WANT TO WROTE A STRING VALUE  TO SUBSTRING
//AND WE GOT THE COUNT OF SUBSTRING IN THE TEXT

func CountSubstringOccurrences(text, substring string) int {
	return strings.Count(text, substring)
}
