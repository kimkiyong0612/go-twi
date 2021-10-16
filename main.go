package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"

	log "github.com/sirupsen/logrus"
)

var (
	assetsStore string = filepath.Join("images", "assets")
	outputStore string = filepath.Join("images", "output")
)

func main() {

	// open image file
	bg, err := openAssetImage("example.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	gopher, err := openAssetImage("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	_ = gopher

	// resize by specifing px
	dst := image.NewRGBA(image.Rect(0, 0, 500, 400))
	draw.CatmullRom.Scale(dst, dst.Bounds(), bg, bg.Bounds(), draw.Over, nil)

	// compose
	offset := image.Pt(bg.Bounds().Dx()/6, bg.Bounds().Dy()/4)
	// draw.Draw(dst, dst.Bounds(), bg, image.ZP, draw.Src)
	draw.Draw(dst, gopher.Bounds().Add(offset), gopher, image.Point{0, 0}, draw.Over)

	// create output file
	outputName := "new.png"
	pOutput, err := os.Create(filepath.Join(outputStore, outputName))
	if err != nil {
		log.Fatal(err)
	}
	defer pOutput.Close()

	// encode to output file
	err = png.Encode(pOutput, dst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSUCCESS: LOOK %s\n", outputName)
}

func openAssetImage(filename string) (image.Image, error) {
	f, err := os.Open(filepath.Join(assetsStore, filename))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	return i, err
}
