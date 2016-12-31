package turtle

import (
	"image"
)

func Decrypt(m image.Image) image.Image {
	secret := image.NewRGBA(m.Bounds())
	return secret
}
