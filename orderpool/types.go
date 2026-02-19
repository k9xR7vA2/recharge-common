package orderpool

type EventType string

const (
	EventTypeCreated    EventType = "created"
	EventTypeExpired    EventType = "expired"
	EventTypeProcessing EventType = "processing"
	EventTypeRemove     EventType = "remove"
)

func (s EventType) String() string {
	return string(s)
}
