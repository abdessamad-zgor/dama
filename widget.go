package dama

import (
	"fmt"
	"maps"
	"slices"
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	dkeybinding "github.com/abdessamad-zgor/dama/keybinding"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Widget interface {
	GetParent() Container

	SetKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback)
	SetMode(mode devent.Mode)
	GetMode() devent.Mode
	GetModeEvents() []devent.Event
	GetWidgetModes() []devent.Mode
	SetAppEvent(eventname devent.AppEventName, callback devent.AppEventCallback)
	GetTraits() []Trait
	Element
}

type widget_s struct {
	Element
	Parent      Container
	Events 		dutils.List[devent.Event]
	Mode		devent.Mode
	Traits		[]Trait
}

func NewWidget(traits ...Trait) Widget {
	wdg := &widget_s {
		NewElement(),
		nil,
		dutils.NewList[devent.Event](),
		devent.NormalMode,
		traits,
	}
	for _, trait := range traits {
		traitKeybindings := trait.GetTraitKeybindings()
		for _, _keybinding := range traitKeybindings {
			keybinding, _ := _keybinding.ToKeybinding()
			wdg.SetKeybinding(keybinding.Mode, keybinding.Pattern, keybinding.Handler)
		}
	}
	return wdg
}

func (widget *widget_s) GetParent() Container {
	return widget.Parent
}

func (widget *widget_s) GetBox() Box {
	box := widget.Element.GetBox()
	box.Element = widget
	return box
}

//func (widget *widget_s) GetModalKeybindings() []devent.DamaEvent {
//
//}

func (widget *widget_s) SetKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback) {
	if mode == devent.InsertMode {
		if !slices.Contains(widget.GetWidgetModes(), devent.InsertMode) {
			modeSwitchingKb := devent.KeybindingToEvent(mode, "<Esc>", func (match dkeybinding.Match) {
				_ = match
				widget.SetMode(devent.NormalMode)
			})
			widget.Events.Add(modeSwitchingKb)
			modeSwitchingKb = devent.KeybindingToEvent(devent.NormalMode, "i", func (match dkeybinding.Match) {
				_ = match
				widget.SetMode(devent.InsertMode)
			})
			widget.Events.Add(modeSwitchingKb)
		}
	} else if mode == devent.VisualMode {
		if !slices.Contains(widget.GetWidgetModes(), devent.VisualMode) {
			modeSwitchingKb := devent.KeybindingToEvent(mode, "<Esc>", func (match dkeybinding.Match) {
				_ = match
				widget.SetMode(devent.NormalMode)
			})
			widget.Events.Add(modeSwitchingKb)
			modeSwitchingKb = devent.KeybindingToEvent(devent.NormalMode, "v", func (match dkeybinding.Match) {
				_ = match
				widget.SetMode(devent.VisualMode)
			})
			widget.Events.Add(modeSwitchingKb)
		}

	}
	keybinding := devent.KeybindingToEvent(mode, pattern, callback)
	widget.Events.Add(keybinding)
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

func (widget *widget_s) SetAppEvent(eventName devent.AppEventName, cb devent.AppEventCallback) {
	appevent := devent.AppEventToEvent(eventName, cb)
	widget.Events.Add(appevent)
}

func (widget *widget_s) GetModeEvents() []devent.Event {
	modeEvents := []devent.Event{}
	for _, event := range widget.Events.Items() {
		if devent.IsKeybinding(event) {
			kb, _ := event.ToKeybinding()
			if kb.Mode == widget.Mode {
				modeEvents = append(modeEvents, kb)
			}
		} else {
			modeEvents = append(modeEvents, event)
		}
	}
	return modeEvents
}

func (widget *widget_s) GetWidgetModes() []devent.Mode {
	modes := make(map[devent.Mode]int)
	for _, event := range widget.Events.Items() {
		if devent.IsKeybinding(event) {
			kb, _ := event.ToKeybinding()
			modes[kb.Mode] = 1
		}
	}
	return slices.Collect(maps.Keys(modes))
}

func (widget *widget_s) Render(screen tcell.Screen) {
	logger.Log("widget.Render() called")
	widget.Element.Render(screen)
	for _, trait := range widget.Traits {
		trait.Render(widget, screen)
	}
}

func (widget *widget_s) GetTraits() []Trait {
	return widget.Traits
}
