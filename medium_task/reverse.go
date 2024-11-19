package medium_task

import "math"

func reverse(x int) int {
	ans := int64(0)
	for x != 0 {
		ans = ans*10 + int64(x%10)
		x /= 10
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return int(ans)
}
