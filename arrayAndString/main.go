package arrayAndString

func RunTask() {

	// Task - Merge sorted array
	// https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
	nums1 := []int{1, 2, 3, 5, 5, 8, 78, 898989}
	nums2 := []int{1, 2, 3, 4, 5, 1234}
	merge(nums1, len(nums1), nums2, len(nums2))

	// Task - Remove element
	// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(nums, 2)

	// Task - Remove duplicated from sorted array
	// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)

	// Task - Remove duplicated from sorted array II
	// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicatesII(nums)

	// Task - Majority element
	// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
	nums = []int{1, 1, 1, 0, 0, 2, 3, 1}
	majorityElement(nums)

	// Task - Rotate element
	// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)

	// Task - Best Time to Buy and Sell Stock
	// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)

	// Task - Best Time to Buy and Sell Stock II
	// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
	prices = []int{7, 1, 5, 3, 6, 4}
	maxProfitII(prices)
}
