package intervals

import (
	"sort"
)

type SliceOfSlices [][]int

func (s SliceOfSlices) Len() int {
	return len(s)
}

func (s SliceOfSlices) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s SliceOfSlices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Sort(SliceOfSlices(intervals))

	ans = append(ans, intervals[0])

	for _, interval := range intervals {
		if ans[len(ans)-1][1] < interval[0] {
			ans = append(ans, interval)
			continue
		}

		if ans[len(ans)-1][1] <= interval[1] {
			ans[len(ans)-1][1] = interval[1]
		}

	}
	return ans
}
