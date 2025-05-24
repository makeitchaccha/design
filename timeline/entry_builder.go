package timeline

import (
	"image"
)

type EntryBuilder struct {
	Entry
}

func NewEntryBuilder(avatar image.Image) *EntryBuilder {
	b := &EntryBuilder{}
	b.Avatar = avatar
	return b
}

func (b *EntryBuilder) AddSeries(series Series) *EntryBuilder {
	b.Series = append(b.Series, series)
	return b
}

func (b *EntryBuilder) Build() Entry {
	return b.Entry
}
