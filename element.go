package dama

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/gdamore/tcell/v2"
)

type DamaElement interface {
	DamaStyle
	Render(screen tcell.Screen, context lcontext.Context)
	GetBox() Box
	SetBox(x uint, y uint, width uint, height uint)
	SetTitle(title string)
	SetTag(tag rune)
	GetTitle() string
	GetTag() rune
	IsNavigable() bool
	Focus()
	Blur()
}

type Element struct {
	X      uint
	Y      uint
	Width  uint
	Height uint
	Tag    rune
	Title  string
	Styling
}

func (element *Element) Render(screen tcell.Screen, context lcontext.Context) {

}

func (element *Element) GetBox() Box {
	return Box{
		element.X,
		element.Y,
		element.Width,
		element.Height,
		nil,
	}
}

func (element *Element) SetBox(x uint, y uint, width uint, height uint) {
	element.X = x
	element.Y = y
	element.Width = width
	element.Height = height
}

func (element *Element) SetTag(tag rune) {
	element.Tag = tag
}

func (element *Element) SetTitle(title string) {
	element.Title = title
}

func (element *Element) GetTag() rune {
	return element.Tag
}

func (element *Element) GetTitle() string {
	return element.Title
}

func (element *Element) IsNavigable() bool {
	return element.Tag != rune(0) && element.Title != ""
}

func (element *Element) Focus() {

}

func (element *Element) Blur() {

}
