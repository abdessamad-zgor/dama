package dama

import (
	devent "github.com/abdessamaad-zgor/dama/event"
	dutils "github.com/abdessamaad-zgor/dama/utils"
	"github.com/abdessamaad-zgor/dama/keystroke"
	"github.com/gdamore/tcell/v2"
)

type EventManager struct {
	App 				*App
	Buffer  			string
	KeystrokeChannel 	chan devent.KeystrokeEvent
	AppEventChannel     chan devent.AppEvent
	Events 				dutils.ExclusionList[devent.Event]
}

func (em *EventManager) RegisterEvents() {

}

func (em *EventManager) HandleTcellEvents() {
	for {
		event := tcell.PollEvent()
		switch (event.type) {
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
			case appEvent := <- em.AppEventChannel:
		}
	}
}

func (em *EventManager) HandleKeybindings() {

}
