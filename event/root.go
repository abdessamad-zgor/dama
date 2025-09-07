package event

import (
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/keystroke"
)

type EventType string

const (
	DAppEvent EventType = "app-event"
	DKeybinding EventType = "keybinding"
)

type Callback = func (event EventDetail)

type DamaEvent struct {
	Type	EventType
	Detail	EventDetail
}

type EventDetail struct {
	Keybinding	*Keybinding
	AppEvent	*AppEvent
}

type AppEventName string

type AppEvent struct {
	Name	AppEventName
	Payload any
	Handler	Callback
}

type KeyEvent struct {
	Keystroke  string
	RecievedAt time.Time
}

type Keybinding struct {
	Pattern		string
	Matcher		keystroke.Matcher
	Handler     Callback
}

func ToKeyEvent(event tcell.Event) KeyEvent {
	ke, _ := event.(*tcell.EventKey)
	key := ke.Key()
	switch key {
	case tcell.KeyRune:
		char := ke.Rune()
		return KeyEvent {
			string(char),
			ke.When(),
		}
	default:
		eventString, _ := keystroke.TcellKeyToString[key]
		return KeyEvent {
			eventString,
			ke.When(),
		}
	}
}

func (event DamaEvent) IsKeybinding() bool {
	return event.Type == DKeybinding
}

func(event DamaEvent) IsAppEvent() bool {
	return event.Type == DAppEvent
}
