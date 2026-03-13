package event

import (
	"errors"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/keybinding"
)

type EventType string

const (
	DAppEvent EventType = "app-event"
	DKeybinding EventType = "keybinding"
)

type Mode string

const (
	InsertMode Mode = "insert"
	NormalMode Mode = "normal"
	VisualMode Mode = "visual"
)

type KeybindingCallback = func (match keybinding.Match)
type AppEventCallback = func ()

type Event interface {
	Type() EventType
	ToAppEvent() (AppEvent, error)
	ToKeybinding() (Keybinding, error)
}

type AppEventName string

type AppEvent struct {
	Name	AppEventName
	Payload any
	Handler	AppEventCallback
}

func (ae AppEvent) Type() EventType {
	return DAppEvent
}

func (ae AppEvent) ToKeybinding() (Keybinding, error) {
	return Keybinding{}, errors.New("Event is not Keybinding.")
}

func (ae AppEvent) ToAppEvent() (AppEvent, error) {
	return ae, nil
}

type Keybinding struct {
	Mode		Mode
	Pattern		string
	Matcher		keybinding.Matcher
	Handler     KeybindingCallback
}

func (kb Keybinding) Type() EventType {
	return DKeybinding
}

func (kb Keybinding) ToKeybinding() (Keybinding, error) {
	return kb, nil
}

func (kb Keybinding) ToAppEvent() (AppEvent, error) {
	return AppEvent{}, errors.New("Event is not AppEvent.")
}

type KeyEvent struct {
	Key			string
	RecievedAt	time.Time
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

func KeybindingToEvent(mode Mode, pattern string, callback KeybindingCallback) Event {
	patternMatcher, err := keybinding.GetMatcher(pattern)
	if err != nil {
		panic(err)
	}
	kb := Keybinding {
		mode,
		pattern,
		patternMatcher,
		callback,
	}
	return kb
}

func AppEventToEvent(eventName AppEventName, callback AppEventCallback) Event {
	appevent := AppEvent{
		eventName,
		nil,
		callback,
	}
	return appevent
}

func IsKeybinding(event Event) bool {
	return event.Type() == DKeybinding
}

func IsAppEvent(event Event) bool {
	return event.Type() == DAppEvent
}
