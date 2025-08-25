package dama

import (
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	"github.com/abdessamad-zgor/dama/keystroke"
	"github.com/gdamore/tcell/v2"
)

type EventManager struct {
	App 				*App
	Buffer  			string
	KeystrokeChannel 	chan devent.KeystrokeEvent
	AppEventChannel     chan devent.AppEvent
	Events 				dutils.EList[devent.DamaEvent]
}

func NewEventManager() EventManager {
	em := EventManager {
		nil,
		"",
		make(chan devent.KeystrokeEvent),
		make(chan devent.AppEvent),
		dutils.NewEList[devent.DamaEvent](),
	}
	return em
}

func (em *EventManager) RegisterEvents() {
	current := &em.App.Navigator.Current
	em.Events.Empty()
	for current != nil {
		currentWidget, ok := current.Element.(*Widget)
		for _, e := range current.Events {
			em.Events.Add(e)
		}
		current = current.Parent
	}
}

func (em *EventManager) HandleTcellEvents() {
	for {
		event := tcell.PollEvent()
		switch event.(type) {
		case tcell.EventKey:
			keystrokeEvent := devent.ToKeystrokeEvent(event)
			em.KeystrokeChannel <- keystrokeEvent
		case tcell.EventResize:
			em.App.Resize()
		}
	}
}

func (em *EventManager) EventLoop() {
	go em.HandleTcellEvents()
	for {
		select {
			case keystrokeEvent := <- em.KeystrokeChannel:
				em.Buffer = em.Buffer + keystrokeEvent.Keystroke
				em.HandleKeybindings()
			case appEvent := <- em.AppEventChannel:
				
		}
	}
}

func (em *EventManager) HandleKeybindings() {
	fullMatches := []devent.Event{}
	partialMatches := []devent.Event{}
	buffer := em.Buffer
	for _, e := range em.Events.Items() {
		// just filter on event type
		if kb, ok := e.Detail.(*devent.Keybinding); ok && kb.Matcher(em.Buffer).IsFull() {
			fullMatches = append(fullMatches, e)
		}
		if ok && kb.Matcher(em.Buffer).IsPartial() {
			partialMatches = append(partialMatches, e)
		}
	}
	dutils.Assert(fullMatches <= 1, "there should be at most 1 full match when handling keybindings")
	if len(fullMatches) == 1 && len(partialMatches) == 0 {
		kb, ok := fullMatches.Detail.(*devent.Keybinding);
		kb.Callback(e.Detail)
	}

	if len(partials) > 0 {
		time.Wait(300)
		if buffer == em.Buffer && len(fulls) == 1 {
			kb, ok := fullMatches.Detail.(*devent.Keybinding);
			kb.Callback(e.Detail)
		}
	}
}

func (em *EventManager) HandleAppEvents() {
	// later baby, not now
}
