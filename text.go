package design

import (
	"strings"

	"golang.org/x/image/font"
)

type TextBox struct {
	Content string
	Font    font.Face
}

func (t TextBox) Valid() bool {
	return t.Content != "" && t.Font != nil
}

func (t TextBox) Lines() int {
	if t.Content == "" {
		return 0
	}
	return strings.Count(t.Content, "\n") + 1
}

func (t TextBox) Height() float64 {
	if !t.Valid() {
		return 0
	}

	return float64(t.Lines() * t.Font.Metrics().Height.Ceil())
}
