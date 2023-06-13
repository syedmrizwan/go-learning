package two_pointers

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	var cleanedString string
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleanedString += string(ch)
		}
	}

	length := len(cleanedString)
	for i := 0; i < length; i++ {
		if cleanedString[i] != cleanedString[length-i-1] {
			return false
		}
	}
	return true
}
