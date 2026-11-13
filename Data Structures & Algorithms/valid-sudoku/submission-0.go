func isValidSudoku(board [][]byte) bool {
	// check rows for duplicates
	for _, rows := range board {
		rowValues := make(map[byte]bool)
		for _, v := range rows {
			if v == '.' {
				continue
			}
			if rowValues[v] {
				return false
			}
			rowValues[v] = true
		}
	}
	// check columns for duplicates
	i := 0
	colValues := make(map[byte]bool)
	for j := 0; j < len(board); {
		val := board[i][j]
		if val == '.' {
			i++
			if i == 9 {
				j++
				i = 0
				clear(colValues)
			}
			continue
		}
		if colValues[val] {
			return false
		}
		colValues[val] = true
		i++
		if i == 9 {
			j++
			i = 0
			clear(colValues)
		}
	}
	// check 3x3 box for duplicates
	var boxes [9]map[byte]bool
	for i := 0; i < len(board); i++ {
		boxes[i] = make(map[byte]bool)
	}
	for rowIndex, rows := range board {
		for columnIndex, val := range rows {
			if val == '.' {
				continue
			}
			// which square
			squareIndex := (rowIndex / 3) * 3 + (columnIndex / 3)
			if boxes[squareIndex][val] {
				return false
			}
			boxes[squareIndex][val] = true
		}
	}

	return true
}
