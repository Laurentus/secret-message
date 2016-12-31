package turtle

import (
	"fmt"
	"image"
	"image/color"
)

var drawColor = color.RGBA{0, 0, 0, 255}
var secret *image.RGBA

const (
	Up = iota
	Right
	Down
	Left
)

const (
	GoUp = iota
	GoLeft
	TurnR
	TurnL
	Stop
	NoCommand
)

var UpColor = color.RGBA{7, 84, 19, 255}
var LeftColor = color.RGBA{139, 57, 137, 255}
var StopColor = color.RGBA{51, 69, 169, 255}
var TurnRColor = color.RGBA{182, 149, 72, 255}
var TurnLColor = color.RGBA{123, 131, 154, 255}

func colorCommand(col color.Color) int {
	if isEqualColor(col, UpColor) {
		return GoUp
	}
	if isEqualColor(col, LeftColor) {
		return GoLeft
	}
	if isEqualColor(col, StopColor) {
		return Stop
	}
	if isEqualColor(col, TurnRColor) {
		return TurnR
	}
	if isEqualColor(col, TurnLColor) {
		return TurnL
	}
	return NoCommand
}

func isEqualColor(col1, col2 color.Color) bool {
	r1, g1, b1, _ := col1.RGBA()
	r2, g2, b2, _ := col2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2
}

func Decrypt(m image.Image) image.Image {
	bounds := m.Bounds()
	secret = image.NewRGBA(bounds)
	fmt.Println(m.ColorModel())
	drawing := false

	r, g, b, a := color.RGBA{7, 84, 19, 255}.RGBA()
	fmt.Println("Success", r, g, b, a)

	fmt.Println("Image size is ", bounds.Min, bounds.Max)
	for y := bounds.Max.Y - 1; y >= 0; y-- {
		for x := bounds.Max.X - 1; x >= 0; x-- {
			command := colorCommand(m.At(x, y))
			if drawing {
				secret.Set(x, y, drawColor)
				drawing = command == NoCommand || command == GoLeft
			} else if command == GoLeft {
				fmt.Println("Drawing")
				drawing = true
			}
		}
	}

	return secret
}

// Draw a line between points p1 and p2
func DrawRect(p1, p2 image.Point) {
	for x := p1.X; x < p2.X; x++ {
		for y := p1.Y; y < p2.Y; y++ {
			secret.Set(x, y, drawColor)
		}
	}
}
