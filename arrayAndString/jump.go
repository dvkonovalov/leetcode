package arrayAndString

func jump(nums []int) int {
	if len(nums) < 3 {
		return len(nums) - 1
	}

	count := 1
	stayStep := nums[0] - 1
	maxPos := nums[0]
	if len(nums) <= maxPos+1 {
		return count
	}
	for i := 1; i < len(nums); i++ {
		if maxPos < nums[i]+i {
			maxPos = nums[i] + i
		}
		if len(nums) <= maxPos+1 {
			return count + 1
		}
		if stayStep == 0 {
			stayStep = maxPos - i
			count++
		}
		stayStep--
	}
	return count + 1
}
