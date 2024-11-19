package arrayAndString

func longestCommonPrefix(strs []string) string {
	answer := ""
	for i := 0; i < len(strs[0]); i++ {
		answer += string(strs[0][i])
		for _, str := range strs[1:] {
			if len(answer) > len(str) {
				return str
			}
			if str[:i+1] != answer {
				return answer[:len(answer)-1]
			}
		}
	}
	return answer
}
