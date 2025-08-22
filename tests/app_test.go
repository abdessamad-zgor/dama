package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/abdessamad-zgor/dama/logger"
	"github.com/abdessamad-zgor/dama"
)

var app dama.DamaApp

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	new_app, err := dama.NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to create new app, because: %v", err))
	}
	app = new_app
}

func teardown() {
	app = new(dama.App)
}

func TestAppRender(t *testing.T) {
	go app.Start()
	app.Exit()
	width, height := app.GetBox().Width, app.GetBox().Height
	logger.Logger.Println("width: ", width, " height: ", height)
	if width <= 0 || height <= 0 {
		t.Fatalf("failed to get screen width and height, %d, %d", app.GetBox().Width, app.GetBox().Height)
	}
}

