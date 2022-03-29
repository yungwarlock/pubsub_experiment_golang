package pubsub

type PubSubMessage struct {
	Message      Message `json:"message"`
	Subscription string  `json:"subscription"`
}

type Message struct {
	Data        []byte            `json:"data,omitempty"`
	Attributes  map[string]string `json:"attributes"`
	PublishTime string            `json:"publishTime"`
	ID          string            `json:"id"`
}
