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
	SetAppEvent(eventname devent.EventName, callback devent.Callback)
	DamaElement
}

type Widget struct {
	*Element
	Parent      *Container
	Events 		dutils.ViewList[string,devent.DamaEvent]
}

func NewWidget() *Widget {
	widget := Widget{
		new(Element),
		nil,
		dutils.NewViewList[string, devent.DamaEvent](),
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

func (widget *Widget) SetKeybinding(pattern string, cb devent.Callback) {
}

func (widget *Widget) SetAppEvent(eventName devent.EventName, cb devent.Callback) {
}

func (widget *Widget) Render(screen tcell.Screen) {

}
