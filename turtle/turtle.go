package turtle

import (
	"image"
	"image/color"
	"image/color/palette"
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
	// Allocate enough space for pictures as dynamic allocation can be slow
	snapshots := make([]*image.Paletted, 200)
	return &Turtle{colorMap: colorMap, snapshots: snapshots}
}

func (t *Turtle) SetCommand(command int, clr color.Color) {
	t.colorMap[clr] = command
}

func (t *Turtle) Decrypt(m image.Image) (image.Image, []*image.Paletted) {
	t.encrypted = m
	t.secret = image.NewPaletted(t.encrypted.Bounds(), palette.Plan9)

	t.findSecret()

	return t.secret, t.snapshots
}

func (t *Turtle) findSecret() {
	bounds := t.encrypted.Bounds()

	// Take starting snapshot to provide nice black start screen
	t.takeSnapshot(0)
	snapshotIdx := 1

	// Iterate diagonal by diagonal to provide nicer gif
	x, y, total := nextDiagonalPoint(0, 0, 0)
	for total < bounds.Max.X+bounds.Max.Y {
		x, y, total = nextDiagonalPoint(x, y, total)
		// Skip pixels outside the image
		if y >= bounds.Max.Y || x >= bounds.Max.X {
			continue
		}

		command := t.colorMap[t.encrypted.At(x, y)]
		if command == GoUp || command == GoLeft {
			t.drawNextLine(x, y, command, NoDir)
			t.takeSnapshot(snapshotIdx)
			snapshotIdx++
		}
	}
}

func nextDiagonalPoint(x, y, total int) (int, int, int) {
	if x == 0 {
		total++
		x = total
		y = 0
	} else {
		y++
		x--
	}

	return x, y, total
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

func (t *Turtle) takeSnapshot(snapshotIdx int) {
	snapshot := image.NewPaletted(t.secret.Bounds(), palette.Plan9)
	snapshot.Pix = make([]uint8, len(t.secret.Pix))
	copy(snapshot.Pix, t.secret.Pix)
	t.snapshots[snapshotIdx] = snapshot
}
