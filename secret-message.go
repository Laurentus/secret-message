package main

import (
	"fmt"
	"github.com/Laurentus/secret-message/turtle"
	"image"
	"image/color"
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

	decrypter := turtle.Decrypter(
		UpColor,
		LeftColor,
		StopColor,
		TurnRColor,
		TurnLColor,
	)

	fmt.Println("Image loaded and ready")
	secret := decrypter.Decrypt(m)
	fmt.Println("Secret decoded and saving to file")

	f, err := os.Create("secret.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, secret)
}
