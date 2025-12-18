package models

import "time"

type MessageRecord struct {
	ProfileID  string
	ProfileURL string
	Content    string
	SentAt     time.Time
}
