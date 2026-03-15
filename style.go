package dama

import (
	_ "github.com/samazee/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type PropertyName string

const (
	BorderColor 	PropertyName 	= "border-color"
	BorderBold 		PropertyName 	= "border-bold"
	Background		PropertyName	= "background"
	Foreground		PropertyName	= "foreground"
	TextBold		PropertyName	= "text-bold"
	TextItalic		PropertyName	= "text-italic"
)

type StyleProperties map[PropertyName]any

type Style interface {
	BorderColor(color tcell.Color) Style
	BorderBold(bold bool) Style
	Background(color tcell.Color) Style
	Foreground(color tcell.Color) Style
	TextBold(bold bool) Style
	TextItalic(italic bool) Style

	GetStyleProperties() StyleProperties
}

func NewStyle() Style {
	var style StyleProperties = make(map[PropertyName]any)
	return &style
}

func (props *StyleProperties) BorderColor(color tcell.Color) Style {
	(*props)[BorderColor] = color
	return props
}

func (props *StyleProperties) BorderBold(bold bool) Style {
	(*props)[BorderBold] = bold
	return props
}

func (props *StyleProperties) Background(color tcell.Color) Style {
	(*props)[Background] = color
	return props
}

func (props *StyleProperties) Foreground(color tcell.Color) Style {
	(*props)[Foreground] = color
	return props
}

func (props *StyleProperties) TextBold(bold bool) Style {
	(*props)[TextBold] = bold
	return props
}

func (props *StyleProperties) TextItalic(italic bool) Style {
	(*props)[TextItalic] = italic
	return props
}

func (props *StyleProperties) GetStyleProperties() StyleProperties {
	return *props
}

func (props StyleProperties) GetBorderColor() (tcell.Color, bool) {
	value, ok := props[BorderColor]
	avalue, _ := value.(tcell.Color)
	return avalue, ok
}

func (props StyleProperties) GetBorderBold() (bool, bool)  {
	value, ok := props[BorderBold]
	avalue, _ := value.(bool)
	return avalue, ok
}

func (props StyleProperties) GetBackground(color tcell.Color) (tcell.Color, bool) {
	value, ok := props[Background]
	avalue, _ := value.(tcell.Color)
	return avalue, ok
}

func (props StyleProperties) GetForeground(color tcell.Color) (tcell.Color, bool) {
	value, ok := props[Foreground]
	avalue, _ := value.(tcell.Color)
	return avalue, ok
}

func (props StyleProperties) GetTextBold(bold bool) (bool, bool) {
	value, ok := props[TextBold]
	avalue, _ := value.(bool)
	return avalue, ok
}

func (props StyleProperties) GetTextItalic(italic bool) (bool, bool) {
	value, ok := props[TextItalic]
	avalue, _ := value.(bool)
	return avalue, ok
}
