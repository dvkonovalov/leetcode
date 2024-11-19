package medium_task

func checkArray(nums []int, k int) bool {
	minus := make([]int, k)
	for i := 0; i < len(nums)-k+1; i++ {
		if minus[0] > nums[i] {
			return false
		}
		razn := nums[i] - minus[0]
		for j := 1; j < k; j++ {
			minus[j-1] = minus[j] + razn
		}
		minus[k-1] = 0

	}
	pos := 0
	for i := len(nums) - k + 1; i < len(nums); i++ {
		if minus[pos] != nums[i] {
			return false
		}
		pos++
	}
	return true
}
