package models

import "time"

type ConnectionRecord struct {
	ProfileID  string
	ProfileURL string
	Note       string
	SentAt     time.Time
}
