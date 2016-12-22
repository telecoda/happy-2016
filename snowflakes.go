package main

import (
	"math/rand"

	termbox "github.com/nsf/termbox-go"
)

type Snowflake struct {
	flake     rune
	x, y      float32
	xVelocity float32
	yVelocity float32
}

func NewSnowflake() *Snowflake {

	flake := &Snowflake{
		flake: rune('⎈'),
	}

	flake.reset()

	return flake
}

func (s *Snowflake) updatePosition() {
	uiWidth, uiHeight := termbox.Size()
	// update position based upon velocity
	s.x += s.xVelocity
	s.y += s.yVelocity

	if int(s.x) < 0 || int(s.x) > uiWidth {
		s.reset()
	}

	if int(s.y) < 0 || int(s.y) > uiHeight {
		s.reset()
	}

}

func (s *Snowflake) reset() {
	uiWidth, _ := termbox.Size()
	s.y = 0
	s.x = float32(rand.Intn(uiWidth))
	s.xVelocity = rand.Float32() - 0.5
	s.yVelocity = rand.Float32()
}

func (s *Snowflake) render() {
	termbox.SetCell(int(s.x), int(s.y), s.flake, termbox.ColorWhite, termbox.ColorBlack)
}
