package hashmap

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mainMap := make(map[byte]byte)
	sMap := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		val, ok := mainMap[s[i]]
		if !ok {
			if _, ok = sMap[t[i]]; ok {
				return false
			}
			mainMap[s[i]] = t[i]
			sMap[t[i]] = true
			continue
		}
		if val != t[i] {
			return false
		}
	}
	return true
}
