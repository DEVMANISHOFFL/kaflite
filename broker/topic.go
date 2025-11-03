package broker

type Topic struct {
	Name     string
	Messages []string
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name:     name,
		Messages: []string{},
	}
}

func (t *Topic) AddMessage(msg string) {
	t.Messages = append(t.Messages, msg)
}
