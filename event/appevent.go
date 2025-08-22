package event

type AppEventName string

type AppEvent struct {
	Name	AppEventName
	Payload any
	Handler	Callback
}

