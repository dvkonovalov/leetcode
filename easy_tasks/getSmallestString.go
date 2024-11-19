package easy_tasks

import "strconv"

func getSmallestString(s string) string {
	for i := 0; i < len(s)-1; i++ {
		tmp, _ := strconv.Atoi(string(s[i]))
		next, _ := strconv.Atoi(string(s[i+1]))
		if tmp%2 == next%2 && tmp > next {
			s = s[:i] + string(s[i+1]) + string(s[i]) + string(s[i+2:])
			return s
		}
	}
	return s
}
