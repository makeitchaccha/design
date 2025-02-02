package timeline

import (
	"time"

	"github.com/makeitchaccha/design"
)

type TimelineBuilder struct {
	Timeline
}

func NewTimelineBuilder(start, end time.Time) *TimelineBuilder {
	b := &TimelineBuilder{}
	b.StartTime = start
	b.EndTime = end
	b.Layout = DefaultLayout()
	b.MainTics, b.SubTics = CalculateTics(end.Sub(start)) // default tics
	b.Margin = design.EdgeInsets{
		Top:    10,
		Left:   10,
		Right:  10,
		Bottom: 10,
	}
	return b
}

func (b *TimelineBuilder) AddEntries(entries ...Entry) *TimelineBuilder {
	b.Entries = append(b.Entries, entries...)
	return b
}

func (b *TimelineBuilder) SetPadding(padding design.EdgeInsets) *TimelineBuilder {
	b.Padding = padding
	return b
}

func (b *TimelineBuilder) SetMargin(margin design.EdgeInsets) *TimelineBuilder {
	b.Margin = margin
	return b
}

func (b *TimelineBuilder) SetTitle(title design.TextBox) *TimelineBuilder {
	b.Title = title
	return b
}

func (b *TimelineBuilder) SetIndicator(indicator time.Time) *TimelineBuilder {
	b.Indicator = indicator
	return b
}

func (b *TimelineBuilder) SetLayout(layout Layout) *TimelineBuilder {
	b.Layout = layout
	return b
}

func (b *TimelineBuilder) SetMainTics(tics Tics) *TimelineBuilder {
	b.MainTics = tics
	return b
}

func (b *TimelineBuilder) SetSubTics(tics Tics) *TimelineBuilder {
	b.SubTics = tics
	return b
}

func (b *TimelineBuilder) Build() Timeline {
	return b.Timeline
}
