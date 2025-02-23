package core

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/gdamore/tcell/v2"
)

type WidgetState map[string]any

type DamaWidget interface {
	GetParent() *Container
	GetBox() Box
	GetEventMap() event.EventMap
	GetKeybindings() event.Keybindings

	GetState() *WidgetState
	SetState(state *WidgetState)

	SetEventListener(key tcell.Key, eventName event.EventName, cb event.Callback)
    Stylable
    Element
}

type Widget struct {
	Parent      *Container
	X           uint
	Y           uint
	Width       uint
	Height      uint
	EventMap    event.EventMap
	Keybindings event.Keybindings
	State       *WidgetState
    Style
}

func (widget *Widget) GetParent() *Container {
	return widget.Parent
}

func (widget *Widget) GetBox() Box {
	return Box{widget.X, widget.Y, widget.Width, widget.Height, widget}
}

func (widget *Widget) GetEventMap() event.EventMap {
	return widget.EventMap
}

func (widget *Widget) GetKeybindings() event.Keybindings {
	return widget.Keybindings
}

func (widget *Widget) GetState() *WidgetState {
	return widget.State
}

func (widget *Widget) SetState(state *WidgetState) {
	widget.State = state
}

func (widget *Widget) SetEventListener(key tcell.Key, eventName event.EventName, cb event.Callback) {
	widget.Keybindings[key] = eventName
	widget.EventMap[eventName] = cb
}

func Render(screen tcell.Screen, context lcontext.context) {

}
