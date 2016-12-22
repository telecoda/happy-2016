package main

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var drawing *Drawing
var snowflakes []*Snowflake

func main() {
	fmt.Printf("Snow sim\n")
	Init()
}

func Init() {
	termbox.Init()
	// init termbox settings
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.SetOutputMode(termbox.OutputNormal)
	//termbox.SetOutputMode(termbox.OutputGrayscale)
	termWidth, termHeight := termbox.Size()
	// drawing is 1 line less than terminal to allow for status bar
	drawing = NewDrawing(termWidth, termHeight)

	drawGround()

	// draw a snowman
	smBottom := termHeight - 2
	smLeft := termWidth - 30

	drawSnowman(smLeft, smBottom)

	// draw a gopher
	gBottom := termHeight - 2
	gLeft := 10

	drawGopher(gLeft, gBottom)

	// draw a tree
	tBottom := termHeight - 2
	tLeft := termWidth/2 - 10

	drawTree(tLeft, tBottom)

	drawMessage()

	snowflakes = make([]*Snowflake, 0)

	addSnowflakes(150)

	eventLoop()

}

func drawMessage() {
	//termbox.SetOutputMode(termbox.OutputGrayscale)
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2

	// prepare image
	// greetings message
	msg1 := "Season Greetings"
	msg2 := "to all the Gophers"
	msg3 := "from @telecoda"

	drawing.SetText(cx-len(msg1)/2-20, cy-10, msg1, termbox.ColorYellow, termbox.ColorRed)
	drawing.SetText(cx-len(msg2)/2+20, cy-10, msg2, termbox.ColorRed, termbox.ColorYellow)
	drawing.SetText(cx-len(msg3)/2, cy-8, msg3, termbox.ColorRed, termbox.ColorGreen)
}

func drawGround() {
	termWidth, termHeight := termbox.Size()
	// draw ground
	ground := ""
	for i := 0; i < termWidth; i++ {
		ground += "-"
	}
	drawing.SetText(0, termHeight-1, ground, termbox.ColorBlack, termbox.ColorWhite)

}

func drawTree(x, y int) {
	// star
	drawing.SetText(x+10, y-15, "⎈", termbox.ColorYellow, termbox.ColorBlack)
	// tree
	drawing.SetText(x+10, y-14, "  ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+10, y-13, "  ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+8, y-12, "      ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+8, y-11, "      ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+6, y-10, "          ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+6, y-9, "          ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+4, y-8, "              ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+4, y-7, "              ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+2, y-6, "                  ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x+2, y-5, "                  ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x, y-4, "                      ", termbox.ColorGreen, termbox.ColorGreen)
	drawing.SetText(x, y-3, "                      ", termbox.ColorGreen, termbox.ColorGreen)
	// trunk
	drawing.SetText(x+8, y-2, "██████", termbox.ColorMagenta, termbox.ColorRed)

	// bucket
	drawing.SetText(x+6, y-1, "          ", termbox.ColorRed, termbox.ColorRed)
	drawing.SetText(x+6, y, "          ", termbox.ColorRed, termbox.ColorRed)

	// baubles
	drawing.SetText(x+10, y-13, "⓿", termbox.ColorRed, termbox.ColorGreen)
	drawing.SetText(x+8, y-11, "⓿", termbox.ColorYellow, termbox.ColorGreen)
	drawing.SetText(x+12, y-11, "⓿", termbox.ColorCyan, termbox.ColorGreen)
	drawing.SetText(x+6, y-9, "⓿", termbox.ColorMagenta, termbox.ColorGreen)
	drawing.SetText(x+10, y-10, "⓿", termbox.ColorBlue, termbox.ColorGreen)
	drawing.SetText(x+14, y-9, "⓿", termbox.ColorWhite, termbox.ColorGreen)
	drawing.SetText(x+8, y-8, "⓿", termbox.ColorRed, termbox.ColorGreen)
	drawing.SetText(x+12, y-8, "⓿", termbox.ColorYellow, termbox.ColorGreen)
	drawing.SetText(x+4, y-7, "⓿", termbox.ColorCyan, termbox.ColorGreen)
	drawing.SetText(x+10, y-7, "⓿", termbox.ColorMagenta, termbox.ColorGreen)
	drawing.SetText(x+16, y-7, "⓿", termbox.ColorBlue, termbox.ColorGreen)
	drawing.SetText(x+6, y-6, "⓿", termbox.ColorCyan, termbox.ColorGreen)
	drawing.SetText(x+14, y-6, "⓿", termbox.ColorRed, termbox.ColorGreen)
	drawing.SetText(x+2, y-5, "⓿", termbox.ColorYellow, termbox.ColorGreen)
	drawing.SetText(x+10, y-5, "⓿", termbox.ColorMagenta, termbox.ColorGreen)
	drawing.SetText(x+18, y-5, "⓿", termbox.ColorRed, termbox.ColorGreen)
	drawing.SetText(x+4, y-4, "⓿", termbox.ColorBlue, termbox.ColorGreen)
	drawing.SetText(x+16, y-4, "⓿", termbox.ColorWhite, termbox.ColorGreen)
	drawing.SetText(x, y-3, "⓿", termbox.ColorRed, termbox.ColorGreen)
	drawing.SetText(x+10, y-3, "⓿", termbox.ColorYellow, termbox.ColorGreen)
	drawing.SetText(x+20, y-3, "⓿", termbox.ColorCyan, termbox.ColorGreen)

}

