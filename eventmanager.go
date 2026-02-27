package dama

import (
	"fmt"
	"time"
	"sync"
	devent "github.com/abdessamad-zgor/dama/event"
	dutils "github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/logger"
)

type EventManager struct {
	App 				*App
	Wg					sync.WaitGroup
	Buffer  			string
	KeyChannel 			chan devent.KeyEvent
	AppEventChannel     chan devent.AppEventName
	Events 				dutils.EList[devent.DamaEvent]
	GlobalEvents		dutils.List[devent.DamaEvent]
}

func NewEventManager(app *App) *EventManager {
	var wg sync.WaitGroup
	var excludeFn dutils.ExcludeFn[devent.DamaEvent] = func (itemList dutils.List[devent.DamaEvent], item devent.DamaEvent) int {
		insertable := true
		toRemove := []devent.DamaEvent{}
		for _, _item := range itemList.Items() {
			if item.IsKeybinding() && _item.IsKeybinding() {
				// more conditions should be introduced to allow for the keybinding with the narrower Pattern to be accepted
				// this whole exclusion principle needs to be rethinked
				if _item.Detail.Keybinding.Matcher(item.Detail.Keybinding.Pattern).IsFull() {
					insertable = false
				} else if item.Detail.Keybinding.Matcher(_item.Detail.Keybinding.Pattern).IsFull() {
					insertable = true
					toRemove = append(toRemove, _item)
				}
				if !insertable {
					return -1
				}
			}
		}
		for _, _item := range toRemove {
			itemList.Remove(_item)
		}
		return itemList.Length()
	}
	em := &EventManager {
		app,
		wg,
		"",
		make(chan devent.KeyEvent),
		make(chan devent.AppEventName),
		dutils.NewEList[devent.DamaEvent](excludeFn),
		dutils.NewList[devent.DamaEvent](),
	}
	return em
}

func (em *EventManager) RegisterEvents() {
	globals := em.GlobalEvents
	current := em.App.Navigator.current
	em.Events.Empty()
	currentWidget, _ := current.element.(*Widget)
	//modalKeybindings := currentWidget.GetModalKeybindings()

	for _, e := range globals.Items() {
		em.Events.Add(e)
	}
	if currentWidget.GetMode() == devent.NormalMode {
		navKeybindings := em.App.Navigator.GetNavigationKeybindings()
		for _, e := range navKeybindings {
			em.Events.Add(e)
		}
	}
	for _, e := range currentWidget.Events.Items() {
		em.Events.Add(e)
	}
	logger.Log(fmt.Sprintf("Registred events: %+v", em.Events.Items()))
}

func (em *EventManager) HandleTcellEvents() {
	logger.Log("Starting tcell event loop.")
	for {
		event := em.App.Screen.PollEvent()
		logger.Log("Recieved tcell event", fmt.Sprintf("%+v", event))
		switch event.(type) {
		case *tcell.EventKey:
			keyEvent := devent.ToKeyEvent(event)
			em.KeyChannel <- keyEvent
		case *tcell.EventResize:
			em.App.Resize()
		}
	}
}


func (em *EventManager) StartEventLoop() {
	logger.Log("Starting App Event Loop")
	em.App.Navigator.Setup()
	go em.HandleTcellEvents()
	em.RegisterEvents()
	for {
		select {
			case keyEvent := <- em.KeyChannel:
				logger.Log("Key Sent: ", keyEvent)
				em.Buffer = em.Buffer + keyEvent.Key
				em.Wg.Add(1)
				go em.HandleKeybindings()
			case appEvent := <- em.AppEventChannel:
				logger.Log("App event sent: ", appEvent)
				em.Wg.Add(1)
				go em.HandleAppEvent(appEvent)	
		}
	}
}

func (em *EventManager) HandleKeybindings() {
	defer em.Wg.Done()
	fulls := []devent.DamaEvent{}
	partials := []devent.DamaEvent{}
	buffer := em.Buffer
	for _, e := range em.Events.Items() {
		if e.IsKeybinding() {
			kb := e.Detail.Keybinding
			if kb.Matcher(em.Buffer).IsFull() {
				fulls = append(fulls, e)
			}
			if kb.Matcher(em.Buffer).IsPartial() {
				partials = append(partials, e)
			}
		}
	}
	//dutils.Assert(len(fulls) <= 1, "there should be at most 1 full match when handling keybindings")
	// if there are no other keybindings that could match the 
	logger.Log("full keybinding matchs: ", fulls)
	logger.Log("partial keybinding matchs: ", partials)
	if len(fulls) == 1 {
		if len(partials) == 0 {
			e := fulls[0]
			logger.Log("Found one full keybinding: ", e)
			kb := e.Detail.Keybinding
			kb.Handler(kb.Matcher(em.Buffer))
			em.Buffer = ""
		}

		time.Sleep(300 * time.Millisecond)
		if buffer == em.Buffer {
			e := fulls[0]
			logger.Log("Found one full keybinding after wait: ", e)
			kb := e.Detail.Keybinding
			kb.Handler(kb.Matcher(em.Buffer))
			em.Buffer = ""
		}
	}
}

func (em *EventManager) HandleAppEvent(eventName devent.AppEventName) {
	defer em.Wg.Done()
	events := []devent.DamaEvent{}
	for _, event := range em.Events.Items() {
		if event.IsAppEvent() && event.Detail.AppEvent.Name == eventName {
			events = append(events, event)
		}
	}
	for _, appevent := range events {
		appevent.Detail.AppEvent.Handler()
	}
}
