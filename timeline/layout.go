package timeline

var DefaultLayout = Layout{
	Margin: EdgeInsets{
		Top:    10,
		Right:  10,
		Bottom: 10,
		Left:   10,
	},
	LabelHeight:     30,
	HeadlineWidth:   100,
	TimelineWidth:   900,
	EntryHeight:     70,
	OnlineBarHeight: 20,
}

type Layout struct {
	Margin          EdgeInsets
	HeadlineWidth   float64
	TimelineWidth   float64
	LabelHeight     float64
	EntryHeight     float64
	OnlineBarHeight float64
}

func (l Layout) Width() float64 {
	return l.HeadlineWidth + l.TimelineWidth + l.Margin.Left + l.Margin.Right
}

func (l Layout) Height(nEntries int) float64 {
	return l.LabelHeight + l.EntryHeight*float64(nEntries) + l.Margin.Top + l.Margin.Bottom
}

func (l Layout) OnlineBarFillingFactor() float64 {
	return l.OnlineBarHeight / l.EntryHeight
}

type EdgeInsets struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}
