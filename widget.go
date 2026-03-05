package dama

import (
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	dkeybinding "github.com/abdessamad-zgor/dama/keybinding"
	"github.com/gdamore/tcell/v2"
)

type Widget interface {
	GetParent() Container

	SetKeybinding(pattern string, callback devent.KeybindingCallback)
	SetModeKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback)
	SetMode(mode devent.Mode)
	GetMode() devent.Mode
	GetEvents() dutils.VList[devent.Mode, devent.DamaEvent]
	SetAppEvent(eventname devent.AppEventName, callback devent.AppEventCallback)
	Element
}

type widget_s struct {
	*element_s
	Parent      *container_s
	Events 		dutils.VList[devent.Mode, devent.DamaEvent]
	Mode		devent.Mode
}

func NewWidget() Widget {
	wdg := widget_s {
		new(element_s),
		nil,
		dutils.NewVList[devent.Mode, devent.DamaEvent](devent.NormalMode, devent.InsertMode, devent.VisualMode),
		devent.NormalMode,
	}

	return &wdg
}

func (widget *widget_s) GetParent() Container {
	return widget.Parent
}

func (widget *widget_s) GetBox() Box {
	box := widget.GetBox()
	box.Element = widget
	return box
}

//func (widget *widget_s) GetModalKeybindings() []devent.DamaEvent {
//
//}

func (widget *widget_s) SetModeKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback) {
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

func (widget *widget_s) SetMode(mode devent.Mode) {
	widget.Mode = mode
	widget.Events.SwitchView(mode)
}

func (widget *widget_s) GetMode() devent.Mode {
	return widget.Mode
}

func (widget *widget_s) SetKeybinding(pattern string, callback devent.KeybindingCallback) {
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

func (widget *widget_s) SetAppEvent(eventName devent.AppEventName, cb devent.AppEventCallback) {
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

func (widget *widget_s) GetEvents() dutils.VList[devent.Mode, devent.DamaEvent] {
	return widget.Events
}

func (widget *widget_s) Render(screen tcell.Screen) {

}
