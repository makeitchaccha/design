package timeline

import (
	"image"
	"image/color"
	"time"
)

type EntryBuilder struct {
	Entry
}

func NewEntryBuilder(avatar image.Image, color color.Color) *EntryBuilder {
	b := &EntryBuilder{}
	b.Avatar = avatar
	b.Color = color
	return b
}

func (b *EntryBuilder) AddSection(start, end time.Time, opts ...SectionOpt) *EntryBuilder {
	section := Section{Start: start, End: end}
	section = section.applied(opts...)
	b.Sections = append(b.Sections, section)
	return b
}

func (b *EntryBuilder) Build() Entry {
	return b.Entry
}
