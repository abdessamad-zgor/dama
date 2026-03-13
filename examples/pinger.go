package main

import (
	"fmt"

	"github.com/abdessamad-zgor/dama"
	_ "github.com/abdessamad-zgor/dama/keybinding"
	_ "github.com/abdessamad-zgor/dama/event"
)

func main() {
	app, err := dama.NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app due to : %v\n", err))
	}

	url := dama.NewWidget(dama.NewEditable())
		
	url.SetTag('U')
	url.SetTitle("URL")

	output := dama.NewWidget(dama.NewEditable())
	
	output.SetTag('O')
	output.SetTitle("Output")

	app.AddElement(url, dama.Top)
	app.AddElement(output, dama.Center)

	app.Start()
}
