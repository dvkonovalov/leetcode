package hashmap

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		pos, exists := mp[target-v]
		if exists {
			return []int{pos, k}
		}
		mp[v] = k
	}
	return nil
}
