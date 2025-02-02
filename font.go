package design

import (
	"os"
	"os/exec"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var DefaultLabelFontFace font.Face
var DefaultTitleFontFace font.Face

func init() {
	// TODO: support windows maybe?
	filename, err := FindFontFile("NotoSans", "Regular")
	if err != nil {
		panic("could not find font file")
	}

	font, err := LoadTrueTypeFont(filename)

	if err != nil {
		panic("could not load font")
	}

	DefaultLabelFontFace = truetype.NewFace(font, &truetype.Options{Size: 15})
	DefaultTitleFontFace = truetype.NewFace(font, &truetype.Options{Size: 20})
}

func FindFontFile(family string, style string) (string, error) {
	// todo: support windows
	filename, err := exec.Command("fc-match", family+":style="+style, "-f", "%{file}").Output()
	return string(filename), err
}

func LoadTrueTypeFont(filename string) (*truetype.Font, error) {

	fontByte, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	font, err := truetype.Parse(fontByte)
	if err != nil {
		return nil, err
	}
	return font, nil
}
