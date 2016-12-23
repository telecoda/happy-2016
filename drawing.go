package main

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/y0ssar1an/q"
)

type Drawing struct {
	width, height int
	displayGrid   bool
	drawBuf       []termbox.Cell
}

//var snowFall = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}
var snowFall = []rune{'▁', '▂', '▃', '▅', '▆', '▇'}

func NewDrawing(width, height int) *Drawing {
	drawing := &Drawing{
		width:       width,
		height:      height,
		displayGrid: false,
		drawBuf:     make([]termbox.Cell, width*height),
	}

	return drawing
}

func (d *Drawing) GetCell(x, y int) *termbox.Cell {
	if d.inBounds(x, y) {
		return &d.drawBuf[d.width*y+x]
	}
	return nil
}

func (d *Drawing) setCell(x, y int, char rune, fg, bg termbox.Attribute) {
	if d.inBounds(x, y) {
		d.drawBuf[d.width*y+x] = termbox.Cell{Ch: char, Fg: fg, Bg: bg}
	}

}

func (d *Drawing) SnowLands(x, y int, cell *termbox.Cell) {

	// get current cell
	if d.isEmpty(cell) {

		d.setCell(x, y, snowFall[0], termbox.ColorWhite, termbox.ColorBlack)
		d.setCell(x+1, y, snowFall[0], termbox.ColorWhite, termbox.ColorBlack)
	} else {
		// find current rune.
		for i := 0; i < len(snowFall); i++ {
			if cell.Ch == snowFall[i] {
				// use NEXT rune
				if i == len(snowFall)-1 {
					// cell is full of snow
					cell.Ch = rune(' ')
					cell.Fg = termbox.ColorBlack
					cell.Bg = termbox.ColorWhite
				} else {
					cell.Ch = snowFall[i+1]
				}
				break
			}
		}
		// update next cell
		nextX := (x - x%2) + 1
		nextCell := drawing.GetCell(nextX, y)
		if nextCell != nil {
			// copy to next cell
			nextCell.Fg = cell.Fg
			nextCell.Bg = cell.Bg
			nextCell.Ch = cell.Ch
		}
	}
}

// check if cell contains ANY snow
func (d *Drawing) containsSnow(cell *termbox.Cell) bool {
	for i := 0; i < len(snowFall); i++ {
		if cell.Ch == snowFall[i] {
			return true
		}
	}

	return false
}

// check if cell completely empty
func (d *Drawing) isEmpty(cell *termbox.Cell) bool {
	if cell.Ch == int32(0) {
		return true
	}

	if cell.Ch == rune(' ') && cell.Fg == termbox.ColorWhite && cell.Bg == termbox.ColorBlack {
		return true
	}

	return false
}

// check if cell contains anything else
func (d *Drawing) containsObstacle(cell *termbox.Cell) bool {
	if d.containsSnow(cell) {
		return false
	}
	if d.isEmpty(cell) {
		return false
	}
	return true
}

func (d *Drawing) SetText(x, y int, message string, fg, bg termbox.Attribute) {
	cx := x
	cy := y
	for _, letter := range message {
		if d.inBounds(cx, cy) {
			d.drawBuf[d.width*cy+cx] = termbox.Cell{Ch: rune(letter), Fg: fg, Bg: bg}
			cx++
			q.Q(cx, cy, letter)
		}
	}
}

func (d *Drawing) inBounds(x, y int) bool {
	// check cursor is on drawing bounds
	if x > d.width-1 || y > d.height-1 {
		return false
	}

	if x < 0 || y < 0 {
		return false
	}

	return true
}

func (d *Drawing) render() {
	termbox.Clear(termbox.ColorGreen, termbox.ColorBlue)
	// copy from drawing buffer to on screen buffer
	// dimension may differ..
	uiWidth, uiHeight := termbox.Size()

	for x := 0; x < uiWidth; x++ {
		for y := 0; y < uiHeight; y++ {
			if d.inBounds(x, y) {
				cell := d.GetCell(x, y)
				if cell != nil {
					termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
				}
			}
		}
	}
}
