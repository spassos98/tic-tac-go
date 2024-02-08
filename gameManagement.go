package main

func isRowComplete(board Board, mark Cell) bool {
	for i := 0; i < len(board); i++ {
		base := board[i][0]
		if base != mark {
			continue
		}
		rowComplete := true
		for j := 1; j < len(board[i]); j++ {
			if board[i][j] != base {
				rowComplete = false
				break
			}
		}
		if rowComplete {
			return true
		}
	}
	return false
}

func isColComplete(board Board, mark Cell) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < cols; i++ {
		base := board[0][i]
		if base != mark {
			continue
		}
		colComplete := true
		for j := 1; j < rows; j++ {
			if board[j][i] != base {
				colComplete = false
				break
			}
		}
		if colComplete {
			return true
		}
	}
	return false
}

func isDiagcomplete(board Board, mark Cell) bool {
	n := len(board)
	base1, base2 := board[0][0], board[0][n-1]
	diag1Complete, diag2Complete := base1 == mark, base2 == mark
	for i := 1; i < len(board); i++ {
		if base1 != board[i][i] {
			diag1Complete = false
		}

		if base2 != board[i][n-1-i] {
			diag2Complete = false
		}

	}
	return diag1Complete || diag2Complete
}

func isGameOver(board Board) int {

	return 0
}

func getGameState(board Board, mark Cell) string {
	state := ""
	if isRowComplete(board, mark) {
		state = "Won with rows"
	}
	if isColComplete(board, mark) {
		state = "Won with cols"
	}

	if isDiagcomplete(board, mark) {
		state = "Won with diag"
	}
	return state
}

func changeMark(mark Cell) Cell {
	if mark == Ex {
		return Circle
	}
	return Ex
}