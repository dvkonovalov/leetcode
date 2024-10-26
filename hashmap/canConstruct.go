package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[rune]int)
	for _, symbol := range magazine {
		mp[symbol]++
	}
	for _, symbol := range ransomNote {
		if val, ok := mp[symbol]; !ok || val == 0 {
			return false
		}
		mp[symbol]--

	}
	return true
}
