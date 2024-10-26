package binarySearch

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		mid = (left + right)
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
