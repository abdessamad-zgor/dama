package event

import (
	"github.com/gdamore/tcell/v2"
)

type EventName string

type EventDetail struct {
	Name   		EventName
	Key    		tcell.Key
	TcellEvent 	tcell.Event
}

type Callback = func (event EventDetail)

type EventType string

const (
	AppEvent EventType = "app-event"
	Keybinding EventType = "keybinding"
)

type KeybindingDetail struct {
	Value	string

}

type DamaEvent struct {
	Type EventType
}
