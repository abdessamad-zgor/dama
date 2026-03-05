package dama

import (
	dtraits "github.com/abdessamad-zgor/dama/traits"
	"github.com/gdamore/tcell/v2"
)

type Element interface {
	dtraits.Style
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

type element_s struct {
	X     	int
	Y      	int
	Width  	int
	Height 	int
	Tag    	rune
	Title  	string
	Focused	bool
	dtraits.Style
}

func (element *element_s) Render(screen tcell.Screen) {

}

func (element *element_s) RenderTag(screen tcell.Screen) {
	if element.Tag != rune(0) {
		offset := 2
		tagText := "["+string(element.Tag)+"]"
		style := tcell.StyleDefault
		elementStyle := element.GetStyleProperties()
		if borderColor, ok := elementStyle.GetBorderColor(); ok {
			style = style.Foreground(borderColor)
		}
		for i, char := range tagText {
			screen.SetContent(int(element.X) + offset + i, int(element.Y), char, nil, style)
		}
	}
}

func (element *element_s) RenderTitle(screen tcell.Screen) {
	if element.Title != "" {
		offset := 2
		if element.Tag != rune(0) {
			offset += 4
		}
		style := tcell.StyleDefault
		elementStyle := element.GetStyleProperties()
		if borderColor, ok := elementStyle.GetBorderColor(); ok {
			style = style.Foreground(borderColor)
		}
		for i, char := range element.Title {
			screen.SetContent(int(element.X) + offset + i, int(element.Y), char, nil, style)
		}
	}	
}

func (element *element_s) GetBox() Box {
	return Box{
		element.X,
		element.Y,
		element.Width,
		element.Height,
		nil,
	}
}

func (element *element_s) SetBox(x int, y int, width int, height int) {
	element.X = x
	element.Y = y
	element.Width = width
	element.Height = height
}

func (element *element_s) SetTag(tag rune) {
	element.Tag = tag
}

func (element *element_s) SetTitle(title string) {
	element.Title = title
}

func (element *element_s) GetTag() rune {
	return element.Tag
}

func (element *element_s) GetTitle() string {
	return element.Title
}

func (element *element_s) IsNavigable() bool {
	return element.Tag != rune(0) && element.Title != ""
}

func (element *element_s) Focus() {

}

func (element *element_s) Blur() {

}

func (element *element_s) IsFocused() bool {
	return element.Focused
}
