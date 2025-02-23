package design

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

func DefaultLabelFontFace() font.Face {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("something went wrong?")
	}

	return truetype.NewFace(font, &truetype.Options{Size: 15})
}

func DefaultTitleFontFace() font.Face {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("something went wrong?")
	}

	return truetype.NewFace(font, &truetype.Options{Size: 20})
}
