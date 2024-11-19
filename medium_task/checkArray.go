package medium_task

func checkArray(nums []int, k int) bool {
	memory := make([]int, k)
	minus := 0
	for i := 0; i < len(nums); i++ {
		minus -= memory[i%k]
		if nums[i]-minus < 0 {
			return false
		}

		if i < len(nums)-k {
			memory[i%k] = nums[i] - minus
			minus = nums[i]
		} else {
			if nums[i]-minus != 0 {
				return false
			}
		}

	}
	return true
}
