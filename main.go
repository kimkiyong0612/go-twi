package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
)

var (
	assetsStore string = filepath.Join("images", "assets")
	outputStore string = filepath.Join("images", "output")
)

func main() {

	// open image file (860*574)
	bg, err := imaging.Open(filepath.Join(assetsStore, "example.jpeg"))
	if err != nil {
		log.Fatal(err)
	}
	_ = bg

	gopher, err := imaging.Open(filepath.Join(assetsStore, "gopher.png"))
	if err != nil {
		log.Fatal(err)
	}
	_ = gopher

	// resize by specifing px
	// dst := image.NewRGBA(image.Rect(0, 0, 500, 400))
	// draw.CatmullRom.Scale(dst, dst.Bounds(), bg, bg.Bounds(), draw.Over, nil)

	// crop preserving the aspect ratio
	// src := imaging.CropAnchor(gopher, 50, 100, imaging.Center)
	// output is 100*100
	src := imaging.Resize(gopher, 100, 50, imaging.Lanczos)

	// compose
	// offset := image.Pt(bg.Bounds().Dx()/6, bg.Bounds().Dy()/4)
	// // draw.Draw(dst, dst.Bounds(), bg, image.ZP, draw.Src)
	// draw.Draw(src, gopher.Bounds().Add(offset), gopher, image.Point{0, 0}, draw.Over)

	// create output file
	outputName := "new.png"
	pOutput, err := os.Create(filepath.Join(outputStore, outputName))
	if err != nil {
		log.Fatal(err)
	}
	defer pOutput.Close()

	// encode to output file
	err = png.Encode(pOutput, src)
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
