package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"

	log "github.com/sirupsen/logrus"
)

func main() {

	assetsStore := filepath.Join("images", "assets")
	outputStore := filepath.Join("images", "output")

	// open image file
	bgImage, err := os.Open(filepath.Join(assetsStore, "example.jpeg"))
	if err != nil {
		log.Fatal(err)
	}
	defer bgImage.Close()

	gopherImage, err := os.Open(filepath.Join(assetsStore, "gopher.png"))
	if err != nil {
		log.Fatal(err)
	}
	defer gopherImage.Close()

	// image decoed
	bg, _, err := image.Decode(bgImage)
	if err != nil {
		log.Fatal(err)
	}
	gopher, _, err := image.Decode(gopherImage)
	if err != nil {
		log.Fatal(err)
	}

	// resize
	// TODO: specify px
	dst := image.NewRGBA(image.Rect(0, 0, 500, 400))
	draw.CatmullRom.Scale(dst, dst.Bounds(), bg, bg.Bounds(), draw.Over, nil)

	// compose
	offset := image.Pt(bg.Bounds().Dx()/6, bg.Bounds().Dy()/4)
	draw.Draw(dst, dst.Bounds(), bg, image.ZP, draw.Src)
	draw.Draw(dst, gopher.Bounds().Add(offset), gopher, image.ZP, draw.Over)

	// create output file
	pOutput, err := os.Create(filepath.Join(outputStore, "new.png"))
	if err != nil {
		log.Fatal(err)
	}
	defer pOutput.Close()

	// encode to output file
	err = png.Encode(pOutput, dst)
	if err != nil {
		log.Fatal(err)
	}

}
