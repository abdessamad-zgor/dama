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

    app.AddElement(url, dama.Center)

    app.Start()
}
