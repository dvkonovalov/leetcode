package bitManipulation

func addBinary(a string, b string) string {
	ans := ""
	one := byte(49)
	prevTransferred := false
	posa, posb := len(a)-1, len(b)-1
	for posa != -1 && posb != -1 {

		switch {
		case prevTransferred && a[posa] == one && b[posb] == one:
			prevTransferred = true
			ans = "1" + ans
		case prevTransferred && (a[posa] == one || b[posb] == one):
			prevTransferred = true
			ans = "0" + ans
		case prevTransferred && a[posa] != one && b[posb] != one:
			prevTransferred = false
			ans = "1" + ans
		case a[posa] == one && b[posb] == one:
			prevTransferred = true
			ans = "0" + ans
		case a[posa] == one || b[posb] == one:
			prevTransferred = false
			ans = "1" + ans
		default:
			prevTransferred = false
			ans = "0" + ans
		}
		posa--
		posb--
	}

	if posb != -1 {
		for posb != -1 {
			switch {
			case prevTransferred && b[posb] == one:
				ans = "0" + ans
				prevTransferred = true
			case prevTransferred && b[posb] != one:
				prevTransferred = false
				ans = "1" + ans
			default:
				ans = string(b[posb]) + ans
			}
			posb--

		}
	}
	if posa != -1 {
		for posa != -1 {
			switch {
			case prevTransferred && a[posa] == one:
				ans = "0" + ans
				prevTransferred = true
			case prevTransferred && a[posa] != one:
				prevTransferred = false
				ans = "1" + ans
			default:
				ans = string(a[posa]) + ans
			}
			posa--

		}
	}
	if prevTransferred {
		ans = "1" + ans
	}
	return ans
}
