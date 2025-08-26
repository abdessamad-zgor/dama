package dama

import (
	"time"
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
)

type EventManager struct {
	App 				*App
	Buffer  			string
	KeyChannel 			chan devent.KeyEvent
	AppEventChannel     chan devent.AppEvent
	Events 				dutils.EList[devent.DamaEvent]
}

func NewEventManager() EventManager {
	var excludeFn dutils.ExcludeFn[devent.DamaEvent]= func (itemList dutils.List[devent.DamaEvent], item devent.DamaEvent) int {
		exists := false
		for _, _item := range itemList.Items() {
			if item == _item {
				exists = true
				break
			}
		}
		if exists {
			return -1
		}
		return itemList.Length()
	}
	em := EventManager {
		nil,
		"",
		make(chan devent.KeyEvent),
		make(chan devent.AppEvent),
		dutils.NewEList[devent.DamaEvent](excludeFn),
	}
	return em
}

func (em *EventManager) RegisterEvents() {
	current := &em.App.Navigator.Current
	em.Events.Empty()
	for current != nil {
		currentWidget, _ := current.Element.(*Widget)
		for _, e := range currentWidget.Events.Items() {
			em.Events.Add(e)
		}
		current = current.Parent
	}
}

func (em *EventManager) HandleTcellEvents() {
	for {
		event := em.App.Screen.PollEvent()
		switch event.(type) {
		case *tcell.EventKey:
			keyEvent := devent.ToKeyEvent(event)
			em.KeyChannel <- keyEvent
		case *tcell.EventResize:
			em.App.Resize()
		}
	}
}

func (em *EventManager) EventLoop() {
	go em.HandleTcellEvents()
	for {
		select {
			case keyEvent := <- em.KeyChannel:
				em.Buffer = em.Buffer + keyEvent.Keystroke
				em.HandleKeybindings()
			case appEvent := <- em.AppEventChannel:
				em.HandleAppEvent(appEvent)	
		}
	}
}

func (em *EventManager) HandleKeybindings() {
	fulls := []devent.DamaEvent{}
	partials := []devent.DamaEvent{}
	buffer := em.Buffer
	for _, e := range em.Events.Items() {
		// this could panic
		kb := e.Detail.Keybinding
		if kb != nil && kb.Matcher(em.Buffer).IsFull() {
			fulls = append(fulls, e)
		}
		if kb != nil && kb.Matcher(em.Buffer).IsPartial() {
			partials = append(partials, e)
		}
	}
	dutils.Assert(len(fulls) <= 1, "there should be at most 1 full match when handling keybindings")
	if len(fulls) == 1 && len(partials) == 0 {
		e := fulls[0]
		kb := e.Detail.Keybinding
		kb.Handler(e.Detail)
	}

	if len(partials) > 0 {
		time.Sleep(300 * time.Millisecond)
		if buffer == em.Buffer && len(fulls) == 1 {
			e := fulls[0]
			kb := e.Detail.Keybinding
			kb.Handler(e.Detail)
		}
	}
}

func (em *EventManager) HandleAppEvent(event devent.AppEvent) {
	// later baby, not now
}
