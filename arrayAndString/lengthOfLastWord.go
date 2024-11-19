package arrayAndString

import "strings"

func lengthOfLastWord(s string) int {
	count := 0
	symbolHas
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
		}
	}
	return len(s[strings.LastIndex(s, " "):])
}
