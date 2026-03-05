package dama

import (
	_ "fmt"
	"os"
	"testing"
	devent "github.com/abdessamad-zgor/dama/event"
	dkeybinding "github.com/abdessamad-zgor/dama/keybinding"
	"github.com/gdamore/tcell/v2"
	"github.com/abdessamad-zgor/dama/logger"
)

type App interface {
	Container
	Start()
	Exit()
	Resize()
	GetNavigator() *Navigator
	GetScreen() tcell.Screen
	GetEventManager() *EventManager
	SetKeybinding(pattern string, callback devent.KeybindingCallback)
}

type app struct {
	Container
	ExitChannel  	chan int
	Screen       	tcell.Screen
	Navigator 		*Navigator
	EventManager	*EventManager
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

func NewApp() (App, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	app := &app{
		NewContainer(),
		make(chan int),
		nil,
		nil,
		nil,
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
	app.SetBox(0, 0, width, height)
	return app, nil
}

func (app *app) SetKeybinding(pattern string, cb devent.KeybindingCallback) {
	patternMatcher, err := dkeybinding.GetMatcher(pattern)
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

func (app *app) DispatchEvent(eventName devent.AppEventName) {
	app.EventManager.AppEventChannel <- eventName
}

func (app *app) Start() {
	logger.Log("Starting App")
	app.Screen.SetStyle(tcell.StyleDefault)
	//app.Init()
	go app.EventManager.StartEventLoop()
	app.Draw()
	logger.Log("App drawn")
	_ = <-app.ExitChannel
	_, ok := app.Screen.(tcell.SimulationScreen)
	if !ok {
		app.Screen.Fini()
	}
}

func (app *app) Draw() {
	app.Screen.Clear()
	app.Container.Render(app.Screen)
	app.Screen.Show()
}

func (app *app) Resize() {
}

func (app *app) Exit() {
	app.EventManager.Wg.Wait()
	app.ExitChannel <- 0
}

func (app *app) GetParent() *Container {
	return nil
}

func (app *app) GetNavigator() *Navigator {
	return app.Navigator
}

func (app *app) GetEventManager() *EventManager {
	return app.EventManager
}

func (app *app) GetScreen() tcell.Screen {
	return app.Screen
}
