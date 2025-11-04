package broker

type Broker struct {
	Topics map[string]*Topic
}

func NewBroker() *Broker {
	return &Broker{
		Topics: make(map[string]*Topic),
	}
}

func (b *Broker) CreateTopic(name string) {
	if _, exists := b.Topics[name]; !exists {
		t := NewTopic(name)
		t.Messages = loadMessages(name)
		b.Topics[name] = t
	}
}

func (b *Broker) Publish(topicName, msg string) {
	if topic, exists := b.Topics[topicName]; exists {
		topic.AddMessage(msg)
	}
}

func (b *Broker) Consume(topicName string) []Message {
	if topic, exists := b.Topics[topicName]; exists {
		return topic.Messages
	}
	return []Message{}
}
