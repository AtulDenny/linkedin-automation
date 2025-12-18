package storage

import (
	"encoding/json"
	"os"

	"linkedin-automation-poc/internal/models"
)

const messagesFile = "messages.json"

func SaveMessage(msg models.MessageRecord) error {
	var all []models.MessageRecord

	if data, err := os.ReadFile(messagesFile); err == nil {
		_ = json.Unmarshal(data, &all)
	}

	all = append(all, msg)

	data, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(messagesFile, data, 0644)
}
