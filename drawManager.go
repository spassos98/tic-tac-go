package main

import (
	"github.com/nsf/termbox-go"
)

type boardCanvas struct {
	boardSize  int
	squareSize int
	lengthSize int
	height     int
	width      int
	centerH    int
	centerW    int
}

func NewBoardCanvas(boardSize int, squareSize int) boardCanvas {
	w, h := termbox.Size()

	lengthSize := boardSize*squareSize + (boardSize - 1)
	c := boardCanvas{boardSize, squareSize, lengthSize, h, w, h / 2, w / 2}
	return c
}

func (c boardCanvas) tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += 1
	}
	termbox.Flush()
}

func (c boardCanvas) drawSquare(x int, y int, size int) {
	for i := 0; i < size; i++ {
		termbox.SetCell(x+i, y, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x+i, y+size-1, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x+size-1, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
	}
}

func (c boardCanvas) drawEx(x int, y int, size int) {
	for i := 0; i < size; i++ {
		termbox.SetCell(x+i, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
		termbox.SetCell(x+size-1-i, y+i, 1, termbox.ColorCyan, termbox.ColorCyan)
	}
}

func (c boardCanvas) drawMark(positionX int, positionY int, mark Cell, boardSize int) {
	topLeftCornerX := c.centerW - c.lengthSize/2
	topLeftCornerY := c.centerH - c.lengthSize/2
	xPos := topLeftCornerX + (c.squareSize+1)*positionX
	yPos := topLeftCornerY + (c.squareSize+1)*positionY

	if mark == Ex {
		c.drawEx(xPos+1, yPos+1, c.squareSize-2)
	} else if mark == Circle {
		c.drawSquare(xPos+1, yPos+1, c.squareSize-2)
	}
	termbox.Flush()
}

func (c boardCanvas) drawBoard() {
	nLines := c.boardSize - 1
	for i := 1; i <= nLines; i++ {
		start := c.centerH - c.lengthSize/2
		end := start + c.lengthSize
		xMaxLeftSize := c.centerW - c.lengthSize/2 - 1
		xPos := xMaxLeftSize + (c.squareSize+1)*i
		for yPos := start; yPos < end; yPos++ {
			termbox.SetCell(xPos, yPos, 1, termbox.ColorCyan, termbox.ColorCyan)
		}
	}

	for i := 1; i <= nLines; i++ {
		start := c.centerW - c.lengthSize/2
		end := start + c.lengthSize
		yMaxLeftSize := c.centerH - c.lengthSize/2 - 1
		yPos := yMaxLeftSize + (c.squareSize+1)*i
		for xPos := start; xPos < end; xPos++ {
			termbox.SetCell(xPos, yPos, 1, termbox.ColorCyan, termbox.ColorCyan)
		}
	}
	termbox.Flush()
}

func (c boardCanvas) getPlayFromPixels(mx int, my int) (int, int) {
	topLeftCornerX := c.centerW - c.lengthSize/2
	topLeftCornerY := c.centerH - c.lengthSize/2
	xPos, yPos := -1, -1
	if mx >= topLeftCornerX && mx <= topLeftCornerX+c.lengthSize {
		xPos = (mx - topLeftCornerX) / (c.squareSize + 1)
	}
	if my >= topLeftCornerY && my <= topLeftCornerY+c.lengthSize {
		yPos = (my - topLeftCornerY) / (c.squareSize + 1)
	}

	if my > 0 && mx > 0 {
		return xPos, yPos
	}
	return -1, -1
}

func (c boardCanvas) drawGameFinished(board Board) {
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
	c.tbprint(c.centerW-len(message)/2, c.centerH, termbox.ColorWhite, termbox.ColorBlack, message)
	termbox.Flush()
}
