package timeline

import "time"

type TimelineBuilder struct {
	Timeline
}

func NewTimelineBuilder(start, end time.Time) *TimelineBuilder {
	b := &TimelineBuilder{}
	b.StartTime = start
	b.EndTime = end
	b.Layout = DefaultLayout
	return b
}

func (b *TimelineBuilder) AddEntries(entries ...Entry) *TimelineBuilder {
	b.Entries = append(b.Entries, entries...)
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

func (b *TimelineBuilder) Build() Timeline {
	return b.Timeline
}
