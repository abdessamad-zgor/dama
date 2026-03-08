package dama

import (
	"fmt"
	"time"
	"sync"
	"slices"
	devent "github.com/abdessamad-zgor/dama/event"
	dkeybinding "github.com/abdessamad-zgor/dama/keybinding"
	dutils "github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/logger"
)

type EventManager struct {
	App 				App
	Wg					sync.WaitGroup
	Buffer  			string
	KeyChannel 			chan devent.KeyEvent
	AppEventChannel     chan devent.AppEventName
	Events 				dutils.List[devent.DamaEvent]
	GlobalEvents		dutils.List[devent.DamaEvent]
}

func NewEventManager(app App) *EventManager {
	var wg sync.WaitGroup

	em := &EventManager {
		app,
		wg,
		"",
		make(chan devent.KeyEvent),
		make(chan devent.AppEventName),
		dutils.NewList[devent.DamaEvent](),
		dutils.NewList[devent.DamaEvent](),
	}
	return em
}

func (em *EventManager) RegisterEvents() {
	globals := em.GlobalEvents
	current := em.App.GetNavigator().GetCurrent()
	em.Events.Empty()
	currentWidget, _ := current.element.(Widget)
	events := dutils.NewList[devent.DamaEvent]()
	// mode switching keybindings
	logger.Log(fmt.Sprintf("current widget: %+v, title: %s, Tag: %c", currentWidget, currentWidget.GetTitle(), currentWidget.GetTag()))
	modes := currentWidget.GetEventModes() 
	logger.Log(fmt.Sprintf("modes: %+v, len: %d", modes, len(modes)))
	if len(modes) > 1 {
		var keybinding devent.DamaEvent
		if currentWidget.GetMode() == devent.InsertMode {
			keybinding = devent.KeybindingToEvent("<Esc>", func (match dkeybinding.Match) {
				_ = match
				currentWidget.SetMode(devent.NormalMode)
			})
			em.Events.Add(keybinding)
		}
		if currentWidget.GetMode() == devent.NormalMode {
			if slices.Contains(modes, devent.InsertMode) {
				keybinding = devent.KeybindingToEvent("i", func (match dkeybinding.Match) {
					_ = match
					currentWidget.SetMode(devent.InsertMode)
				})
				em.Events.Add(keybinding)
			}
			if slices.Contains(modes, devent.InsertMode) {
				keybinding = devent.KeybindingToEvent("v", func (match dkeybinding.Match) {
					_ = match
					currentWidget.SetMode(devent.VisualMode)
				})
				em.Events.Add(keybinding)
			}
		}
		if currentWidget.GetMode() == devent.VisualMode {
			keybinding = devent.KeybindingToEvent("<Esc>", func (match dkeybinding.Match) {
				_ = match
				currentWidget.SetMode(devent.NormalMode)
			})
			em.Events.Add(keybinding)
		} 
	}
	// exit keybinding
	keybinding := devent.KeybindingToEvent("<C-C>", func (match dkeybinding.Match) {
		_ = match
		em.App.Exit()
		logger.Log("Exit signal sent")
	})
	em.Events.Add(keybinding)

	for _, e := range globals.Items() {
		events.Add(e)
	}
	if currentWidget != nil {
		if currentWidget.GetMode() == devent.NormalMode {
			navKeybindings := em.App.GetNavigator().GetNavigationKeybindings()
			for _, e := range navKeybindings {
				events.Add(e)
			}
		}
		for _, e := range currentWidget.GetEvents().Items() {
			events.Add(e)
		}
	}
	for _, item := range events.Items() {
		toRemove := []devent.DamaEvent{}
		for _, _item := range em.Events.Items() {
			if item.IsKeybinding() && _item.IsKeybinding() {
				if _item.Detail.Keybinding.Matcher(item.Detail.Keybinding.Pattern).IsFull() {
					break
				} else if item.Detail.Keybinding.Matcher(_item.Detail.Keybinding.Pattern).IsFull() {
					em.Events.Add(item)
					toRemove = append(toRemove, _item)
					break
				}
			}
		}
		for _, _item := range toRemove {
			em.Events.Remove(_item)
		}
		if len(toRemove) == 0 {
			em.Events.Add(item)
		}
	}
	logger.Log(fmt.Sprintf("Registred events: %+v", em.Events.Items()))
}

func (em *EventManager) HandleTcellEvents() {
	logger.Log("Starting tcell event loop.")
	for {
		event := em.App.GetScreen().PollEvent()
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
	em.App.GetNavigator().Setup()
	go em.HandleTcellEvents()
	em.RegisterEvents()
	em.App.Render(em.App.GetScreen())
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
		em.App.Render(em.App.GetScreen())
		em.RegisterEvents()
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
			if kb.Matcher(buffer).IsFull() {
				fulls = append(fulls, e)
			}
			if kb.Matcher(buffer).IsPartial() {
				partials = append(partials, e)
			}
		}
	}
	// dutils.Assert(len(fulls) <= 1, "there should be at most 1 full match when handling keybindings")
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
	}else if len(fulls) == 0 && len(partials) == 0 {
		em.Buffer = ""
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
