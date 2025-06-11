package dama

import (
	devent "github.com/abdessaad-zgor/dama/event"
	dutils "github.com/abdessaad-zgor/dama/utils"
)

type EventManager struct {
	App 	*App
	Events 	dutils.ExclusionList[devent.Event]
}


