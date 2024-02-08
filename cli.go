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
	board := Board{
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
	}
	printBoard(board)
	// play := -1
	// plays := 0
	// currentMark := Ex

	// for plays < 9 {
	// 	_, err := fmt.Scanf("%d", &play)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	//
	// 	if play < 0 || play > 8 {
	// 		fmt.Println("Please enter a number between 0 and 8")
	// 		continue
	// 	}
	//
	// 	fmt.Println("")
	// 	positionX := play / 3
	// 	positionY := play % 3
	//
	// 	if board[positionX][positionY] != Empty {
	// 		fmt.Println("That position is already occupied, try again")
	// 		continue
	// 	}
	// 	board[positionX][positionY] = currentMark
	// 	currentMark = changeMark(currentMark)
	// 	plays += 1
	// 	boardErr := printBoard(board)
	// 	if boardErr != nil {
	// 		fmt.Println("Could not print board, stopping game...")
	// 	}
	// 	gameState := getGameState(board)
	// 	if gameState != "" {
	// 		fmt.Println(gameState)
	// 		break
	// 	}
	// }
}
