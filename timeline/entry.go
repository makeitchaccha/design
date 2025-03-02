package timeline

import (
	"image"
	"image/color"
	"time"
)

type Entry struct {
	Avatar   image.Image
	Color    color.Color
	Sections []Section
}

// Section represents a time Section in a timeline entry.
// It is defined by a start and end time.
type Section struct {
	Start time.Time
	End   time.Time
	Alpha float64 // 0(transparent) to 1(opaque)
}

func (s Section) applied(opts ...SectionOpt) Section {
	for _, opt := range opts {
		opt(&s)
	}
	return s
}

type SectionOpt func(*Section)

func WithAlpha(alpha float64) SectionOpt {
	return func(s *Section) {
		s.Alpha = alpha
	}
}
