package easy_tasks

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	if numRows < 1 {
		return ans
	}
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	ans = append(ans, []int{1, 1})
	if numRows == 2 {
		return ans
	}
	for i := 2; i < numRows; i++ {
		newSlice := make([]int, 0)
		newSlice = append(newSlice, 1)
		for j := 1; j < i; j++ {
			newSlice = append(newSlice, ans[i-1][j]+ans[i-1][j-1])
		}
		newSlice = append(newSlice, 1)
		ans = append(ans, newSlice)
	}
	return ans
}
