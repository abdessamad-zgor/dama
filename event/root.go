package event

import (
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/gdamore/tcell/v2"
)

type EventName string

const (
	InsertMode    EventName = "insert-mode"
	VisualMode    EventName = "visual-mode"
	NormalMode    EventName = "normal-mode"
	TagNavigation EventName = "tag-navigation"
	MoveCursor    EventName = "move-cursor"
	Escape        EventName = "escape"
	Quit          EventName = "quit"
	Confirm       EventName = "confirm"
	Key           EventName = "key"
	CR		 	  EventName = "carriage-return"
	Help          EventName = "help"
	Save          EventName = "save"
	Left          EventName = "left"
	Right         EventName = "right"
	Top           EventName = "top"
	Bottom        EventName = "bottom"
)

type Event struct {
	Name   EventName
	Key    tcell.Key
	TEvent tcell.Event
}

type Callback = func(context lcontext.Context, event Event)

type EventMap = map[EventName]Callback
type Keybindings = map[tcell.Key]EventName

var AppEventMap EventMap

func DefaultEventMap() EventMap {
	appEventMap := make(EventMap)

	return appEventMap
}

func DefaultKeybindings() Keybindings {
	appKeybindings := make(Keybindings)
	return appKeybindings
}
