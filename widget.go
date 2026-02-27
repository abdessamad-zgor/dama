package dama

import (
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	dkeybinding "github.com/abdessamad-zgor/dama/keybinding"
	"github.com/gdamore/tcell/v2"
)

type DamaWidget interface {
	GetParent() *Container

	SetKeybinding(pattern string, callback devent.KeybindingCallback)
	SetModeKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback)
	SetMode(mode devent.Mode)
	GetMode() devent.Mode
	SetAppEvent(eventname devent.AppEventName, callback devent.AppEventCallback)
	DamaElement
}

type Widget struct {
	*Element
	Parent      *Container
	Events 		dutils.VList[devent.Mode, devent.DamaEvent]
	Mode		devent.Mode
}

func NewWidget() *Widget {
	widget := Widget{
		new(Element),
		nil,
		dutils.NewVList[devent.Mode, devent.DamaEvent](devent.NormalMode, devent.InsertMode, devent.VisualMode),
		devent.NormalMode,
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

//func (widget *Widget) GetModalKeybindings() []devent.DamaEvent {
//
//}

func (widget *Widget) SetModeKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback) {
	patternMatcher, err := dkeybinding.GetMatcher(pattern)
	if err != nil {
		panic(err)
	}
	keybinding := devent.DamaEvent{
		devent.DKeybinding,
		devent.EventDetail{
			&devent.Keybinding{
				pattern,
				patternMatcher,
				callback,
			},
			nil,
		},
	}

	widget.Events.AddToView(mode, keybinding)
}

func (widget *Widget) SetMode(mode devent.Mode) {
	widget.Mode = mode
	widget.Events.SwitchView(mode)
}

func (widget *Widget) GetMode() devent.Mode {
	return widget.Mode
}

func (widget *Widget) SetKeybinding(pattern string, callback devent.KeybindingCallback) {
	patternMatcher, err := dkeybinding.GetMatcher(pattern)
	if err != nil {
		panic(err)
	}
	keybinding := devent.DamaEvent{
		devent.DKeybinding,
		devent.EventDetail{
			&devent.Keybinding{
				pattern,
				patternMatcher,
				callback,
			},
			nil,
		},
	}
	widget.Events.AddToView(devent.NormalMode, keybinding)
}

func (widget *Widget) SetAppEvent(eventName devent.AppEventName, cb devent.AppEventCallback) {
	appevent := devent.DamaEvent{
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
