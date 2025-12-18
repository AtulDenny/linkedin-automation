package search

import (
	"fmt"

	"linkedin-automation-poc/internal/logger"
	"linkedin-automation-poc/internal/models"
)

type Service struct {
	log *logger.Logger
}

func New(log *logger.Logger) *Service {
	return &Service{log: log}
}

// Search simulates LinkedIn-style profile search.
// In real environments, this would use Rod DOM parsing.
// Here, we return demo profiles for safe evaluation.
func (s *Service) Search(c Criteria) ([]models.Profile, error) {
	s.log.Info(fmt.Sprintf(
		"Searching profiles | title=%s company=%s location=%s page=%d",
		c.JobTitle, c.Company, c.Location, c.Page,
	))

	// Simulated paginated data
	profiles := []models.Profile{
		{
			ID:       fmt.Sprintf("p-%d-1", c.Page),
			Name:     "Alice Johnson",
			Title:    c.JobTitle,
			Company:  c.Company,
			Location: c.Location,
			URL:      "https://linkedin.com/in/alice-demo",
		},
		{
			ID:       fmt.Sprintf("p-%d-2", c.Page),
			Name:     "Bob Smith",
			Title:    c.JobTitle,
			Company:  c.Company,
			Location: c.Location,
			URL:      "https://linkedin.com/in/bob-demo",
		},
	}

	return profiles, nil
}
