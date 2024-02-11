package main

import "fmt"

type move struct {
	x int
	y int
}

func duplicateBoard(board Board) Board {
	duplicate := make(Board, len(board))
	for i := range board {
		duplicate[i] = make([]Cell, len(board[i]))
		copy(duplicate[i], board[i])
	}
	return duplicate
}

func getPossibleMoves(board Board) []move {
	var moves []move
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == Empty {
				moves = append(moves, move{x: i, y: j})
			}
		}
	}
	return moves
}

func findMinWithIndex(arr []int) (int, int) {
	minValue := 1 << 30
	minIdx := 0
	for idx, val := range arr {
		if val < minValue {
			minValue = val
			minIdx = idx
		}
	}
	return minValue, minIdx
}

func findMaxWithIndex(arr []int) (int, int) {
	maxValue := 0
	maxIdx := 0
	for idx, val := range arr {
		if val > maxValue {
			maxValue = val
			maxIdx = idx
		}
	}
	return maxValue, maxIdx
}

func printPossibleMoves(moves []move) {
	fmt.Println("Possible Moves")
	for _, m := range moves {
		fmt.Printf("\t [%d, %d]\n", m.x, m.y)
	}
}

func minmax(board Board, depth int) (int, move) {
	if len(board) > 3 && depth >= len(board)*2-1 {
		return 0, move{x: -1, y: -1}
	}
	isFirstPlayer := depth%2 == 0
	var mark Cell
	if isFirstPlayer {
		mark = Ex
	} else {
		mark = Circle
	}
	gameState, playerMark := getGameState(board)
	if gameState == Win {
		if playerMark == Ex {
			return 10, move{x: -1, y: -1}
		} else if playerMark == Circle {
			return -10, move{x: -1, y: -1}
		}
	}

	if gameState == Draw {
		return 0, move{x: -1, y: -1}
	}

	var scores []int
	var moves []move
	possibleMoves := getPossibleMoves(board)

	newBoard := duplicateBoard(board)
	for _, currentMove := range possibleMoves {
		newBoard[currentMove.x][currentMove.y] = mark
		score, _ := minmax(newBoard, depth+1)
		scores = append(scores, score)
		moves = append(moves, currentMove)
		newBoard[currentMove.x][currentMove.y] = Empty
		if isFirstPlayer && score > 1 {
			break
		} else if !isFirstPlayer && score < 0 {
			break
		}
	}

	var bestValue, bestMoveIdx int
	if isFirstPlayer {
		bestValue, bestMoveIdx = findMaxWithIndex(scores)
	} else {
		bestValue, bestMoveIdx = findMinWithIndex(scores)
	}

	return bestValue, moves[bestMoveIdx]
}
