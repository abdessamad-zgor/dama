package event

type EventType string

const (
	DAppEvent EventType = "app-event"
	DKeybinding EventType = "keybinding"
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
