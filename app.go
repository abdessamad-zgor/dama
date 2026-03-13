package dama

import (
	"fmt"
	"os"
	"testing"
	devent "github.com/abdessamad-zgor/dama/event"
	_ "github.com/abdessamad-zgor/dama/keybinding"
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
	SetKeybinding(mode devent.Mode, pattern string, callback devent.KeybindingCallback)
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

func (app *app) SetKeybinding(mode devent.Mode, pattern string, cb devent.KeybindingCallback) {
	keybinding := devent.KeybindingToEvent(mode, pattern, cb)
	app.EventManager.GlobalEvents.Add(keybinding)
}

func (app *app) DispatchEvent(eventName devent.AppEventName) {
	app.EventManager.AppEventChannel <- eventName
}

func (app *app) Start() {
	logger.Log("Starting App")
	app.Screen.SetStyle(tcell.StyleDefault)
	go app.EventManager.StartEventLoop()
	logger.Log("App rendered")
	_ = <-app.ExitChannel
	logger.Log("Exit signal recieved,  exiting")
	app.EventManager.Wg.Wait()
	_, ok := app.Screen.(tcell.SimulationScreen)
	if !ok {
		app.Screen.Fini()
		logger.Log("App exited, screen fini")
	}
}

func (app *app) Render(screen tcell.Screen) {
	defer func() {
		if x := recover(); x != nil {
			// recovering from a panic; x contains whatever was passed to panic()
			logger.Log(fmt.Sprintf("run time panic: %v", x))

			// if you just want to log the panic, panic again
			panic(x)
		}
	}()
	app.EventManager.Wg.Wait()
	app.Screen.Clear()
	app.Container.Render(screen)
	//app.Screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	//app.Screen.ShowCursor(2,2)
	//app.Screen.SetCursorStyle(tcell.CursorStyleSteadyBar)
	//app.Screen.ShowCursor(3,3)
	app.Screen.Show()
}

func (app *app) Resize() {
}

func (app *app) Exit() {
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
