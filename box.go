package dama

import (
	"reflect"

	lcontext "github.com/abdessamad-zgor/dama/context"
	"github.com/abdessamad-zgor/dama/logger"
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	X       uint
	Y       uint
	Width   uint
	Height  uint
	Element DamaElement
}

func (box Box) Render(screen tcell.Screen, context lcontext.Context) {
    stylable, ok := box.Element.(DamaStylable)
    logger.Logger.Println("inside box render", ok, " element type: ", reflect.TypeOf(box.Element))
    if ok {
        boxStyle := stylable.GetStyle()
        logger.Logger.Println("inside box render stylable")
        if boxStyle.Border != nil {
            borderStyle := tcell.StyleDefault.Foreground(boxStyle.Border.Color).Bold(boxStyle.Border.Bold)
            for xi := range box.Width {
                for yi := range box.Height {
                    borderX, borderY := xi+box.X, yi+box.Y
                    if borderX == box.X || borderX == box.X+box.Width-1 {
                        screen.SetContent(int(borderX), int(borderY), tcell.RuneVLine, nil, borderStyle)
                    }
                    if borderY == box.Y || borderY == box.Y+box.Height-1 {
                        screen.SetContent(int(borderX), int(borderY), tcell.RuneHLine, nil, borderStyle)
                    }
                }
            }
            logger.Logger.Println("inside box render border")

            screen.SetContent(int(box.X), int(box.Y), tcell.RuneULCorner, nil, borderStyle)
            screen.SetContent(int(box.X+box.Width-1), int(box.Y), tcell.RuneURCorner, nil, borderStyle)
            screen.SetContent(int(box.X), int(box.Y+box.Height-1), tcell.RuneLLCorner, nil, borderStyle)
            screen.SetContent(int(box.X+box.Width-1), int(box.Y+box.Height-1), tcell.RuneLRCorner, nil, borderStyle)
        }
    }
}
