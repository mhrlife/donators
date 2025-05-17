package utils

import "strings"

var shouldBeEscaped = "_*[]()~`>#+-=|{}.!"

// EscapeMarkdown escapes special symbols for Telegram MarkdownV2 syntax
func EscapeMarkdown(s string) string {
	var result []rune
	for _, r := range s {
		if strings.ContainsRune(shouldBeEscaped, r) {
			result = append(result, '\\')
		}
		result = append(result, r)
	}
	return string(result)
}
