package arrayAndString

func removeElement(nums []int, val int) int {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[pos] = nums[i]
		pos++
	}
	return pos
}
