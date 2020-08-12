package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

func importImage(path string) *ebiten.Image {
	importedImage, _, _ := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	return importedImage
}

func importFont(size float64) *font.Face {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 100
	mplusNormalFont := truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return &mplusNormalFont
}

//lint:ignore U1000 Stubs
func importSound(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
}
