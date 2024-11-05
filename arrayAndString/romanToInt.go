package arrayAndString

func romanToInt(s string) int {
	mp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	currentMinNumber := mp[s[0]]
	chislo := 0
	ans := 0
	for _, symbol := range s {
		chislo = mp[byte(symbol)]
		if chislo > currentMinNumber {
			ans -= currentMinNumber * 2
		}
		ans += chislo
	}
	return ans
}
