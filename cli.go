package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
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

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	boardSize := 3
	squareSize := 7
	canvas := NewBoardCanvas(boardSize, squareSize)
	board := initBoard(boardSize)
	currentMark := Ex
	canvas.drawBoard()
	takingPlays := true
	var start time.Time
	endScreenTime := 1 * time.Second
mainloop:
	for {
		mx, my := -1, -1
		if !takingPlays {
			if time.Since(start) >= endScreenTime {
				canvas.drawGameFinished(board)
				switch ev := termbox.PollEvent(); ev.Type {
				case termbox.EventKey:
					if ev.Key == termbox.KeyEsc {
						break mainloop
					}
				case termbox.EventMouse:
					if ev.Key == termbox.MouseLeft {
						break mainloop
					}
				}
			}
			continue
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
		case termbox.EventResize:
			termbox.Clear(termbox.ColorBlack, termbox.ColorDefault)
			canvas.resizeBoardCanvas()
			canvas.redrawEverything(board)
		}
		xPos, yPos := canvas.getPlayFromPixels(mx, my)
		if xPos > -1 && yPos > -1 {
			if board[xPos][yPos] != Empty {
				canvas.printOccupiedCellMessage()
				continue mainloop
			}
			board[xPos][yPos] = currentMark
			canvas.drawMark(xPos, yPos, currentMark)
			canvas.clearOccupiedCellMessage()

			if isGameFinished(board) {
				takingPlays = false
				start = time.Now()
				continue mainloop
			}

			cpuXPos, cpuYPos := takeCpuTurn(board, Circle)
			canvas.drawMark(cpuXPos, cpuYPos, Circle)

			if isGameFinished(board) {
				takingPlays = false
				start = time.Now()
				continue mainloop
			}
		}
	}

}
