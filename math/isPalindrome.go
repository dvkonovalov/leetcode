package math

import "strconv"

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	for i := 0; i < len(strX)/2; i++ {
		if strX[i] != strX[len(strX)-i-1] {
			return false
		}
	}
	return true
}
