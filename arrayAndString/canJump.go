package arrayAndString

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	count := 0
	curPos := nums[0]
	for curPos < len(nums)-1 {
		if curPos == 0 {
			return false
		}
		if nums[curPos] > count {
			curPos = curPos + nums[curPos]
			count = 0
		} else {
			count++
			curPos--
		}

	}
	return true
}
