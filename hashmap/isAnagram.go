package hashmap

func isAnagram(s string, t string) bool {
	mp := make(map[int32]int)
	val := 0
	var ok bool

	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		mp[c]++
	}
	for _, c := range t {
		if val, ok = mp[c]; !ok || val == 0 {
			return false
		}
		mp[c]--
	}
	return true
}
