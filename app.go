package dama

import (
	_ "fmt"
	"os"
	"testing"

	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/event"
	_ "github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type DamaApp interface {
	DamaContainer
	DamaWidget
	Start()
}

type App struct {
	EventChannel chan event.Event
	ExitChannel  chan int
	Screen       tcell.Screen
	State        *WidgetState
	*Container
	Navigator *Navigator
	event.EventMap
	event.Keybindings
	lcontext.Context
}

func initAppScreen() (tcell.Screen, error) {
	isTesting := testing.Testing()
	isDebug := os.Getenv("DEBUG")
	var screen tcell.Screen
	if isTesting || isDebug != "" {
		screen = tcell.NewSimulationScreen("UTF-8")
	} else {
		n_screen, err := tcell.NewScreen()
		if err != nil {
			return nil, err
		}
		screen = n_screen
	}

	return screen, nil
}

func NewApp() (*App, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	app := &App{
		make(chan event.Event),
		make(chan int),
		nil,
		&WidgetState{},
		NewContainer(),
		NewNavigator(),
		make(event.EventMap),
		make(event.Keybindings),
		make(lcontext.Context),
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
	app.Container.Render(app.Screen, app.Context)
	go app.EventLoop()
	_ = <-app.ExitChannel
	app.Screen.Fini()
}

func (app *App) SetupNavigation() {
	navigables := app.GetNavigables()
	app.Navigator.IndexItems(navigables)
	app.Navigator.Setup()
}

func (app *App) StartKeyEventMapper() {
	for {
		app.Screen.Show()
		ev := app.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			key := ev.Key()
			kevent, ok := app.Keybindings[key]
			if ok {
				app.EventChannel <- event.Event{kevent, key, ev}
			}
			if key == tcell.KeyCtrlC {
				app.ExitChannel <- 0
			}
		case *tcell.EventResize:
			app.Screen.Sync()
		}
	}
}

func (app *App) EventLoop() {
	go app.StartKeyEventMapper()
	for {
		select {
		case event := <-app.EventChannel:
			callback, ok := app.EventMap[event.Name]
			if ok {
				callback(app.Context, event)
			}
		case _, _ = <-lcontext.DispatchContextChannel:
		}
		app.Container.Render(app.Screen, app.Context)
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
