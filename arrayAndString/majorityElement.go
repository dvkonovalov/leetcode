package arrayAndString

func majorityElement(nums []int) int {
	dictNumber := make(map[int]int)
	maximum := 0
	for _, num := range nums {
		if _, ok := dictNumber[num]; ok {
			dictNumber[num]++
		} else {
			dictNumber[num] = 1
		}
		if dictNumber[maximum] < dictNumber[num] {
			maximum = num
		}
	}
	return maximum
}
