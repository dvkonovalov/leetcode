package graphBFS

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	minMut := 1<<63 - 1
	newArr := make([]string, 0)
	for i := 0; i < len(bank); i++ {
		if diffGenLessOne(startGene, bank[i]) {
			newArr2 := append(newArr, bank[i+1:]...)

			countMut := 1 + minMutation(bank[i], endGene, newArr2)
			if countMut < minMut && countMut > 0 {
				minMut = countMut
			}
		}
		newArr = append(newArr, bank[i])
	}
	if minMut == 1<<63-1 {
		return -1
	}
	return minMut
}

func diffGenLessOne(start string, end string) bool {
	flag := false
	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			if flag {
				return false
			} else {
				flag = true
			}
		}
	}
	return true
}
