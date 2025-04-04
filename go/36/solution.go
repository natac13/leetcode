package main

const (
	size = 9
)

func isValidSudoku(board [][]byte) bool {

	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := range 9 {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c]
			squareIdx := (r/3)*3 + c/3

			if rows[r][val] || cols[c][val] || squares[squareIdx][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			squares[squareIdx][val] = true
		}
	}

	return true

}
