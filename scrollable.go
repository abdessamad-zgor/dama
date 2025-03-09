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

type Scrollable interface {
	GetViewport() Viewport
	Scroll(direction Direction)
}

type Scroll struct {
	Viewport Viewport
}

func (scroll *Scroll) GetViewport() Viewport {
	return scroll.Viewport
}

func (scroll *Scroll) Scroll(direction Direction) {
	switch direction {
	case Left:
		if scroll.Viewport.OffsetX != scroll.Viewport.Width {
			scroll.Viewport.OffsetX += 1
		}
	case Right:
		if scroll.Viewport.OffsetX != 0 {
			scroll.Viewport.OffsetX -= 1
		}
	case Top:
		if scroll.Viewport.OffsetY != 0 {
			scroll.Viewport.OffsetY -= 1
		}
	case Bottom:
		if scroll.Viewport.OffsetY != scroll.Viewport.Height {
			scroll.Viewport.OffsetY += 1
		}
	}
}
