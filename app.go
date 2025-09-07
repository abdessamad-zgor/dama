package dama

import (
	_ "fmt"
	"os"
	"testing"
	devent "github.com/abdessamad-zgor/dama/event"
	dkeystroke "github.com/abdessamad-zgor/dama/keystroke"
	"github.com/gdamore/tcell/v2"
)

type DamaApp interface {
	DamaContainer
	Start()
	Exit()
	GetNavigator() Navigator
	SetKeybinding(pattern string, callback devent.Callback)
}

type App struct {
	*Container
	ExitChannel  	chan int
	Screen       	tcell.Screen
	Navigator 		Navigator
	EventManager	EventManager
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
		NewContainer(),
		ExitChannel: make(chan int),
		Screen: nil,
	}
	app.Navigator = NewNavigator(app)
	app.EventManager = NewEventManager(app)
	screen, err := initAppScreen()
	if err != nil {
		return nil, err
	}
	if err = screen.Init(); err != nil {
		return nil, err
	}
	app.Screen = screen
	width, height := app.Screen.Size()
	app.X = 0
	app.Y = 0
	app.Width = (width)
	app.Height = (height)
	return app, nil
}

func (app *App) SetKeybinding(pattern string, cb devent.Callback) {
	patternMatcher, err := dkeystroke.GetMatcher(pattern)
	if err != nil {
		panic(err)
	}
	keybinding := devent.DamaEvent{
		devent.DKeybinding,
		devent.EventDetail{
			&devent.Keybinding{
				pattern,
				patternMatcher,
				cb,
			},
			nil,
		},
	}

	app.EventManager.GlobalEvents.Add(keybinding)
}

func (app *App) DispatchEvent(eventName devent.AppEventName) {
	app.EventManager.AppEventChannel <- eventName
}

func (app *App) Start() {
	app.Screen.SetStyle(tcell.StyleDefault)
	app.Init()
	go app.EventManager.EventLoop()
	app.Draw()
	_ = <-app.ExitChannel
	_, ok := app.Screen.(tcell.SimulationScreen)
	if !ok {
		app.Screen.Fini()
	}
}

func (app *App) Draw() {
	app.Screen.Clear()
	app.Container.Render(app.Screen)
	app.Screen.Show()
}

func (app *App) Resize() {
}

func (app *App) Exit() {
	app.ExitChannel <- 0
}

func (app *App) GetParent() *Container {
	return nil
}


func (app *App) GetNavigator() Navigator {
	return app.Navigator
}

