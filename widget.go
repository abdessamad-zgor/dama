package dama

import (
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
)

type DamaWidget interface {
	GetParent() *Container

	SetKeybinding(pattern string, callback devent.Callback)
	SetAppEvent(eventname devent.EventName, callback devent.Callback)
	DamaElement
}

type Widget struct {
	*Element
	Parent      *Container
	Events 		dutils.VList[devent.DamaEvent]
}

func NewWidget() *Widget {
	widget := Widget{
		new(Element),
		nil,
		dutils.NewVList[devent.DamaEvent](),
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
	patternMatcher, err := keystroke.GetMatcher(pattern)
	if err != nil {
		panic(err)
	}
	keybinding := devent.Event{
		devent.DKeybinding,
		devent.EventDetail{
			devent.Keybinding{
				pattern,
				patternMatcher,
				cb,
			},
		},
	}

	widget.Events.Add(keybinding)
}

func (widget *Widget) SetAppEvent(eventName devent.EventName, cb devent.Callback) {
	appevent := devent.Event{
		devent.DAppEvent,
		devent.EventDetail{
			nil,
			&devent.AppEvent{
				eventName,
				nil,
				cb,
			},
		},
	}

	widget.Events.Add(appevent)
}

func (widget *Widget) Render(screen tcell.Screen) {

}
