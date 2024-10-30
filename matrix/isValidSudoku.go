package matrix

func isValidSudoku(board [][]byte) bool {
	var rows [9]map[int]bool
	var cols [9]map[int]bool
	var boxes [9]map[int]bool

	// Инициализируем каждую позицию массива пустой картой
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]bool)
		cols[i] = make(map[int]bool)
		boxes[i] = make(map[int]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			value := int(board[i][j] - '0')

			_, rowsOk := rows[i][value]
			_, colsOk := cols[j][value]
			_, boxesOk := boxes[(i/3)*3+j/3][value]
			if rowsOk || colsOk || boxesOk {
				return false
			}
			rows[i][value] = true
			cols[j][value] = true
			boxes[(i/3)*3+j/3][value] = true
		}
	}
	return true
}
