package main

import (
	"errors"
	"fmt"
	"strings"
)

type Cell int
type Board [][]Cell

const (
	Empty Cell = iota
	Ex
	Circle
)

func getCellRepresentation(cellValue Cell) (string, error) {
	switch cellValue {
	case Empty:
		return "_", nil
	case Ex:
		return "X", nil
	case Circle:
		return "O", nil
	}
	return "", errors.New("No representation found for given cell value")
}

func printBoard(board Board) error {
	for i := 0; i < len(board); i++ {
		boardRow := make([]string, len(board))
		var err error
		for j := 0; j < len(board[i]); j++ {
			boardRow[j], err = getCellRepresentation(board[i][j])
			if err != nil {
				return errors.New("Could not reprenst cell value")
			}
		}
		fmt.Printf("%s\n", strings.Join(boardRow, " "))
	}
	return nil
}

func isRowComplete(board Board) bool {
	for i := 0; i < len(board); i++ {
		base := board[i][0]
		if base == Empty {
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

func isColComplete(board Board) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < cols; i++ {
		base := board[0][i]
		if base == Empty {
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

func isDiagcomplete(board Board) bool {
	n := len(board)
	base1, base2 := board[0][0], board[0][n-1]
	diag1Complete, diag2Complete := base1 != Empty, base2 != Empty
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

func getGameState(board Board) string {
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

func changeMark(mark Cell) Cell {
	if mark == Ex {
		return Circle
	}
	return Ex
}

func main() {
	board := Board{
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
	}
	printBoard(board)
	play := -1
	plays := 0
	currentMark := Ex

	for plays < 9 {
		_, err := fmt.Scanf("%d", &play)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if play < 0 || play > 8 {
			fmt.Println("Please enter a number between 0 and 8")
			continue
		}

		fmt.Println("")
		positionX := play / 3
		positionY := play % 3

		if board[positionX][positionY] != Empty {
			fmt.Println("That position is already occupied, try again")
			continue
		}
		board[positionX][positionY] = currentMark
		currentMark = changeMark(currentMark)
		plays += 1
		boardErr := printBoard(board)
		if boardErr != nil {
			fmt.Println("Could not print board, stopping game...")
		}
		gameState := getGameState(board)
		if gameState != "" {
			fmt.Println(gameState)
			break
		}
	}
}
