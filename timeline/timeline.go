package timeline

import (
	"image"
	"image/color"
	"io"
	"time"

	"github.com/fogleman/gg"
	"github.com/makeitchaccha/design"
	"github.com/makeitchaccha/rendering/chart/timeline"
	"github.com/makeitchaccha/rendering/layout"
)

type Timeline struct {
	Title     design.TextBox
	StartTime time.Time
	Indicator time.Time
	EndTime   time.Time
	Entries   []Entry
	Margin    design.EdgeInsets
	Padding   design.EdgeInsets
	Layout    Layout
	MainTics  Tics
	SubTics   Tics
}

func (t Timeline) Width() float64 {
	return t.Layout.Width() + t.Margin.Horizontal() + t.Padding.Horizontal()
}

func (t Timeline) Height() float64 {
	return t.GridTop() + t.Layout.Height(len(t.Entries)) + t.Margin.Bottom + t.Padding.Bottom
}

func (t Timeline) GridTop() float64 {
	height := t.Margin.Top + t.Padding.Top
	if t.Title.Valid() {
		height += t.Title.Height()
	}
	if t.MainTics.Valid() {
		height += t.MainTics.Label.Height()
	}
	if t.SubTics.Valid() {
		height += t.SubTics.Label.Height()
	}
	return height
}

func (t Timeline) GridLeft() float64 {
	return t.Margin.Left
}

func (t Timeline) Generate() io.Reader {
	nEntries := len(t.Entries)

	width := t.Width()
	height := t.Height()

	dc := gg.NewContext(int(width), int(height))
	dc.SetColor(color.White)
	dc.Clear()

	cellHeights := make([]float64, nEntries)
	for i := range t.Entries {
		cellHeights[i] = t.Layout.EntryHeight
	}
	grid := layout.NewGrid(t.GridLeft(), t.GridTop(), []float64{t.Layout.HeadlineWidth, t.Layout.TimelineWidth}, cellHeights)

	headerGrid, _ := grid.ColAsSubgrid(0)

	if headerGrid.Bounds().Dx() > 0 {
		for idx, f := range headerGrid.ForEachCellRenderFunc {
			entry := t.Entries[idx.Row]
			f(dc, func(dc *gg.Context, x, y, w, h float64) error {
				dc.Push()
				dc.DrawCircle(x+w/2, y+h/2, float64(entry.Avatar.Bounds().Dx())/2)
				dc.Clip()
				dc.DrawImageAnchored(entry.Avatar, int(x+w/2), int(y+h/2), 0.5, 0.5)
				dc.ResetClip()
				dc.Pop()
				return nil
			})
		}
	}

	timelineGrid, _ := grid.ColAsSubgrid(1)

	timelineBounds := timelineGrid.Bounds()

	// draw title
	if t.Title.Valid() {
		dc.Push()
		dc.SetColor(color.RGBA{66, 66, 66, 255})
		dc.SetFontFace(t.Title.Font)
		dc.DrawStringAnchored(t.Title.Content, timelineBounds.Cx(), timelineBounds.Min.Y-t.Padding.Top-t.MainTics.Label.Height()-t.SubTics.Label.Height(), 0.5, 0)
		dc.Pop()
	}

	// draw tics on an hour intervals

	total := t.EndTime.Sub(t.StartTime).Seconds()
	main, sub := t.MainTics, t.SubTics
	anchor := 0.0
	for _, tics := range []Tics{main, sub} {
		if !tics.shouldDraw() {
			// skip if no tics are set
			continue
		}

		_, offset := t.StartTime.Zone()
		dOffset := time.Duration(offset) * time.Second
		current := t.StartTime.Add(dOffset).Truncate(tics.Interval).Add(-dOffset)
		if current.Before(t.StartTime) {
			current = current.Add(tics.Interval) // move to the next hour to avoid drawing a tic at the start
		}
		for ; !current.After(t.EndTime); current = current.Add(tics.Interval) {
			x := timelineBounds.Min.X + (timelineBounds.Dx() * current.Sub(t.StartTime).Seconds() / total)
			// draw a tic and label on the top

			if tics.Label.Valid() {
				dc.SetFontFace(tics.Label.Font)
				dc.SetColor(color.RGBA{66, 66, 66, 255})
				dc.DrawStringAnchored(current.Format(tics.Label.Content), x, timelineBounds.Min.Y-t.Padding.Top-anchor, 0.5, 0)
			}

			dc.SetColor(tics.Color)
			dc.DrawLine(x, timelineBounds.Min.Y, x, timelineBounds.Max.Y)
			dc.Stroke()
		}
		anchor += tics.Label.Height()
	}

	builder := timeline.NewTimelineBuilder()

	for _, entry := range t.Entries {

		entryBuilder := timeline.NewEntryBuilder()
		for _, series := range entry.Series {
			if series.Color == nil {
				series.Color = extractMainColor(entry.Avatar)
			}

			seriesBuilder := timeline.NewSeriesBuilder(series.FillingFactor, series.Color)

			for _, section := range series.Sections {

				s := section.Start.Sub(t.StartTime).Seconds() / total
				e := section.End.Sub(t.StartTime).Seconds() / total
				seriesBuilder.AddSection(s, e, timeline.WithAlpha(section.Alpha))
			}

			entryBuilder.AddSeries(seriesBuilder.Build())
		}

		builder.AddEntry(entryBuilder.Build())
	}

	builder.Build().RenderInGrid(dc, timelineGrid)

	// Draw start and end time vertical lines
	dc.SetColor(color.RGBA{0, 105, 92, 255})
	dc.DrawLine(timelineBounds.Min.X, timelineBounds.Min.Y, timelineBounds.Min.X, timelineBounds.Max.Y)
	dc.DrawLine(timelineBounds.Max.X, timelineBounds.Min.Y, timelineBounds.Max.X, timelineBounds.Max.Y)
	dc.Stroke()

	// Draw indicator line if it is set
	if !t.Indicator.IsZero() {
		x := timelineBounds.Min.X + (timelineBounds.Dx() * t.Indicator.Sub(t.StartTime).Seconds() / total)
		dc.SetColor(color.RGBA{128, 0, 0, 255})
		dc.DrawLine(x, timelineBounds.Min.Y, x, timelineBounds.Max.Y)
		dc.Stroke()
	}

	r, w := io.Pipe()
	go func() {
		dc.EncodePNG(w)
		w.Close()
	}()

	return r
}

func extractMainColor(img image.Image) color.Color {
	bounds := img.Bounds()
	colorCount := make(map[color.Color]int)

	// HACK: Prevent all white or transparent images from being considered
	colorCount[color.Black] = 1

	// Iterate over each pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()
			// Ignore white like colors
			if r > 0xBFFF && g > 0xBFFF && b > 0xBFFF {
				continue
			}
			// also ignore transparent pixels
			if a == 0 {
				continue
			}
			colorCount[c]++
		}
	}

	// Find the most frequent color
	var mainColor color.Color
	maxCount := 0
	for c, count := range colorCount {
		if count > maxCount {
			maxCount = count
			mainColor = c
		}
	}

	return mainColor
}
