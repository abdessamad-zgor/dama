package core

import (
    lcontext "github.com/abdessamad-zgor/dama/context"
    "github.com/gdamore/tcell/v2"
)

type Element interface {
    Render(screen tcell.Screen, context lcontext.Context)
}
