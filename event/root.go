package event

import (
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/keybinding"
)

type EventType string

const (
	DAppEvent EventType = "app-event"
	DKeybinding EventType = "keybinding"
)

type KeybindingCallback = func (match keybinding.Match)
type AppEventCallback = func ()

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
	Handler	AppEventCallback
}

type KeyEvent struct {
	Key			string
	RecievedAt	time.Time
}

type Keybinding struct {
	Pattern		string
	Matcher		keybinding.Matcher
	Handler     KeybindingCallback
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
		eventString, _ := keybinding.TcellKeyToString[key]
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
