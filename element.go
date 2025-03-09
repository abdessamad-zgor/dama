package dama

import (
    lcontext "github.com/abdessamad-zgor/dama/context"
    "github.com/gdamore/tcell/v2"
)

type DamaElement interface {
    Render(screen tcell.Screen, context lcontext.Context)
    GetBox() Box
}
