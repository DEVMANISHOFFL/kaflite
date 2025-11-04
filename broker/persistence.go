package broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func loadMessages(topic string) []Message {
	file := fmt.Sprintf("data/%s.log", topic)
	data, err := os.ReadFile(file)
	if err != nil {
		return []Message{}
	}

	var msgs []Message
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if len(lines) == 0 {
			continue
		}
		var m Message
		json.Unmarshal(line, &m)
		msgs = append(msgs, m)
	}
	return msgs
}
