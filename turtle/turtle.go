package turtle

import (
	"image"
	"image/color"
	"image/color/palette"
	"time"
)

var drawColor = color.RGBA{0, 255, 0, 255}

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
	colorMap  map[color.Color]int
	secret    *image.Paletted
	encrypted image.Image
	snapshots []*image.Paletted
}

func Decrypter() *Turtle {
	// If color doesn't match, map will return NoCommand as zero value
	colorMap := make(map[color.Color]int)
	snapshots := make([]*image.Paletted, 50)
	return &Turtle{colorMap: colorMap, snapshots: snapshots}
}

func (t *Turtle) SetCommand(command int, clr color.Color) {
	t.colorMap[clr] = command
}

func (t *Turtle) Decrypt(m image.Image) (image.Image, []*image.Paletted) {
	t.encrypted = m
	t.secret = image.NewPaletted(t.encrypted.Bounds(), palette.Plan9)

	done := make(chan bool)
	go t.findSecret(done)
	t.takeSnapshots(done)

	return t.secret, t.snapshots
}

func (t *Turtle) findSecret(done chan bool) {
	bounds := t.encrypted.Bounds()
	for y := bounds.Max.Y - 1; y >= 0; y-- {
		for x := bounds.Max.X - 1; x >= 0; x-- {
			command := t.colorMap[t.encrypted.At(x, y)]
			if command == GoUp || command == GoLeft {
				t.drawNextLine(x, y, command, NoDir)
				time.Sleep(1 * time.Millisecond)
			}
		}
	}
	done <- true
}

func (t *Turtle) takeSnapshots(done chan bool) {
	i := 0

	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(10 * time.Millisecond)
			snapshot := image.NewPaletted(t.secret.Bounds(), palette.Plan9)
			snapshot.Pix = make([]uint8, len(t.secret.Pix))
			copy(snapshot.Pix, t.secret.Pix)
			t.snapshots[i] = snapshot
			i++
		}
	}
}

func (t *Turtle) drawNextLine(x, y, command, direction int) {
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

// Draw line until stop is reached
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
		command = t.colorMap[t.encrypted.At(curX, curY)]
	}

	if direction == Left || direction == Up {
		drawRect(t.secret, image.Pt(curX, curY), image.Pt(x, y))
	} else {
		drawRect(t.secret, image.Pt(x, y), image.Pt(curX, curY))
	}

	if command != Stop {
		t.drawNextLine(curX, curY, command, direction)
	}
}

// Fill in rectangle - though only used to draw lines
func drawRect(dst *image.Paletted, p1, p2 image.Point) {
	for x := p1.X; x <= p2.X; x++ {
		for y := p1.Y; y <= p2.Y; y++ {
			dst.Set(x, y, drawColor)
		}
	}
}
