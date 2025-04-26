package event

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/gdamore/tcell/v2"
)

type EventName string

var EventChannel chan Event = make(chan Event)

const (
	Escape        EventName = "escape"
	Quit          EventName = "quit"
	Confirm       EventName = "confirm"
	Help          EventName = "help"
	Save          EventName = "save"
)

type Event struct {
	Name   EventName
	Key    tcell.Key
	TEvent tcell.Event
}

type KeyEvent struct {
	Key 	tcell.Key
	TEvent 	tcell.Event
}

type EventCallback = func(context lcontext.Context, event Event)
type KeybindingCallback = func(context lcontext.Context, event KeyEvent)

type EventMap = map[EventName]EventCallback
type Keybindings = map[tcell.Key]KeybindingCallback

var AppEventMap EventMap

func DefaultEventMap() EventMap {
	appEventMap := make(EventMap)

	return appEventMap
}

func DefaultKeybindings() Keybindings {
	appKeybindings := make(Keybindings)
	return appKeybindings
}
