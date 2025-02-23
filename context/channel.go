package context

type DispatchEventName string

const (
    HighlightWidget DispatchEventName = "highlight-widget"
    QueueRender     DispatchEventName = "queue-render"
)

type Dispatch struct {
    Event   DispatchEventName
    Payload any
}

var DispatchContextChannel chan Dispatch

func init() {
    DispatchContextChannel = make(chan Dispatch)
}

func DispatchEvent(eventname DispatchEventName, payload any) {
    DispatchContextChannel<-Dispatch{Event: eventname, Payload: payload} 
}
