package tests

import (
	"testing"
	"github.com/abdessamad-zgor/dama"
)

func TestNavigator(t *testing.T) {
	app.SetLayout(dama.NewGridLayout(2, 2))
	widget1 := dama.NewWidget()
	widget1.SetTag('A')
	widget1.SetTitle("Widget 1")
	app.AddElement(widget1, dama.GridPosition{0, 0, 1, 1})
	widget2 := dama.NewWidget()
	widget2.SetTag('B')
	widget2.SetTitle("Widget 2")
	app.AddElement(widget2, dama.GridPosition{1, 1, 1, 1})

	go app.Start()
	app.Exit()

	if len(app.GetNavigator().Index) != 2 {
		t.Log("app.Navigator.Index = ", len(app.GetNavigator().Index))
	}

	if app.GetNavigator().Current.Element.GetTag() != 'A' {
		t.Log("Wrong current tag bro, we got " + string(app.GetNavigator().Current.Element.GetTag()) + " while it should be 'A'")
	}
}
