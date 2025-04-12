package dama

import (
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

var DefaultStyle Style = Style{}

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
	Foreground tcell.Color
	Bold       bool
	Italic     bool
}

type DamaStyle interface {
	BorderColor(color tcell.Color) DamaStyle
	BorderBold(bold bool) DamaStyle
	Padding(padding Spacing) DamaStyle
	Margin(margin Spacing) DamaStyle
	Background(color tcell.Color) DamaStyle
	Foreground(color tcell.Color) DamaStyle
	Bold(bold bool) DamaStyle
	Italic(italic bool) DamaStyle

	GetStyle() Style
	SetStyle(style Style)
}

type Styling struct {
	Style Style
}

func (styling *Styling) BorderColor(color tcell.Color) DamaStyle {
	logger.Logger.Println("styling before: ", styling)
	logger.Logger.Println("styling after: ", styling)
	if styling.Style.Border == nil {
		styling.Style.Border = new(Border)
	}
	styling.Style.Border.Color = color
	return styling
}

func (styling *Styling) BorderBold(bold bool) DamaStyle {
	if styling.Style.Border == nil {
		styling.Style.Border = new(Border)
	}
	styling.Style.Border.Bold = bold
	return styling
}

func (styling *Styling) Padding(padding Spacing) DamaStyle {
	styling.Style.Padding = padding
	return styling
}

func (styling *Styling) Margin(margin Spacing) DamaStyle {
	styling.Style.Margin = margin
	return styling
}

func (styling *Styling) Background(color tcell.Color) DamaStyle {
	styling.Style.Background = color
	return styling
}

func (styling *Styling) Foreground(color tcell.Color) DamaStyle {
	styling.Style.Foreground = color
	return styling
}

func (styling *Styling) Bold(bold bool) DamaStyle {
	styling.Style.Bold = bold
	return styling
}

func (styling *Styling) Italic(italic bool) DamaStyle {
	styling.Style.Italic = italic
	return styling
}

func (styling *Styling) GetStyle() Style {
	return styling.Style
}

func (styling *Styling) SetStyle(style Style) {
	styling.Style = style
}
