package graphGeneral

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	islands := 0

	var dfs func(int, int)
	dfs = func(row int, col int) {
		if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] != '1' {
			return
		}
		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				islands++
			}
		}
	}
	return islands
}
