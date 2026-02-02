package main

import (
	"fmt"

	"github.com/abdessamad-zgor/dama"
	"github.com/abdessamad-zgor/dama/traits"
)

type Input struct {
	dama.DamaWidget
	traits.Editable
}



func main() {
	app, err := dama.NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app due to : %v\n", err))
	}

	url := Input {
		dama.NewWidget(),
		traits.NewEditable(),
	}
	url.SetTag('U')
	url.SetTitle("URL")
	url.SetKeybinding("<char>", func (event dama.EventDetail) {

	})

	output := Input {
		dama.NewWidget(),
		traits.NewEditable(),
	}
	output.SetTag('O')
	output.SetTitle("Output")

	app.AddElement(url, dama.Top)
	app.AddElement(output, dama.Center)

	app.Start()
}
