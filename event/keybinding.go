package event

import (
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/keystroke"
)

type KeyEvent struct {
	Keystroke  string
	RecievedAt time.Time
}

func ToKeyEvent(event tcell.Event) KeyEvent {
	ke, _ := event.(*tcell.EventKey)
	key := ke.Key()
	switch key {
	case tcell.KeyRune:
		char := ke.Rune()
		return KeyEvent{
			string(char),
			ke.When(),
		}
	default:
		eventString, _ := keystroke.TcellKeyToString[key]
		return KeyEvent{
			eventString,
			ke.When(),
		}
	}
}

type Keybinding struct {
	Pattern		string
	Matcher		keystroke.Matcher
	Handler     Callback
}
