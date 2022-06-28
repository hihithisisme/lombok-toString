package parser

import "strings"

const openSquare = '['
const closeSquare = ']'

const openRound = '('
const closeRound = ')'

func beginsWithRoundBracket(str string) bool {
	return beginsWithBracket(str, openRound)
}

func beginsWithSquareBracket(str string) bool {
	return beginsWithBracket(str, openSquare)
}

func beginsWithBracket(str string, openBracket int32) bool {
	for _, char := range str {
		if char == openBracket {
			return true
		} else if strings.Contains(SPECIAL_CHARS, string(char)) {
			return false
		}
	}
	return false
}

func endsWithSquareBracket(str string) bool {
	return endsWithBracket(str, closeSquare)
}

func endsWithRoundBracket(str string) bool {
	return endsWithBracket(str, closeRound)
}

func endsWithBracket(str string, closeBracket int32) bool {
	for i := len(str) - 1; i >= 0; i-- {
		char := string(str[i])
		if char == string(closeBracket) {
			return true
		} else if char == "," {
			continue
		} else {
			return false
		}
	}
	return false
}
