package orderpool

type EventType string

const (
	EventTypeCreated    EventType = "created"
	EventTypeExpired    EventType = "expired"
	EventTypeProcessing EventType = "processing"
	EventTypeRemove     EventType = "remove"
)