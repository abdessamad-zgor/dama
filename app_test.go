package dama

import (
	"fmt"
	"os"
	"testing"
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
  
}
