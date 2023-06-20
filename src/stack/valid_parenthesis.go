package stack

func isValid(s string) bool {
	parenthesisStack := []rune{}
	closeToOpenBracketMap := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	for _, currentBracket := range s {
		if lastOpenBracket, ok := closeToOpenBracketMap[currentBracket]; ok {
			lastOpenBracketIndex := len(parenthesisStack) - 1

			if lastOpenBracketIndex >= 0 && parenthesisStack[lastOpenBracketIndex] == lastOpenBracket {
				parenthesisStack = parenthesisStack[:lastOpenBracketIndex]
			} else {
				return false
			}
		} else {
			parenthesisStack = append(parenthesisStack, currentBracket)
		}

	}
	return len(parenthesisStack) == 0
}
