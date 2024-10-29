package arrayAndString

import "strings"

func intToRoman(num int) string {
	ans := ""
	ans = ans + strings.Repeat("M", num/1000)
	num %= 1000
	if num/900 > 0 {
		num %= 900
		ans = ans + "CM"
	}

	ans = ans + strings.Repeat("D", num/500)
	num %= 500
	if num/400 > 0 {
		num %= 400
		ans = ans + "CD"
	}

	ans = ans + strings.Repeat("C", num/100)
	num %= 100
	if num/90 > 0 {
		num %= 90
		ans = ans + "XC"
	}

	ans = ans + strings.Repeat("L", num/50)
	num %= 50
	if num/40 > 0 {
		num %= 40
		ans = ans + "XL"
	}

	ans = ans + strings.Repeat("X", num/10)
	num %= 10
	if num/9 > 0 {
		num %= 9
		ans = ans + "IX"
	}

	ans = ans + strings.Repeat("V", num/5)
	num %= 5
	if num/4 > 0 {
		num %= 4
		ans = ans + "IV"
	}

	ans = ans + strings.Repeat("I", num)

	return ans

}
