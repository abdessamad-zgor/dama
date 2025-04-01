package dama

import (
	"fmt"
	"os"
	"testing"

	"github.com/abdessamad-zgor/dama/logger"
	//"github.com/abdessamad-zgor/dama/elements"
)

var app DamaApp

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	new_app, err := NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to create new app, because: %v", err))
	}
	app = new_app
}

func teardown() {
	app = new(App)
}

func TestAppRender(t *testing.T) {
	app.Start()
    width, height := app.GetBox().Width, app.GetBox().Height 
    logger.Logger.Println("width: ", width, " height: ", height)
	if width <= 0 || height <= 0 {
		t.Fatalf("failed to get screen width and height, %d, %d", app.GetBox().Width, app.GetBox().Height)
	}
}

func TestMescillinous(t *testing.T) {
    ints := []int{}
    ints = append(ints, 1)
    ints = append(ints, 2)
    ints = append(ints, 3)

    logger.Logger.Println("ints: ", ints)
}
