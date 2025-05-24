package timeline

import (
	"image/color"
	"time"
)

type SeriesBuilder struct {
	Series
}

func NewSeriesBuilder(fillingFactor float64, color color.Color) *SeriesBuilder {
	b := &SeriesBuilder{}
	b.FillingFactor = fillingFactor
	b.Color = color
	b.Sections = make([]Section, 0)
	return b
}

func (b *SeriesBuilder) AddSection(start, end time.Time, opts ...SectionOpt) *SeriesBuilder {
	section := Section{Start: start, End: end, Alpha: 1.0}
	section = section.applied(opts...)
	b.Sections = append(b.Sections, section)
	return b
}

func (b *SeriesBuilder) Build() Series {
	return b.Series
}
