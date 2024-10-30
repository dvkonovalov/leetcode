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
			// [первый, первый] [второй, второй]
			ans = append(ans, interval)
			continue
		}

		if ans[len(ans)-1][1] >= interval[1] && ans[len(ans)-1][0] <= interval[0] {
			// [первый, первый]
			continue
		}

		if ans[len(ans)-1][0] >= interval[0] && ans[len(ans)-1][1] <= interval[1] {
			// [второй, второй]
			ans[len(ans)-1][0] = interval[0]
			ans[len(ans)-1][1] = interval[1]
			continue
		}

		if ans[len(ans)-1][0] <= interval[0] && ans[len(ans)-1][1] <= interval[1] {
			// [первый, второй]
			ans[len(ans)-1][1] = interval[1]
		}

		if ans[len(ans)-1][0] >= interval[0] && ans[len(ans)-1][1] >= interval[1] {
			// [второй, первый]
			ans[len(ans)-1][0] = interval[0]
		}

	}
	return ans
}
