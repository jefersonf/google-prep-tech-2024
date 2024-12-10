package day1

func checkRow(board [][]byte, i, j int) bool {
	seen := make(map[byte]struct{})
	for columnIndex := 0; columnIndex < len(board); columnIndex++ {
		number := board[i][columnIndex]
		if number == '.' {
			continue
		}
		if _, ok := seen[number]; ok {
			return true
		}
		seen[number] = struct{}{}
	}
	return false
}

func checkColumn(board [][]byte, i, j int) bool {
	seen := make(map[byte]struct{})
	for rowIndex := 0; rowIndex < len(board); rowIndex++ {
		number := board[rowIndex][j]
		if number == '.' {
			continue
		}
		if _, ok := seen[number]; ok {
			return true
		}
		seen[number] = struct{}{}
	}
	return false
}

func checkSubSquare(board [][]byte, i, j int) bool {
	seen := make(map[byte]struct{})
	topRowIndex := int(i/3) * 3
	leftColIndex := int(j/3) * 3
	for i := topRowIndex; i < topRowIndex+3; i++ {
		for j := leftColIndex; j < leftColIndex+3; j++ {
			number := board[i][j]
			if number == '.' {
				continue
			}
			if _, ok := seen[number]; ok {
				return true
			}
			seen[number] = struct{}{}
		}
	}

	return false
}

func isRepeated(board [][]byte, i, j int) bool {
	return checkRow(board, i, j) || checkColumn(board, i, j) || checkSubSquare(board, i, j)
}

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != '.' && isRepeated(board, i, j) {
				return false
			}
		}
	}
	return true
}
