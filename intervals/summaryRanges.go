package intervals

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	tmp := 0
	answer := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			if nums[i-1] != nums[tmp] {
				answer = append(answer, fmt.Sprintf("%d->%d", nums[tmp], nums[i-1]))
			} else {
				answer = append(answer, strconv.Itoa(nums[i-1]))
			}
			tmp = i
		}
	}

	if nums[len(nums)-1] != nums[tmp] {
		answer = append(answer, fmt.Sprintf("%d->%d", nums[tmp], nums[len(nums)-1]))
	} else {
		answer = append(answer, strconv.Itoa(nums[len(nums)-1]))
	}
	return answer
}
