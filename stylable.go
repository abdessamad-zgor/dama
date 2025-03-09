package dama

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/gdamore/tcell/v2"
)

type StyleMap map[event.EventName]tcell.Style

type Stylable interface {
	GetStyleMap() StyleMap
	GetEventStyle(event event.EventName) tcell.Style
	SetEventStyle(event event.EventName, tstyle tcell.Style)
    GetContextStyle(context lcontext.Context) tcell.Style
}

type Style struct {
	StyleMap StyleMap
}

func (style Style) GetStyleMap() StyleMap {
	return style.StyleMap
}

func (style Style) GetEventStyle(event event.EventName) tcell.Style {
	return style.StyleMap[event]
}

func (style Style) SetEventStyle(event event.EventName, tstyle tcell.Style) {
	style.StyleMap[event] = tstyle
}

func (style Style) GetContextStyle(context lcontext.Context) tcell.Style {
    return tcell.StyleDefault
}
