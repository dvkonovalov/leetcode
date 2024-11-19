package twoPointers

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	pos := 0
	for _, char := range t {
		if char == int32(s[pos]) {
			pos++

			if pos == len(s) {
				return true
			}
		}
	}

	return false
}
