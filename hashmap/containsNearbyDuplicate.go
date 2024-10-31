package hashmap

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i, num := range nums {
		pos, found := mp[num]
		if !found {
			mp[num] = i
			continue
		}
		if i-pos <= k {
			return true
		}
		mp[num] = i
	}
	return false
}
