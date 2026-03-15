package dama

import (
	"fmt"
	"time"
	"sync"
	_ "slices"
	dutils "github.com/samazee/dama/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/samazee/dama/logger"
)

type EventManager struct {
	App 				App
	Wg					sync.WaitGroup
	Buffer  			string
	KeyChannel 			chan KeyEvent
	AppEventChannel     chan AppEventDispatch
	Events 				dutils.List[Event]
	GlobalEvents		dutils.List[Event]
}

func NewEventManager(app App) *EventManager {
	var wg sync.WaitGroup

	em := &EventManager {
		app,
		wg,
		"",
		make(chan KeyEvent, 1),
		make(chan AppEventDispatch, 1),
		dutils.NewList[Event](),
		dutils.NewList[Event](),
	}
	return em
}

func (em *EventManager) RegisterEvents() {
	globals := em.GlobalEvents
	current := em.App.GetNavigator().GetCurrent()
	em.Events.Empty()
	currentWidget, _ := current.element.(Widget)
	events := dutils.NewList[Event]()
	// mode switching keybindings
	logger.Log(fmt.Sprintf("current widget: %+v", currentWidget))
	// exit keybinding
	keybinding := KeybindingToEvent(NormalMode, "<C-C>", func (match Match) {
		_ = match
		em.App.Exit()
		logger.Log("Exit signal sent")
	})
	em.Events.Add(keybinding)

	for _, e := range globals.Items() {
		events.Add(e)
	}
	if currentWidget != nil {
		if currentWidget.GetMode() == NormalMode {
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
			if IsKeybinding(item) && IsKeybinding(_item) {
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
		elements := em.App.GetElements()
		for _, element := range elements {
			widget, ok := element.(Widget)
			if ok {
				if currentWidget.GetTag() != widget.GetTag() && currentWidget.GetTitle() != widget.GetTitle() {
					for _, event := range widget.GetModeEvents() {
						if IsAppEvent(event) {
							appevent, _ := event.ToAppEvent()
							em.Events.Add(appevent)
						}
					}
				}
			}
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
			keyEvent := ToKeyEvent(event)
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
	fulls := []Keybinding{}
	partials := []Keybinding{}
	buffer := em.Buffer
	for _, e := range em.Events.Items() {
		if IsKeybinding(e) {
			kb, _ := e.ToKeybinding()
			if kb.Matcher(buffer).IsFull() {
				fulls = append(fulls, kb)
			}
			if kb.Matcher(buffer).IsPartial() {
				partials = append(partials, kb)
			}
		}
	}
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

func (em *EventManager) HandleAppEvent(eventDispatch AppEventDispatch) {
	defer em.Wg.Done()
	appevents := []AppEvent{}
	for _, event := range em.Events.Items() {
		if IsAppEvent(event) {
			appevent, _ := event.ToAppEvent()
			if appevent.Name == eventDispatch.Name {
				appevents = append(appevents, appevent)
			}
		}
	}
	for _, appevent := range appevents {
		appevent.Handler(eventDispatch.Payload)
	}
}
