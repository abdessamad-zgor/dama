package dama

import (
	"fmt"
	"testing"

	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	_ "github.com/abdessamad-zgor/dama/utils"
	"github.com/gdamore/tcell/v2"
)

type DamaApp interface {
	DamaContainer
    DamaWidget
	Start() 
	StartEventLoop()
}

type App struct {
	EventChannel chan event.Event
	Screen       tcell.Screen
    State        *WidgetState
	*Container
	event.EventMap
	event.Keybindings
	lcontext.Context
}

func initAppScreen() (tcell.Screen, error) {
	isTesting := testing.Testing()
	var screen tcell.Screen
	if !isTesting {
		new_screen, err := tcell.NewScreen()
		if err != nil {
			return nil, err
		}
        screen = new_screen
	} else {
		screen = tcell.NewSimulationScreen("UTF-8")
	}
    return screen, nil
}

func NewApp() (*App, error) {
	app := &App{
        CreateContainer(),
        
    }
    screen, err := initAppScreen()
    if err != nil {
        return nil, err
    }
	if err = screen.Init(); err != nil {
		return nil, err
	}
    app.Screen = screen
	width, height := app.Screen.Size()
	app.EventChannel = make(chan event.Event)
	app.X = 0
	app.Y = 0
	app.Width = uint(width)
	app.Height = uint(height)
	return app, nil
}

func (app *App) Start() {
	app.Screen.Clear()
	app.Screen.SetStyle(tcell.StyleDefault)
	app.Render(app.Screen, app.Context)
	app.Screen.Show()
	go app.StartKeyEventMapper()
	go app.StartEventLoop()
}

func (app *App) StartKeyEventMapper() {
	for {
		ev := app.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			key := ev.Key()
			kevent, ok := app.Keybindings[key]
			if ok {
				app.EventChannel <- event.Event{kevent, key}
			}
		case *tcell.EventResize:
			app.Screen.Sync()
		}
	}
}

func (app *App) StartEventLoop() {
	for {
		select {
		case event := <-app.EventChannel:
			callback, ok := app.EventMap[event.Name]
			if ok {
				callback(app.Context, event)
			}
		case dispatchEvent, _ := <-lcontext.DispatchContextChannel:
			switch dispatchEvent.Event {
			case lcontext.HighlightWidget:
			case lcontext.QueueRender:
				value, ok := app.Context.GetValue(lcontext.RenderQueue)
				payload := dispatchEvent.Payload
				if ok {
					queue, qOk := value.([]func())
					rendefFn, fOk := payload.(func())

					if qOk && fOk {
						queue = append(queue, rendefFn)
						app.Context.SetValue(lcontext.RenderQueue, queue)
					} else {
						panic(fmt.Sprintf("Invalid context value '%s' : %v or invalid cast %v.", lcontext.RenderQueue, value, payload))
					}
				} else {
					app.Context.SetValue(lcontext.RenderQueue, [](func()){})
				}
			}
		}
	}
}

func (app *App) GetParent() *Container {
    return nil
}

func (app *App) GetEventMap() event.EventMap {
    return app.EventMap
}

func (app *App) GetKeybindings() event.Keybindings {
    return app.Keybindings
}

func (app *App) GetState() *WidgetState {
    return app.State
}

func (app *App) SetState(state *WidgetState) {
    app.State = state
}

func (app *App) SetEventListener(key tcell.Key, eventName event.EventName, cb event.Callback) {
	app.Keybindings[key] = eventName
	app.EventMap[eventName] = cb
}

