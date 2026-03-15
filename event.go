package dama

import (
	"errors"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/constants"
)

type EventType string
type Mode string
type AppEventName string

type AppEventCallback = func (payload any)
type KeybindingCallback = func (match Match)

const (
	DAppEvent EventType = "app-event"
	DKeybinding EventType = "keybinding"
)

const (
	InsertMode Mode = "insert"
	NormalMode Mode = "normal"
	VisualMode Mode = "visual"
)

type Event interface {
	Type() EventType
	ToAppEvent() (AppEvent, error)
	ToKeybinding() (Keybinding, error)
}

type AppEvent struct {
	Name	AppEventName
	Handler	AppEventCallback
}

type AppEventDispatch struct {
	Name	AppEventName
	Payload	any
}

type Keybinding struct {
	Mode		Mode
	Pattern		string
	Matcher		Matcher
	Handler     KeybindingCallback
}

type KeyEvent struct {
	Key			string
	RecievedAt	time.Time
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

func (kb Keybinding) Type() EventType {
	return DKeybinding
}

func (kb Keybinding) ToKeybinding() (Keybinding, error) {
	return kb, nil
}

func (kb Keybinding) ToAppEvent() (AppEvent, error) {
	return AppEvent{}, errors.New("Event is not AppEvent.")
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
		eventString, _ := constants.TcellKeyToString[key]
		return KeyEvent {
			eventString,
			ke.When(),
		}
	}
}

func KeybindingToEvent(mode Mode, pattern string, callback KeybindingCallback) Event {
	patternMatcher, err := GetMatcher(pattern)
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

