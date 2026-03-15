package main

import (
	"fmt"
	"os/exec"
	"github.com/abdessamad-zgor/dama"
)

func main() {
	app, err := dama.NewApp()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app due to : %v\n", err))
	}

	urlEditable := dama.NewEditable()
	url := dama.NewWidget(urlEditable)
	var PingUrl dama.AppEventName = "ping-url"
	url.SetKeybinding(dama.NormalMode, "<CR>", func (match dama.Match) {
		_ = match
		app.DispatchEvent(PingUrl, urlEditable.GetContents())
	})
		
	url.SetTag('U')
	url.SetTitle("URL")

	outputEditable := dama.NewEditable()
	output := dama.NewWidget(outputEditable)
	output.SetAppEvent(PingUrl, func (payload any) {
		urlString, _ := payload.(string)
		cmd := exec.Command("ping", "-c", "1", "-W", "2", urlString)
		outputString, err := cmd.CombinedOutput()
		errString := ""
		if err != nil {
			errString = fmt.Sprintf("Command failed with error: %v, output: %s", err, string(outputString))
		}
		if len(errString) != 0 {
			outputEditable.SetContents(errString)
		} else {
			outputEditable.SetContents(string(outputString))
		}
	})
	
	output.SetTag('O')
	output.SetTitle("Output")

	app.AddElement(url, dama.Top)
	app.AddElement(output, dama.Center)

	app.Start()
}
