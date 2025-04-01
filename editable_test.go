package dama

import "testing"

func TestEditbale(t *testing.T) {
    editable := NewEditable()

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

    if contents != "hello\nworld" {
        t.Error("editable does not insert runes correctely", contents)
    }    

    if cursor.Column != len("world") {
        t.Error("editable does not position the cursor correctely", *cursor)
    }
} 
