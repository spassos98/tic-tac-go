package main

import "fmt"

type GameState int

const (
	Win     GameState = iota
	Draw    GameState = iota
	Running GameState = iota
)

func initBoard(boardSize int) Board {
	var board Board
	for i := 0; i < boardSize; i++ {
		row := make([]Cell, boardSize)
		board = append(board, row)
	}
	return board
}

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

func boardIsFull(board Board) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

func getGameStatePlayer(board Board, mark Cell) GameState {
	if isRowComplete(board, mark) || isColComplete(board, mark) || isDiagcomplete(board, mark) {
		return Win
	}
	if boardIsFull(board) {
		return Draw
	}
	return Running
}

func getGameState(board Board) (GameState, Cell) {
	gameStateP1 := getGameStatePlayer(board, Ex)
	gameStateP2 := getGameStatePlayer(board, Circle)

	// If it's draw it's draw for both players
	if gameStateP1 == Draw {
		return gameStateP1, Empty
	}

	if gameStateP1 == Win {
		return gameStateP1, Ex
	}

	if gameStateP2 == Win {
		return gameStateP2, Circle
	}

	return gameStateP1, Empty
}

func checkGameState(board Board) GameState {
	gameState, playerState := getGameState(board)
	if gameState == Win {
		if playerState == Ex {
			fmt.Println("Player 1 Won!")
		} else if playerState == Circle {
			fmt.Println("Player 2 Won!")
		}
		return Win
	}
	if gameState == Draw {
		fmt.Println("It's a Draw")
		return Draw
	}
	return Running
}

func isGameFinished(board Board) bool {
	gameState := checkGameState(board)
	if gameState == Win || gameState == Draw {
		return true
	}
	return false
}
