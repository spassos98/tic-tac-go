package main

import (
	"fmt"
	"strings"
)

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func isRowComplete(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		base := board[i][0]
		if base == "_" {
			return false
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

func isColComplete(board [][]string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < cols; i++ {
		base := board[0][i]
		if base == "_" {
			return false
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

func isDiagcomplete(board [][]string) bool {
	n := len(board)
	base1, base2 := board[0][0], board[0][n-1]
	diag1Complete, diag2Complete := base1 != "_", base2 != "_"
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

func getGameState(board [][]string) string {
	state := ""
	if isRowComplete(board) {
		state = "Won with rows"
	}
	if isColComplete(board) {
		state = "Won with cols"
	}

	if isDiagcomplete(board) {
		state = "Won with diag"
	}
	return state
}

func changeMark(mark string) string {
	if mark == "X" {
		return "O"
	}
	return "X"
}

func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	printBoard(board)
	play := -1
	plays := 0
	currentMark := "X"

	for plays < 9 {
		_, err := fmt.Scanf("%d", &play)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("")
		positionX := play / 3
		positionY := play % 3
		board[positionX][positionY] = currentMark
		currentMark = changeMark(currentMark)
		plays += 1
		printBoard(board)
		gameState := getGameState(board)
		if gameState != "" {
			fmt.Println(gameState)
			break
		}
	}
}
