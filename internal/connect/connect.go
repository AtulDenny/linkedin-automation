package connect

import (
	"fmt"
	"time"

	"linkedin-automation-poc/internal/limiter"
	"linkedin-automation-poc/internal/logger"
	"linkedin-automation-poc/internal/models"
	"linkedin-automation-poc/internal/storage"
)

type Service struct {
	log     *logger.Logger
	limiter *limiter.DailyLimiter
}

func New(log *logger.Logger, dailyLimit int) *Service {
	return &Service{
		log:     log,
		limiter: limiter.NewDailyLimiter(dailyLimit),
	}
}

// Send simulates navigating to a profile and sending a connection request.
// In production, this would use Rod DOM interaction.
// For the POC, it is fully logged and persisted.
func (s *Service) Send(p models.Profile, note string) error {
	if err := s.limiter.Allow(); err != nil {
		s.log.Warn(err.Error())
		return err
	}

	s.log.Info(fmt.Sprintf("Navigating to profile: %s", p.URL))
	time.Sleep(800 * time.Millisecond)

	s.log.Info("Clicking Connect button")
	time.Sleep(600 * time.Millisecond)

	if len(note) > 300 {
		note = note[:300]
	}

	s.log.Info("Sending personalized note")
	time.Sleep(700 * time.Millisecond)

	rec := models.ConnectionRecord{
		ProfileID:  p.ID,
		ProfileURL: p.URL,
		Note:       note,
		SentAt:     time.Now(),
	}

	_ = storage.SaveConnection(rec)
	s.log.Info("Connection request sent (demo)")

	return nil
}
