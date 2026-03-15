package tests

import (
	"testing"

	"github.com/samazee/dama"
)

func TestEditable(t *testing.T) {
    editable := dama.NewEditable()

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

    cursor := editable.GetCursor()
    contents := editable.GetContents()

    if contents != "hello\nworld" {
        t.Error("editable does not insert runes correctly", contents)
    }    

    if cursor.Column != len("world") || cursor.Line != 1 {
        t.Error("editable does not position the cursor correctely", cursor)
    }
	
	editable.RemoveRune()
    cursor = editable.GetCursor()
    contents = editable.GetContents()

    if contents != "hello\nworl" {
        t.Error("editable does not insert runes correctly", contents)
    }    

    if cursor.Column != len("worl") || cursor.Line != 1 {
        t.Error("editable does not position the cursor correctely", cursor)
    }
} 
