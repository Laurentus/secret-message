package main

import (
	"fmt"
	"github.com/Laurentus/secret-message/turtle"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("wundernut.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	var UpColor = color.RGBA{7, 84, 19, 255}
	var LeftColor = color.RGBA{139, 57, 137, 255}
	var StopColor = color.RGBA{51, 69, 169, 255}
	var TurnRColor = color.RGBA{182, 149, 72, 255}
	var TurnLColor = color.RGBA{123, 131, 154, 255}

	decrypter := turtle.Decrypter()
	decrypter.SetCommand(turtle.GoUp, UpColor)
	decrypter.SetCommand(turtle.GoLeft, LeftColor)
	decrypter.SetCommand(turtle.Stop, StopColor)
	decrypter.SetCommand(turtle.TurnR, TurnRColor)
	decrypter.SetCommand(turtle.TurnL, TurnLColor)

	fmt.Println("Image loaded and ready")
	bounds := m.Bounds()
	fmt.Println("Image size is ", bounds.Min, bounds.Max)
	secret := decrypter.Decrypt(m)
	fmt.Println("Secret decoded and saving to file")

	f, err := os.Create("secret.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, secret)

	// Construct GIF from images.
	fmt.Println("Creating secret gif")
	outGif := &gif.GIF{}
	images := []image.Image{secret}
	for _, inImage := range images {
		outGif.Image = append(outGif.Image, inImage.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	g, err := os.Create("secret.gif")
	if err != nil {
		panic(err)
	}
	defer g.Close()
	gif.EncodeAll(g, outGif)
}
