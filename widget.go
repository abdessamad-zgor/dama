package dama

import (
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
)

type WidgetState map[string]any

type DamaWidget interface {
	GetParent() *Container

	SetKeybinding(pattern string, callback devent.Callback)
	//SetKeybindings(callback devent.Callback, keys ...tcell.Key)
	SetAppEvent(eventname devent.EventName, callback devent.Callback)
	DamaElement
}

type Widget struct {
	*Element
	Parent      *Container
	Events 		dutils.List[devent.DamaEvent]
}

func NewWidget() *Widget {
	widget := Widget{
		new(Element),
		nil,
		dutils.NewList[devent.DamaEvent](),
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

func (widget *Widget) SetKeybinding(key tcell.Key, cb devent.KeybindingCallback) {
	widget.Keybindings[key] = cb
}

func (widget *Widget) SetKeybindings(cb devent.KeybindingCallback, keys ...tcell.Key) {
	for _, key := range keys {
		widget.Keybindings[key] = cb
	}
}

func (widget *Widget) SetEventCallback(eventName devent.EventName, cb devent.EventCallback) {
	widget.EventMap[eventName] = cb
}

func (widget *Widget) Render(screen tcell.Screen, context lcontext.Context) {

}
