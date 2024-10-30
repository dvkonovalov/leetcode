package slidingWindow

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start := 0
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if pos, ok := mp[s[i]]; ok && pos >= start {
			start = pos + 1
		}
		mp[s[i]] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}
