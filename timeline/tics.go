package timeline

import (
	"image/color"
	"time"

	"github.com/makeitchaccha/design"
)

type Tics struct {
	Interval time.Duration
	Color    color.Color
	Label    design.TextBox
}

func (t Tics) Valid() bool {
	return t.shouldDraw() && t.Label.Valid()
}

func (t Tics) shouldDraw() bool {
	return t.Interval != 0 && t.Color != nil
}

func CalculateTics(d time.Duration) (Tics, Tics) {
	base := chooseBaseDuration(d)
	upgrade := upgradeDuration(base)

	return Tics{
			Interval: base,
			Label:    design.TextBox{Content: ChooseFormat(base), Font: design.DefaultLabelFontFace()},
			Color:    color.RGBA{200, 200, 200, 255},
		},
		Tics{
			Interval: upgrade,
			Label:    design.TextBox{Content: ChooseFormat(upgrade), Font: design.DefaultLabelFontFace()},
			Color:    color.RGBA{100, 100, 100, 255},
		}
}

func ChooseFormat(d time.Duration) string {
	if d < 60*time.Second {
		return "15:04:05"
	}
	if d < 24*time.Hour {
		return "15:04"
	}

	return "01/02"
}

func chooseBaseDuration(d time.Duration) time.Duration {
	presets := []time.Duration{
		10 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		30 * time.Minute,
		1 * time.Hour,
		4 * time.Hour,
		8 * time.Hour,
		12 * time.Hour,
		24 * time.Hour,
	}

	for _, p := range presets {
		if d < 12*p {
			return p
		}
	}

	return 24 * time.Hour
}

func upgradeDuration(d time.Duration) time.Duration {
	if d < 24*time.Hour {
		return 24 * time.Hour
	}
	if d < 7*24*time.Hour {
		return 7 * 24 * time.Hour // 1 week (?)
	}

	return 30 * 24 * time.Hour // (??) would the call last for 30 days?
}
