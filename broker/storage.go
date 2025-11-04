package broker

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func saveMessage(topic string, msg Message) error {
	dir := "data"
	os.Mkdir(dir, os.ModePerm)
	file := filepath.Join(dir, fmt.Sprintf("%s.log", topic))

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := json.Marshal(msg)
	_, err = f.Write(append(data, '\n'))
	return err
}
