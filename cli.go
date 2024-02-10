package main

import (
	"errors"
	"fmt"
	"strings"
)

type Cell int
type GameMode int
type Board [][]Cell

const (
	Empty  Cell = iota
	Ex     Cell = iota
	Circle Cell = iota
)

const (
	PlayerVsCPU    GameMode = iota
	PlayerVsPlayer GameMode = iota
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
	fmt.Println("")
	return nil
}

func initBoard(boardSize int) Board {
	var board Board
	for i := 0; i < boardSize; i++ {
		row := make([]Cell, boardSize)
		board = append(board, row)
	}
	return board
}

func main() {
	fmt.Println("Welcome to tic-tac-go!")
	boardSize := readBoardSize()
	gameMode := readGameMode()
	board := initBoard(boardSize)
	printBoard(board)
	currentMark := Ex

	for {
		takePlayerTurn(board, currentMark)
		if isGameFinished(board) {
			break
		}

		printBoard(board)

		if gameMode == PlayerVsCPU {
			takeCpuTurn(board, Circle)
		} else if gameMode == PlayerVsPlayer {
			takePlayerTurn(board, Circle)
		}

		printBoard(board)

		if isGameFinished(board) {
			break
		}
	}
}
