package elements

import (
	"github.com/abdessamad-zgor/dama"
	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Input struct {
	*dama.Widget
	*dama.Editable
	*dama.Navigable
	*dama.Scrollable
    *dama.Stylable
}

func NewInput() Input {
    input := Input{
        dama.NewWidget(),
        new(dama.Editable),
        new(dama.Navigable),
        new(dama.Scrollable),
        &dama.Stylable{
            dama.DefaultStyle,
        },
    }

    input.BorderColor(tcell.ColorDefault)
    return input
}

func (input Input) Render(screen tcell.Screen, context lcontext.Context) {
    box := input.GetBox()
    box.Element = input
    logger.Logger.Println("box: ", box)
    box.Render(screen, context)
}
