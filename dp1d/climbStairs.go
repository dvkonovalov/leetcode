package dp1d

var mp = make(map[int]int)

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	var nm1, nm2 int
	var ok bool
	if nm1, ok = mp[n-1]; !ok {
		nm1 = climbStairs(n - 1)
		mp[n-1] = nm1
	}
	if nm2, ok = mp[n-2]; !ok {
		nm2 = climbStairs(n - 2)
		mp[n-2] = nm2
	}
	return nm1 + nm2
}
