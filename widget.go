package dama

import (
	"fmt"
	"maps"
	"slices"
	dutils "github.com/samazee/dama/utils"
	"github.com/samazee/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Widget interface {
	GetParent() Container

	SetKeybinding(mode Mode, pattern string, callback KeybindingCallback)
	SetMode(mode Mode)
	GetMode() Mode
	GetModeEvents() []Event
	GetWidgetModes() []Mode
	SetAppEvent(eventname AppEventName, callback AppEventCallback)
	GetTraits() []Trait
	Element
}

type widget_s struct {
	Element
	Parent      Container
	Events 		dutils.List[Event]
	Mode		Mode
	Traits		[]Trait
}

func NewWidget(traits ...Trait) Widget {
	wdg := &widget_s {
		NewElement(),
		nil,
		dutils.NewList[Event](),
		NormalMode,
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

//func (widget *widget_s) GetModalKeybindings() []DamaEvent {
//
//}

func (widget *widget_s) SetKeybinding(mode Mode, pattern string, callback KeybindingCallback) {
	if mode == InsertMode {
		if !slices.Contains(widget.GetWidgetModes(), InsertMode) {
			modeSwitchingKb := KeybindingToEvent(mode, "<Esc>", func (match Match) {
				_ = match
				widget.SetMode(NormalMode)
			})
			widget.Events.Add(modeSwitchingKb)
			modeSwitchingKb = KeybindingToEvent(NormalMode, "i", func (match Match) {
				_ = match
				widget.SetMode(InsertMode)
			})
			widget.Events.Add(modeSwitchingKb)
		}
	} else if mode == VisualMode {
		if !slices.Contains(widget.GetWidgetModes(), VisualMode) {
			modeSwitchingKb := KeybindingToEvent(mode, "<Esc>", func (match Match) {
				_ = match
				widget.SetMode(NormalMode)
			})
			widget.Events.Add(modeSwitchingKb)
			modeSwitchingKb = KeybindingToEvent(NormalMode, "v", func (match Match) {
				_ = match
				widget.SetMode(VisualMode)
			})
			widget.Events.Add(modeSwitchingKb)
		}

	}
	keybinding := KeybindingToEvent(mode, pattern, callback)
	widget.Events.Add(keybinding)
	logger.Log(fmt.Sprintf("widget events: %+v", widget.Events))
}

func (widget *widget_s) SetMode(mode Mode) {
	widget.Mode = mode
	if mode == InsertMode {
		widget.BorderColor(tcell.ColorTeal)
	} else if mode == NormalMode {
		widget.BorderColor(tcell.ColorLime)
	} else if mode == VisualMode {
		widget.BorderColor(tcell.ColorYellow)
	}
}

func (widget *widget_s) GetMode() Mode {
	return widget.Mode
}

func (widget *widget_s) SetAppEvent(eventName AppEventName, cb AppEventCallback) {
	appevent := AppEventToEvent(eventName, cb)
	widget.Events.Add(appevent)
}

func (widget *widget_s) GetModeEvents() []Event {
	modeEvents := []Event{}
	for _, event := range widget.Events.Items() {
		if IsKeybinding(event) {
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

func (widget *widget_s) GetWidgetModes() []Mode {
	modes := make(map[Mode]int)
	for _, event := range widget.Events.Items() {
		if IsKeybinding(event) {
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
