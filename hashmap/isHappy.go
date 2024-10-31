package hashmap

func isHappy(n int) bool {
	mp := make(map[int]bool)
	for n != 1 {
		if _, ok := mp[n]; ok {
			return false
		}
		mp[n] = true
		newn := 0
		for n > 0 {
			newn += (n % 10) * (n % 10)
			n /= 10
		}
		n = newn
	}
	return true
}
