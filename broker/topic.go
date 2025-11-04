package broker

import "time"

type Message struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type Topic struct {
	Name     string
	Messages []Message
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name:     name,
		Messages: []Message{},
	}
}

func (t *Topic) AddMessage(msg string) {
	m := Message{Value: msg, Timestamp: time.Now()}
	t.Messages = append(t.Messages, m)
	saveMessage(t.Name, m)
}
