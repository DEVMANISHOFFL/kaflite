package broker

import "testing"

func TestCreateTopic(t *testing.T) {
	b := NewBroker()
	b.CreateTopic("test")
	if _, exists := b.Topics["test"]; !exists {
		t.Errorf("expected topic to be created")
	}
}
