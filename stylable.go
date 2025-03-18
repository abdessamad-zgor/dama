package dama

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/gdamore/tcell/v2"
)

type StyleMap map[event.EventName]tcell.Style

type DamaStylable interface {
	GetStyleMap() StyleMap
	GetEventStyle(event event.EventName) tcell.Style
	SetEventStyle(event event.EventName, tstyle tcell.Style)
  GetContextStyle(context lcontext.Context) tcell.Style
}

type Stylable struct {
	StyleMap StyleMap
}

func (stylable Stylable) GetStyleMap() StyleMap {
	return stylable.StyleMap
}

func (stylable Stylable) GetEventStyle(event event.EventName) tcell.Style {
	return stylable.StyleMap[event]
}

func (stylable Stylable) SetEventStyle(event event.EventName, tstyle tcell.Style) {
	stylable.StyleMap[event] = tstyle
}

func (stylable Stylable) GetContextStyle(context lcontext.Context) tcell.Style {
    return tcell.StyleDefault
}
