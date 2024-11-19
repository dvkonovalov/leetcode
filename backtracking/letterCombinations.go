package backtracking

var Mp map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans := make([]string, 0)
	number, _ := Mp[string(digits[0])]

	if len(digits) == 1 {
		for _, symbol := range number {
			ans = append(ans, string(symbol))
		}
		return ans
	}
	nextCombination := letterCombinations(digits[1:])
	for _, symbol := range number {
		for _, combination := range nextCombination {
			ans = append(ans, string(symbol)+combination)
		}
	}
	return ans
}
