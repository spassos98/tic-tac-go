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
	}
}
