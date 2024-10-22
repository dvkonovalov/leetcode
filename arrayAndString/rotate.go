package arrayAndString

func rotate(nums []int, k int) {
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
