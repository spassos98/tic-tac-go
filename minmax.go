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

func findMinWithIndex(arr []int, operation int) (int, int) {
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

func minmax(board Board, depth int) int {
	var scores []int
	var moves []move
	possibleMoves := getPossibleMoves(board)

	isFirstPlayer := depth%2 == 0
	var mark Cell
	if isFirstPlayer {
		mark = Ex
	} else {
		mark = Circle
	}

	newBoard := duplicateBoard(board)
	for _, currentMove := range possibleMoves {
		newBoard[currentMove.x][currentMove.y] = mark
		score := minmax(newBoard, depth+1)
		scores = append(scores, score)
		moves = append(moves, currentMove)
		newBoard[currentMove.x][currentMove.y] = Empty
	}

	var bestValue, bestMoveIdx int
	if isFirstPlayer {
		bestValue, bestMoveIdx = findMaxWithIndex(scores)
	} else {
		bestValue, bestMoveIdx = findMaxWithIndex(scores)
	}

	fmt.Printf("Best move is %d %d\n", moves[bestMoveIdx].x, moves[bestMoveIdx].y)
	return bestValue
}
