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

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += 1
	}
	termbox.Flush()
}

func drawSquare(x int, y int, size int) {
	for i := 0; i < size; i++ {
		termbox.SetCell(x+i, y, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x+i, y+size-1, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x+size-1, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
	}
}

func drawEx(x int, y int, size int) {
	for i := 0; i < size; i++ {
		termbox.SetCell(x+i, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x+size-1-i, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
	}
}

func drawMark(positionX int, positionY int, mark Cell, boardSize int) {
	w, h := termbox.Size()
	centerH := h / 2
	centerW := w / 2
	squareSize := 5
	lengthSize := boardSize*squareSize + (boardSize - 1)
	topLeftCornerX := centerW - lengthSize/2
	topLeftCornerY := centerH - lengthSize/2
	xPos := topLeftCornerX + (squareSize+1)*positionX
	yPos := topLeftCornerY + (squareSize+1)*positionY

	if mark == Ex {
		drawEx(xPos+1, yPos+1, squareSize-2)
	} else if mark == Circle {
		drawSquare(xPos+1, yPos+1, squareSize-2)
	}
	termbox.Flush()
}

func drawBoard() {
	w, h := termbox.Size()
	centerH := h / 2
	centerW := w / 2
	boardSize := 3
	squareSize := 5
	lengthSize := boardSize*squareSize + (boardSize - 1)
	nLines := boardSize - 1
	for i := 1; i <= nLines; i++ {
		start := centerH - lengthSize/2
		end := start + lengthSize
		xMaxLeftSize := centerW - lengthSize/2 - 1
		xPos := xMaxLeftSize + (squareSize+1)*i
		for yPos := start; yPos < end; yPos++ {
			termbox.SetCell(xPos, yPos, 1, termbox.ColorCyan, termbox.ColorCyan)
		}
	}

	for i := 1; i <= nLines; i++ {
		start := centerW - lengthSize/2
		end := start + lengthSize
		yMaxLeftSize := centerH - lengthSize/2 - 1
		yPos := yMaxLeftSize + (squareSize+1)*i
		for xPos := start; xPos < end; xPos++ {
			termbox.SetCell(xPos, yPos, 1, termbox.ColorCyan, termbox.ColorCyan)
		}
	}
	termbox.Flush()
}

func getPlayFromPixels(mx int, my int) (int, int) {
	w, h := termbox.Size()
	centerH := h / 2
	centerW := w / 2
	boardSize := 3
	squareSize := 5
	lengthSize := boardSize*squareSize + (boardSize - 1)
	topLeftCornerX := centerW - lengthSize/2
	topLeftCornerY := centerH - lengthSize/2
	xPos, yPos := -1, -1
	if mx >= topLeftCornerX && mx <= topLeftCornerX+lengthSize {
		xPos = (mx - topLeftCornerX) / (squareSize + 1)
	}
	if my >= topLeftCornerY && my <= topLeftCornerY+lengthSize {
		yPos = (my - topLeftCornerY) / (squareSize + 1)
	}

	if my > 0 && mx > 0 {
		return xPos, yPos
	}
	return -1, -1
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
	board := initBoard(boardSize)
	currentMark := Ex
	drawBoard()
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
			tbprint(centerW-len(message)/2, centerH, termbox.ColorWhite, termbox.ColorBlack, message)
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
			xPos, yPos := getPlayFromPixels(mx, my)
			if xPos > -1 && yPos > -1 {
				board[xPos][yPos] = currentMark
				drawMark(xPos, yPos, currentMark, len(board))

				if isGameFinished(board) {
					takingPlays = false
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					tbprint(0, 0, termbox.ColorWhite, termbox.ColorBlack, "Game finished")
					termbox.Flush()
					continue mainloop
				}

				cpuXPos, cpuYPos := takeCpuTurn(board, Circle)
				drawMark(cpuXPos, cpuYPos, Circle, len(board))

				if isGameFinished(board) {
					takingPlays = false
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					tbprint(0, 0, termbox.ColorWhite, termbox.ColorBlack, "Game finished")
					termbox.Flush()
				}
			}
		}
	}

}
