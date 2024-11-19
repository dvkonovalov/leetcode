package easy_tasks

func backspaceCompare(s string, t string) bool {
	s = convertStr(s)
	t = convertStr(t)
	return s == t
}

func convertStr(s string) string {
	ans := ""
	for _, char := range s {
		if char == '#' && len(ans) == 0 {
			continue
		}
		if char == '#' {
			ans = ans[:len(ans)-1]
		} else {
			ans = ans + string(char)
		}
	}
	return ans
}
