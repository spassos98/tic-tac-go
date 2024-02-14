package main

import (
	"errors"
	"fmt"
	"github.com/nsf/termbox-go"
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

var current string

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	boardSize := 3
	squareSize := 5
	canvas := NewBoardCanvas(boardSize, squareSize)
	board := initBoard(boardSize)
	currentMark := Ex
	canvas.drawBoard()
	// drawBoard()
	takingPlays := true
mainloop:
	for {
		mx, my := -1, -1
		if !takingPlays {
			termbox.Clear(termbox.ColorBlack, termbox.ColorDefault)
			gameState, playerMark := getGameState(board)
			var message string
			if gameState == Win {
				if playerMark == Ex {
					message = "Player 1 won"
				} else {
					message = "Player 2 won"
				}
			} else {
				message = "It's a ddddraw"
			}
			w, h := termbox.Size()
			centerW := w / 2
			centerH := h / 2
			canvas.tbprint(centerW-len(message)/2, centerH, termbox.ColorWhite, termbox.ColorBlack, message)
			termbox.Flush()
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
		case termbox.EventMouse:
			if ev.Key == termbox.MouseLeft && takingPlays {
				mx, my = ev.MouseX, ev.MouseY
			}
		}
		if takingPlays {
			xPos, yPos := canvas.getPlayFromPixels(mx, my)
			if xPos > -1 && yPos > -1 {
				board[xPos][yPos] = currentMark
				canvas.drawMark(xPos, yPos, currentMark, len(board))

				if isGameFinished(board) {
					takingPlays = false
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					canvas.tbprint(0, 0, termbox.ColorWhite, termbox.ColorBlack, "Game finished")
					termbox.Flush()
					continue mainloop
				}

				cpuXPos, cpuYPos := takeCpuTurn(board, Circle)
				canvas.drawMark(cpuXPos, cpuYPos, Circle, len(board))

				if isGameFinished(board) {
					takingPlays = false
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					canvas.tbprint(0, 0, termbox.ColorWhite, termbox.ColorBlack, "Game finished")
					termbox.Flush()
				}
			}
		}
	}

}
