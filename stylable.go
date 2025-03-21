package dama

import (
	"github.com/gdamore/tcell/v2"
)

var DefaultStyle Style = Style {}


type Spacing struct {
	Top    uint
	Bottom uint
	Left   uint
	Right  uint
}

type Border struct {
	Color tcell.Color
	Bold  bool
}

type Style struct {
	Border     *Border
	Padding    Spacing
	Margin     Spacing
	Background tcell.Color
	Foreground  tcell.Color
	Bold   bool
	Italic bool
}

type DamaStylable interface {
    BorderColor(color tcell.Color) DamaStylable
    BorderBold(bold bool) DamaStylable
    Padding(padding Spacing) DamaStylable
    Margin(margin Spacing) DamaStylable
    Background(color tcell.Color) DamaStylable
    Foreground(color tcell.Color) DamaStylable
    Bold(bold bool) DamaStylable
    Italic(italic bool) DamaStylable

    GetStyle() Style
    SetStyle(style Style) 
}

type Stylable struct {
	Style Style
}

func (stylable *Stylable) BorderColor(color tcell.Color) *Stylable {
    stylable.Style.Border.Color = color
    return stylable
}

func (stylable *Stylable) BorderBold(bold bool) *Stylable {
    stylable.Style.Border.Bold = bold
    return stylable
}

func (stylable *Stylable) Padding(padding Spacing) *Stylable {
    stylable.Style.Padding = padding
    return stylable
}

func (stylable *Stylable) Margin(margin Spacing) *Stylable {
    stylable.Style.Margin = margin
    return stylable
}

func (stylable *Stylable) Background(color tcell.Color) *Stylable {
    stylable.Style.Background = color
    return stylable
}

func (stylable *Stylable) Foreground(color tcell.Color) *Stylable {
    stylable.Style.Foreground = color
    return stylable
}

func (stylable *Stylable) Bold(bold bool) *Stylable {
    stylable.Style.Bold = bold
    return stylable
}

func (stylable *Stylable) Italic(italic bool) *Stylable {
    stylable.Style.Italic = italic
    return stylable
}

func (stylable *Stylable) GetStyle() Style {
    return stylable.Style
}

func (stylable *Stylable) SetStyle(style Style) {
    stylable.Style = style
}
