package timeline

func DefaultLayout() Layout {
	return Layout{
		HeadlineWidth:   100,
		TimelineWidth:   900,
		EntryHeight:     70,
		OnlineBarHeight: 20,
	}
}

type Layout struct {
	HeadlineWidth   float64
	TimelineWidth   float64
	EntryHeight     float64
	OnlineBarHeight float64
}

func (l Layout) Width() float64 {
	return l.HeadlineWidth + l.TimelineWidth
}

func (l Layout) Height(nEntries int) float64 {
	return l.EntryHeight * float64(nEntries)
}

func (l Layout) OnlineBarFillingFactor() float64 {
	return l.OnlineBarHeight / l.EntryHeight
}
