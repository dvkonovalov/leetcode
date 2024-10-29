package math

func plusOne(digits []int) []int {
	nextInc := true
	pos := len(digits) - 1
	for nextInc {
		if pos < 0 {
			return append([]int{1}, digits...)
		}

		digits[pos] += 1
		if digits[pos] > 9 {
			digits[pos] -= 10
		} else {
			nextInc = false
		}
		pos--

	}
	return digits
}
