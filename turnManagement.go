package main

import "fmt"

func takePlayerTurn(board Board, mark Cell) {
	boardSize := len(board)
	var play int

	for {
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
		board[positionX][positionY] = mark
		break
	}
}

func takeCpuTurn(board Board, mark Cell) {
	_, bestMove := minmax(board, 1)
	board[bestMove.x][bestMove.y] = mark
}
