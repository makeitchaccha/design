package timeline

import "image/color"

type Series struct {
	FillingFactor float64
	Color         color.Color
	Sections      []Section
}
