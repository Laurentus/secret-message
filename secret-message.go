package main

import (
	"fmt"
	"github.com/Laurentus/secret-message/turtle"
	"image"
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

	fmt.Println("Image loaded and ready")
	secret := turtle.Decrypt(m)
	fmt.Println("Secret decoded and saving to file")

	f, err := os.Create("secret.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, secret)
}
