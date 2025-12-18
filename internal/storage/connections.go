package storage

import (
	"encoding/json"
	"os"

	"linkedin-automation-poc/internal/models"
)

const connectionsFile = "connections.json"

func SaveConnection(rec models.ConnectionRecord) error {
	var all []models.ConnectionRecord

	if data, err := os.ReadFile(connectionsFile); err == nil {
		_ = json.Unmarshal(data, &all)
	}

	all = append(all, rec)

	data, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(connectionsFile, data, 0644)
}
