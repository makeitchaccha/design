package design

type EdgeInsets struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

func (e EdgeInsets) Horizontal() float64 {
	return e.Left + e.Right
}

func (e EdgeInsets) Vertical() float64 {
	return e.Top + e.Bottom
}
