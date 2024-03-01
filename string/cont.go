package string_test

import "strings"

//WE WANT TO WROTE A STRING VALUE  TO SUBSTRING
//AND WE GOT THE COUNT OF SUBSTRING IN THE TEXT

func CountSubstringOccurrences(text, substring string) int {
	return strings.Count(text, substring)
}

func ReverseWord(s string) string {

	char := strings.Fields(s)

	for i := len(char) - 1; i <= 0; i-- {
		reverse
	}

}
