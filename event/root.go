package event

import (
	"github.com/gdamore/tcell/v2"
    lcontext "github.com/abdessamad-zgor/dama/context"
)

type EventName string

const (
	InsertMode EventName = "insert-mode"
	VisualMode EventName = "visual-mode"
	NormalMode EventName = "normal-mode"
	Escape     EventName = "escape"
	Quit       EventName = "quit"
	Confirm    EventName = "confirm"
	Key        EventName = "key"
	Help       EventName = "help"
	Save       EventName = "save"
	Left       EventName = "left"
	Right      EventName = "right"
	Top        EventName = "top"
	Bottom     EventName = "bottom"
)

type Event struct {
	Name EventName
	Key  tcell.Key
}

type Callback = func(context lcontext.Context, event Event)

type EventMap = map[EventName]Callback
type Keybindings = map[tcell.Key]EventName

var AppEventMap EventMap

func SetDefaultEventMap() EventMap {
	appEventMap := make(EventMap)
	appEventMap[Key] = func(appcontext lcontext.Context, event Event) {
        // send rune to widget
        // key := event.Key()
	}

    appEventMap[Left] = func(appcontext lcontext.Context, event Event) {

    }
        
    return appEventMap
}
