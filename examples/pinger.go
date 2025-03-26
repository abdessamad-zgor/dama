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

    url := new(elements.Input)
    url.SetTitle("URL")
    url.SetTag('U')

    app.AddElement(url, dama.Center)

    app.Start()
}
