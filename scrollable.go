package dama

type Direction string

const (
	Left   Direction = "left"
	Right  Direction = "right"
	Top    Direction = "top"
	Bottom Direction = "bottom"
)

type Viewport struct {
	OffsetX int
	OffsetY int
	Width   int
	Height  int
}

type DamaScrollable interface {
	GetViewport() Viewport
	Scroll(direction Direction)
}

type Scrollable struct {
	Viewport Viewport
}

func (scrollable *Scrollable) GetViewport() Viewport {
	return scrollable.Viewport
}

func (scrollable *Scrollable) Scroll(direction Direction) {
	switch direction {
	case Left:
		if scrollable.Viewport.OffsetX != scrollable.Viewport.Width {
			scrollable.Viewport.OffsetX += 1
		}
	case Right:
		if scrollable.Viewport.OffsetX != 0 {
			scrollable.Viewport.OffsetX -= 1
		}
	case Top:
		if scrollable.Viewport.OffsetY != 0 {
			scrollable.Viewport.OffsetY -= 1
		}
	case Bottom:
		if scrollable.Viewport.OffsetY != scrollable.Viewport.Height {
			scrollable.Viewport.OffsetY += 1
		}
	}
}
