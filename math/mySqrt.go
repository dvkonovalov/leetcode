package math

func mySqrt(x int) int {
	left := 0
	right := x

	for left <= right {
		mid := (right + left) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return right

}
