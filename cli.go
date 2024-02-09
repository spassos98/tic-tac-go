package main

import (
	"errors"
	"fmt"
	"strings"
)

type Cell int
type Board [][]Cell

const (
	Empty  Cell = iota
	Ex     Cell = iota
	Circle Cell = iota
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

func main() {
	fmt.Println("Welcome to tic-tac-go!")
	fmt.Println("Please enter the size of the board (an integer)")
	var boardSize int
	fmt.Scanf("%d", &boardSize)
	fmt.Printf("You selected a size of %d\n", boardSize)
	var board Board
	for i := 0; i < boardSize; i++ {
		row := make([]Cell, boardSize)
		board = append(board, row)
	}
	printBoard(board)
	play := -1
	currentMark := Ex

	for {
		// User turn
		_, err := fmt.Scanf("%d", &play)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if play < 0 || play > (boardSize*boardSize)-1 {
			fmt.Println("Please enter a number between 0 and 8")
			continue
		}

		fmt.Println("")
		positionX := play / boardSize
		positionY := play % boardSize

		if board[positionX][positionY] != Empty {
			fmt.Println("That position is already occupied, try again")
			continue
		}
		board[positionX][positionY] = currentMark
		gameState, playerState := getGameState(board)
		if gameState == Win {
			printBoard(board)
			if playerState == Ex {
				fmt.Println("Player 1 Won!")
			} else if playerState == Circle {
				fmt.Println("Player 2 Won!")
			}
			break
		}
		if gameState == Draw {
			fmt.Println("It's a Draw")
			break
		}

		// CPU turn
		_, bestMove := minmax(board, 1)
		board[bestMove.x][bestMove.y] = Circle
		printBoard(board)
		gameState, playerState = getGameState(board)
		if gameState == Win {
			if playerState == Ex {
				fmt.Println("Player 1 Won!")
			} else if playerState == Circle {
				fmt.Println("Player 2 Won!")
			}
			break
		}
		if gameState == Draw {
			fmt.Println("It's a Draw")
		}
	}
}
