package elements

import (
	"fmt"
	"os"
	"testing"

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


func TestInput(t *testing.T) {
    url := NewInput()
    url.SetTag('U')
    url.SetTitle("URL")

    app.AddElement(url, dama.Center)
    app.Start()
}
