package dama

import (
	_ "fmt"
	"os"
	"testing"

	"github.com/gdamore/tcell/v2"
)

type DamaApp interface {
	DamaContainer
	Start()
	Exit()
	GetNavigator() Navigator
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
		make(chan int),
		nil,
		NewNavigator(),
		NewEventManager(),
	}
	app.Navigator.Root.Element = app 
	app.EventManager.App = app
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

func (app *App) Start() {
	app.Screen.SetStyle(tcell.StyleDefault)
	app.SetupNavigation()
	app.Draw()
	go app.EventManager.EventLoop()
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

func (app *App) SetupNavigation() {
	//navigables := app.GetNavigables()
	//app.Navigator.GetNavigationTree(navigables)
	//tags := app.Navigator.SetupKeybindings()
	//// TODO: replace this section
	//if len(navigables) >= 1 {
	//	app.Navigate(navigables[0].GetTag())
	//}
}

func (app *App) Navigate(tag rune) {
	//previousWidget, pOk := app.Navigator.Current.Element.(DamaWidget)
	//if app.Navigator.Navigate(tag) {
	//	app.UpdateNavigator()
	//	widget, ok := app.Navigator.Current.Element.(DamaWidget)
	//	if pOk {
	//		previousEventMap := previousWidget.GetEventMap()
	//		previousKeybindings := previousWidget.GetKeybindings()
	//		for key, _ := range previousKeybindings {
	//			delete(app.Keybindings, key)
	//		}

	//		for key, _ := range previousEventMap {
	//			delete(app.EventMap, key)
	//		}
	//	}

	//	if ok {
	//		elementEventMap := widget.GetEventMap()
	//		elementKeybindings := widget.GetKeybindings()
	//		for key, value := range elementKeybindings {
	//			app.SetKeybinding(key, value)
	//		}

	//		for key, value := range elementEventMap {
	//			app.SetEventCallback(key, value)
	//		}
	//	}
	//}
}

func (app *App) UpdateNavigator() {
	//tags := app.Navigator.SetupKeybindings()
}

func (app *App) GetParent() *Container {
	return nil
}

func (app *App) GetNavigator() Navigator {
	return app.Navigator
}

