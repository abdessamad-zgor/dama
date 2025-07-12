package event

import (
	"github.com/gdamore/tcell/v2"
)
type EventType string

:q
:q
const (
	AppEvent EventType = "app-event"
	Keybinding EventType = "keybinding"
)

type Callback = func (event EventDetail)

type Event struct {
	Type	EventType
	Detail	EventDetail
}

type EventDetail struct {
	Keybinding	*Keybinding
	AppEvent	*AppEvent
}
