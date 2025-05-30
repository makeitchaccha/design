package timeline

import (
	"image"
	"time"
)

// Entry represents a timeline entry.
// It contains an avatar and a timeline entry.
type Entry struct {
	Avatar image.Image
	Series []Series
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
