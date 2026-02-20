package orderpool

type EventType string

const (
	EventTypeCreated    EventType = "created"    //订单创建
	EventTypeExpired    EventType = "expired"    //订单过期
	EventTypeProcessing EventType = "processing" //订单处理中
	EventTypeCompleted  EventType = "completed"  //订单完结
	EventTypeRemove     EventType = "remove"     //订单删除
)

func (s EventType) String() string {
	return string(s)
}
