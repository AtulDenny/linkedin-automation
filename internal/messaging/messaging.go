package messaging

import (
	"strings"
	"time"

	"linkedin-automation-poc/internal/logger"
	"linkedin-automation-poc/internal/models"
	"linkedin-automation-poc/internal/storage"
)

type Service struct {
	log *logger.Logger
}

func New(log *logger.Logger) *Service {
	return &Service{log: log}
}

// SendFollowUp simulates messaging an accepted connection
func (s *Service) SendFollowUp(p models.Profile, template string) error {
	s.log.Info("Checking connection acceptance (demo)")

	// Demo: assume accepted after connect
	time.Sleep(600 * time.Millisecond)

	msg := applyTemplate(template, p)

	s.log.Info("Sending follow-up message")
	time.Sleep(800 * time.Millisecond)

	rec := models.MessageRecord{
		ProfileID:  p.ID,
		ProfileURL: p.URL,
		Content:    msg,
		SentAt:     time.Now(),
	}

	_ = storage.SaveMessage(rec)
	s.log.Info("Message sent (demo)")

	return nil
}

func applyTemplate(tpl string, p models.Profile) string {
	out := tpl
	out = strings.ReplaceAll(out, "{{name}}", p.Name)
	out = strings.ReplaceAll(out, "{{company}}", p.Company)
	return out
}
