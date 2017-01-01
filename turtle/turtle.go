package turtle

import (
	"image"
	"image/color"
)

var drawColor = color.RGBA{0, 0, 0, 255}

const (
	Up = iota
	Right
	Down
	Left
	NoDir
)

const (
	NoCommand = iota
	GoUp
	GoLeft
	TurnR
	TurnL
	Stop
)

type Turtle struct {
	colorMap map[color.Color]int
	secret   *image.RGBA
	code     image.Image
}

func Decrypter(UpColor, LeftColor, StopColor, TurnRColor, TurnLColor color.Color) *Turtle {
	// If color doesn't match, map will return NoCommand as zero value
	colorMap := map[color.Color]int{
		UpColor:    GoUp,
		LeftColor:  GoLeft,
		StopColor:  Stop,
		TurnRColor: TurnR,
		TurnLColor: TurnL,
	}
	return &Turtle{colorMap: colorMap}
}

func (t *Turtle) Decrypt(m image.Image) image.Image {
	t.code = m
	bounds := m.Bounds()
	t.secret = image.NewRGBA(bounds)

	for y := bounds.Max.Y - 1; y >= 0; y-- {
		for x := bounds.Max.X - 1; x >= 0; x-- {
			command := t.colorMap[m.At(x, y)]
			if command == GoUp || command == GoLeft {
				t.DrawNextLine(x, y, command, NoDir)
			}
		}
	}

	return t.secret
}

func (t *Turtle) DrawNextLine(x, y, command, direction int) {
	switch command {
	case GoUp:
		t.drawLine(x, y, Up)
	case GoLeft:
		t.drawLine(x, y, Left)
	case TurnR:
		t.drawLine(x, y, (direction+1)%4)
	case TurnL:
		t.drawLine(x, y, (direction+3)%4)
	}
}

func (t *Turtle) drawLine(x, y, direction int) {
	curX := x
	curY := y
	command := NoCommand
	for command == NoCommand {
		switch direction {
		case Up:
			curY--
		case Down:
			curY++
		case Left:
			curX--
		case Right:
			curX++
		}
		command = t.colorMap[t.code.At(curX, curY)]
	}

	if direction == Left || direction == Up {
		drawRect(t.secret, image.Pt(curX, curY), image.Pt(x, y))
	} else {
		drawRect(t.secret, image.Pt(x, y), image.Pt(curX, curY))
	}

	if command != Stop {
		t.DrawNextLine(curX, curY, command, direction)
	}
}

// Fill in rectangle - though only used to draw lines
func drawRect(dst *image.RGBA, p1, p2 image.Point) {
	for x := p1.X; x <= p2.X; x++ {
		for y := p1.Y; y <= p2.Y; y++ {
			dst.Set(x, y, drawColor)
		}
	}
}