func drawSnowman(x, y int) {
	// hat
	drawing.SetText(x+1, y-14, "       ", termbox.ColorBlack, termbox.ColorBlue)
	drawing.SetText(x+1, y-13, "  ███  ", termbox.ColorBlue, termbox.ColorYellow)
	drawing.SetText(x-1, y-12, "           ", termbox.ColorBlack, termbox.ColorBlue)
	// head
	drawing.SetText(x+1, y-11, "       ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x, y-10, "         ", termbox.ColorBlack, termbox.ColorWhite)
	// eye
	drawing.SetText(x+2, y-10, "⓿", termbox.ColorBlue, termbox.ColorWhite)
	drawing.SetText(x+5, y-10, "⓿", termbox.ColorBlue, termbox.ColorWhite)
	drawing.SetText(x, y-9, " *     * ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x, y-8, "  *****  ", termbox.ColorBlack, termbox.ColorWhite)
	// nose
	drawing.SetText(x+4, y-9, "▓", termbox.ColorRed, termbox.ColorYellow)
	drawing.SetText(x+1, y-7, "       ", termbox.ColorBlack, termbox.ColorWhite)
	// body
	drawing.SetText(x-1, y-6, "           ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+4, y-6, "⓿", termbox.ColorRed, termbox.ColorWhite)
	drawing.SetText(x-1, y-5, "           ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x-1, y-4, "           ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+4, y-4, "⓿", termbox.ColorRed, termbox.ColorWhite)
	drawing.SetText(x-1, y-3, "           ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x-1, y-2, "           ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+4, y-2, "⓿", termbox.ColorRed, termbox.ColorWhite)
	drawing.SetText(x-1, y-1, "           ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x, y, "         ", termbox.ColorBlack, termbox.ColorWhite)
}

func drawGopher(x, y int) {
	// ears
	drawing.SetText(x+2, y-20, "  ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+24, y-20, "  ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x, y-19, "      ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-19, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+22, y-19, "      ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+24, y-19, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	// head
	drawing.SetText(x+8, y-19, "            ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+10, y-18, "        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+12, y-17, "    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+12, y-16, "    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+12, y-15, "    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+10, y-14, "        ", termbox.ColorBlack, termbox.ColorCyan)
	// eyes
	drawing.SetText(x+4, y-18, "      ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+18, y-18, "      ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+2, y-17, "          ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+16, y-17, "          ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+2, y-16, "          ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+6, y-16, "  ", termbox.ColorBlack, termbox.ColorBlack)
	drawing.SetText(x+16, y-16, "          ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+20, y-16, "  ", termbox.ColorBlack, termbox.ColorBlack)
	drawing.SetText(x+2, y-15, "          ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+16, y-15, "          ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+4, y-14, "      ", termbox.ColorBlack, termbox.ColorWhite)
	drawing.SetText(x+18, y-14, "      ", termbox.ColorBlack, termbox.ColorWhite)
	// body
	drawing.SetText(x+4, y-13, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+4, y-12, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+4, y-11, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+4, y-10, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+4, y-9, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+4, y-8, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+4, y-7, "                    ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-6, "                        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-5, "                        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-4, "                        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-3, "                        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-2, "                        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+2, y-1, "                        ", termbox.ColorBlack, termbox.ColorCyan)
	drawing.SetText(x+5, y, "                  ", termbox.ColorBlack, termbox.ColorCyan)
	// nose
	drawing.SetText(x+12, y-14, "    ", termbox.ColorBlack, termbox.ColorBlack)
	drawing.SetText(x+10, y-13, "▓▓▓▓▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+10, y-12, "▓▓▓▓▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	// teeth
	drawing.SetText(x+12, y-11, "    ", termbox.ColorWhite, termbox.ColorWhite)
	drawing.SetText(x+12, y-10, "    ", termbox.ColorWhite, termbox.ColorWhite)
	// arms
	drawing.SetText(x+2, y-11, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+24, y-11, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x, y-10, "▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+24, y-10, "▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x, y-9, "▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+24, y-9, "▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x, y-8, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+26, y-8, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	// feet
	drawing.SetText(x+4, y, "▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+20, y, "▓▓▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+6, y+1, "▓▓", termbox.ColorWhite, termbox.ColorYellow)
	drawing.SetText(x+20, y+1, "▓▓", termbox.ColorWhite, termbox.ColorYellow)

}

func addSnowflakes(number int) {
	for i := 0; i < number; i++ {
		snowflake := NewSnowflake()
		snowflakes = append(snowflakes, snowflake)
	}
}

func eventLoop() {

	go renderLoop()
mainloop:
	for {

		// global event handling first
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			//lastKeyEvent = ev
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
		}
	}

	termbox.Flush()

}

func renderLoop() {
	for {
		// update position of flakes
		updateFlakes()
		drawing.render()
		renderFlakes()
		termbox.Flush()
		time.Sleep(200 * time.Millisecond)
	}
}

func updateFlakes() {
	// move flakes based on velocity
	for i, _ := range snowflakes {
		snowflake := snowflakes[i]
		snowflake.updatePosition()
		// if cell that snowflake overlaps is occupied
		// increase level of snow
		cell := drawing.GetCell(int(snowflake.x), int(snowflake.y))

		if cell != nil {

			if drawing.containsSnow(cell) {
				// increase snow level in current cell
				drawing.SnowLands(int(snowflake.x), int(snowflake.y), cell)
				// reset flake
				snowflake.reset()
				return
			}
			// check if cell below is full
			cellBelow := drawing.GetCell(int(snowflake.x), int(snowflake.y+1))
			if cellBelow != nil {
				if drawing.containsObstacle(cellBelow) {
					// increase snow level in current cell
					drawing.SnowLands(int(snowflake.x), int(snowflake.y), cell)
					// reset flake
					snowflake.reset()
					return
				}
			}
		}
	}
}

func renderFlakes() {
	for i, _ := range snowflakes {
		snowflakes[i].render()
	}
}
