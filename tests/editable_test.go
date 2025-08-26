package tests

import (
	"testing"

	"github.com/abdessamad-zgor/dama/logger"
	"github.com/abdessamad-zgor/dama/traits"
)

func TestEditbale(t *testing.T) {
    editable := traits.NewEditable()

    editable.AddRune('h')
    editable.AddRune('e')
    editable.AddRune('l')
    editable.AddRune('l')
    editable.AddRune('o')
    editable.AddRune('\n')
    editable.AddRune('w')
    editable.AddRune('o')
    editable.AddRune('r')
    editable.AddRune('l')
    editable.AddRune('d')

    cursor := editable.Cursor
    contents := editable.Contents

    logger.Logger.Println("contents: ", contents, " cursor: ", cursor)

    if contents != "hello\nworld" {
        t.Error("editable does not insert runes correctely", contents)
    }    

    if cursor.Column != len("world") || cursor.Line != 1 {
        t.Error("editable does not position the cursor correctely", cursor)
    }
} 
