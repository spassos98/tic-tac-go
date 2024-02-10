package main

import "fmt"

func readBoardSize() int {
	fmt.Println("Please enter the size of the board (an integer)")
	var boardSize int
	fmt.Scanf("%d", &boardSize)
	fmt.Printf("You selected a size of %d\n", boardSize)
	return boardSize
}

func readGameMode() GameMode {
	fmt.Println("Please select a game mode")
	fmt.Println("1. Player vs CPU")
	fmt.Println("2. Player vs Player")

	var option int
	var gameMode GameMode
	fmt.Scanf("%d", &option)
	if option == 1 {
		gameMode = PlayerVsCPU
	} else if option == 2 {
		gameMode = PlayerVsPlayer
	}
	return gameMode
}
