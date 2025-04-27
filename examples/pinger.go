package main

import (
	"fmt"

	"github.com/abdessamad-zgor/dama"
	"github.com/abdessamad-zgor/dama/elements"
)

func main() {
	app, err := dama.NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app due to : %v\n", err))
	}

	url := elements.NewInput()
	url.SetTag('U')
	url.SetTitle("URL")

	output := elements.NewTextArea()
	output.SetTag('O')
	output.SetTitle("Output")

	app.AddElement(url, dama.Top)
	app.AddElement(output, dama.Center)

	app.Start()
}
