package dama

import (
	"fmt"
	"maps"
	"slices"
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	_ "github.com/abdessamad-zgor/dama/keybinding"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Widget interface {
	GetParent() Container

	SetKeybinding(pattern string, callback devent.KeybindingCallback)
	SetModeKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback)
	SetMode(mode devent.Mode)
	GetMode() devent.Mode
	GetEvents() dutils.List[devent.DamaEvent]
	GetEventModes() []devent.Mode
	SetAppEvent(eventname devent.AppEventName, callback devent.AppEventCallback)
	Element
}

type widget_s struct {
	Element
	Parent      Container
	Events 		map[devent.Mode]dutils.List[devent.DamaEvent]
	Mode		devent.Mode
}

func NewWidget() Widget {
	wdg := widget_s {
		NewElement(),
		nil,
		make(map[devent.Mode]dutils.List[devent.DamaEvent]),
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
	keybinding := devent.KeybindingToEvent(pattern, callback)
	if _, ok := widget.Events[mode]; !ok {
		widget.Events[mode] = dutils.NewList[devent.DamaEvent]()
	}
	widget.Events[mode].Add(keybinding)
	logger.Log(fmt.Sprintf("widget events: %+v", widget.Events))
}

func (widget *widget_s) SetMode(mode devent.Mode) {
	widget.Mode = mode
	if mode == devent.InsertMode {
		widget.BorderColor(tcell.ColorTeal)
	} else if mode == devent.NormalMode {
		widget.BorderColor(tcell.ColorLime)
	} else if mode == devent.VisualMode {
		widget.BorderColor(tcell.ColorYellow)
	}
}

func (widget *widget_s) GetMode() devent.Mode {
	return widget.Mode
}

func (widget *widget_s) SetKeybinding(pattern string, callback devent.KeybindingCallback) {
	keybinding := devent.KeybindingToEvent(pattern, callback)
	if _, ok := widget.Events[devent.NormalMode]; !ok {
		widget.Events[devent.NormalMode] = dutils.NewList[devent.DamaEvent]()
	}
	widget.Events[devent.NormalMode].Add(keybinding)
}

func (widget *widget_s) SetAppEvent(eventName devent.AppEventName, cb devent.AppEventCallback) {
	appevent := devent.AppEventToEvent(eventName, cb)
	widget.Events[widget.Mode].Add(appevent)
}

func (widget *widget_s) GetEvents() dutils.List[devent.DamaEvent] {
	return widget.Events[widget.Mode]
}

func (widget *widget_s) GetEventModes() []devent.Mode {
	return slices.Collect(maps.Keys(widget.Events))
}

//func (widget *widget_s) Render(screen tcell.Screen) {
//	widget.Render()
//}
