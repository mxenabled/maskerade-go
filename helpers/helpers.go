package helpers

import (
	"strings"
)

func ReplaceCharacterAtIndexInString(str string, replacementToken string, indices []int) string {
	var b strings.Builder

	b.Grow(len(str))

	for i, rune := range str {
		if intInSlice(i, indices) {
			b.WriteString(replacementToken)
		} else {
			b.WriteRune(rune)
		}
	}

	return b.String()
}

func intInSlice(a int, arr []int) bool {
	for _, b := range arr {
		if b == a {
			return true
		}
	}
	return false
}
