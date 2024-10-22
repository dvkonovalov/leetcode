package arrayAndString

func merge(nums1 []int, m int, nums2 []int, n int) {
	size := m + n
	firstPos := 0
	secondPos := 0
	res := make([]int, size)
	for i := 0; i < size; i++ {
		if secondPos >= n || firstPos < m && nums1[firstPos] < nums2[secondPos] {
			res[i] = nums1[firstPos]
			firstPos = firstPos + 1
		} else {
			res[i] = nums2[secondPos]
			secondPos = secondPos + 1
		}
	}
	copy(nums1, res)
}
