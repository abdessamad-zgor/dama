package dama

import (
	"github.com/gdamore/tcell/v2"
)

type Direction string

const (
	Left   Direction = "left"
	Right  Direction = "right"
	Top    Direction = "top"
	Bottom Direction = "bottom"
	Center Direction = "center"
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
	Trait
}

type scrollable_s struct {
	Viewport Viewport
}

func (scrollable *scrollable_s) GetViewport() Viewport {
	return scrollable.Viewport
}

func (scrollable *scrollable_s) Scroll(direction Direction) {
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

func (scrollable *scrollable_s) Render(widget Widget, screen tcell.Screen) {

}
