package dama

import (
	dtraits "github.com/abdessamad-zgor/dama/traits"
	"github.com/gdamore/tcell/v2"
)

type DamaElement interface {
	dtraits.DamaStyle
	Render(screen tcell.Screen)
	GetBox() Box
	SetBox(x int, y int, width int, height int)
	SetTitle(title string)
	SetTag(tag rune)
	GetTitle() string
	GetTag() rune
	IsNavigable() bool
	RenderTag(screen tcell.Screen)
	RenderTitle(screen tcell.Screen)
	Focus()
	Blur()
	IsFocused() bool
}

type Element struct {
	X     	int
	Y      	int
	Width  	int
	Height 	int
	Tag    	rune
	Title  	string
	Focused	bool
	dtraits.Styling
}

func (element *Element) Render(screen tcell.Screen) {

}

func (element *Element) RenderTag(screen tcell.Screen) {
	if element.Tag != rune(0) {
		offset := 2
		tagText := "["+string(element.Tag)+"]"
		color := tcell.ColorDefault
		if element.Style.Border != nil {
			color = element.Style.Border.Color
		}
		for i, char := range tagText {
			screen.SetContent(int(element.X) + offset + i, int(element.Y), char, nil, tcell.StyleDefault.Foreground(color))
		}
	}
}

func (element *Element) RenderTitle(screen tcell.Screen) {
	if element.Title != "" {
		offset := 2
		if element.Tag != rune(0) {
			offset += 4
		}
		color := tcell.ColorDefault
		if element.Style.Border != nil {
			color = element.Style.Border.Color
		}
		for i, char := range element.Title {
			screen.SetContent(int(element.X) + offset + i, int(element.Y), char, nil, tcell.StyleDefault.Foreground(color))
		}
	}	
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

func (element *Element) SetBox(x int, y int, width int, height int) {
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

func (element *Element) IsFocused() bool {
	return element.Focused
}
