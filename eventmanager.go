package dama

import (
	"fmt"
	"time"
	"sync"
	_ "slices"
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
	Events 				dutils.List[devent.Event]
	GlobalEvents		dutils.List[devent.Event]
}

func NewEventManager(app App) *EventManager {
	var wg sync.WaitGroup

	em := &EventManager {
		app,
		wg,
		"",
		make(chan devent.KeyEvent),
		make(chan devent.AppEventName),
		dutils.NewList[devent.Event](),
		dutils.NewList[devent.Event](),
	}
	return em
}

func (em *EventManager) RegisterEvents() {
	globals := em.GlobalEvents
	current := em.App.GetNavigator().GetCurrent()
	em.Events.Empty()
	currentWidget, _ := current.element.(Widget)
	events := dutils.NewList[devent.Event]()
	// mode switching keybindings
	logger.Log(fmt.Sprintf("current widget: %+v", currentWidget))
	// exit keybinding
	keybinding := devent.KeybindingToEvent(devent.NormalMode, "<C-C>", func (match dkeybinding.Match) {
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
		for _, e := range currentWidget.GetModeEvents() {
			events.Add(e)
		}
	}
	for _, item := range events.Items() {
		toRemove := []int{}
		for i, _item := range em.Events.Items() {
			if devent.IsKeybinding(item) && devent.IsKeybinding(_item) {
				_itemKb, _ := _item.ToKeybinding()
				itemKb, _ := item.ToKeybinding()
				if _itemKb.Matcher(itemKb.Pattern).IsFull() {
					break
				} else if itemKb.Matcher(_itemKb.Pattern).IsFull() {
					em.Events.Add(item)
					toRemove = append(toRemove, i)
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
	fulls := []devent.Keybinding{}
	partials := []devent.Keybinding{}
	buffer := em.Buffer
	for _, e := range em.Events.Items() {
		if devent.IsKeybinding(e) {
			kb, _ := e.ToKeybinding()
			if kb.Matcher(buffer).IsFull() {
				fulls = append(fulls, kb)
			}
			if kb.Matcher(buffer).IsPartial() {
				partials = append(partials, kb)
			}
		}
	}
	// dutils.Assert(len(fulls) <= 1, "there should be at most 1 full match when handling keybindings")
	// if there are no other keybindings that could match the 
	logger.Log("full keybinding matchs: ", fulls)
	logger.Log("partial keybinding matchs: ", partials)
	if len(fulls) == 1 {
		if len(partials) == 0 {
			kb := fulls[0]
			logger.Log("Found one full keybinding: ", kb)
			kb.Handler(kb.Matcher(em.Buffer))
			em.Buffer = ""
			return
		}

		time.Sleep(300 * time.Millisecond)
		if buffer == em.Buffer {
			kb := fulls[0]
			logger.Log("Found one full keybinding after wait: ", kb)
			kb.Handler(kb.Matcher(em.Buffer))
			em.Buffer = ""
		}
	}else if len(fulls) == 0 && len(partials) == 0 {
		em.Buffer = ""
	}
}

func (em *EventManager) HandleAppEvent(eventName devent.AppEventName) {
	defer em.Wg.Done()
	appevents := []devent.AppEvent{}
	for _, event := range em.Events.Items() {
		if devent.IsAppEvent(event) {
			appevent, _ := event.ToAppEvent()
			appevents = append(appevents, appevent)
		}
	}
	for _, appevent := range appevents {
		appevent.Handler()
	}
}
