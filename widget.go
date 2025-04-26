package dama

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	"github.com/gdamore/tcell/v2"
)

type WidgetState map[string]any

type DamaWidget interface {
	GetParent() *Container
	GetEventMap() event.EventMap
	GetKeybindings() event.Keybindings

	GetState() *WidgetState
	SetState(state *WidgetState)

	SetKeybinding(key tcell.Key, callback event.KeybindingCallback)
	SetKeybindings(callback event.KeybindingCallback, keys ...tcell.Key)
	SetEventCallback(eventname event.EventName, callback event.EventCallback)
	DamaElement
}

type Widget struct {
	*Element
	Parent      *Container
	EventMap    event.EventMap
	Keybindings event.Keybindings
	State       *WidgetState
}

func NewWidget() *Widget {
	widget := Widget{
		new(Element),
		nil,
		make(event.EventMap),
		make(event.Keybindings),
		nil,
	}

	return &widget
}

func (widget *Widget) GetParent() *Container {
	return widget.Parent
}

func (widget *Widget) GetBox() Box {
	box := widget.Element.GetBox()
	box.Element = widget
	return box
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

func (widget *Widget) SetKeybinding(key tcell.Key, cb event.KeybindingCallback) {
	widget.Keybindings[key] = cb
}

func (widget *Widget) SetKeybindings(cb event.KeybindingCallback, keys ...tcell.Key) {
	for _, key := range keys {
		widget.Keybindings[key] = cb
	}
}

func (widget *Widget) SetEventCallback(eventName event.EventName, cb event.EventCallback) {
	widget.EventMap[eventName] = cb
}

func (widget *Widget) Render(screen tcell.Screen, context lcontext.Context) {

}
