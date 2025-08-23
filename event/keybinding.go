package event

import (
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/keystroke"
)

type KeystrokeEvent struct {
	Keystroke  string
	RecievedAt time.Time
}

func ToKeytrokeEvent(event tcell.EventKey) KeystrokeEvent {
	key := event.Key()
	switch key {
	case tcell.KeyRune:
		char := event.Rune()
		return KeystrokeEvent{
			string(char),
			event.When(),
		}
	default:
		eventString, _ := keystroke.TcellKeyToString[key]
		return KeystrokeEvent{
			eventString,
			event.When(),
		}
	}
}

type Keybinding struct {
	Pattern		string
	Matcher		keystroke.Matcher
	Handler     Callback
}
