package search

import "linkedin-automation-poc/internal/models"

func Deduplicate(input []models.Profile) []models.Profile {
	seen := make(map[string]bool)
	var result []models.Profile

	for _, p := range input {
		if !seen[p.URL] {
			seen[p.URL] = true
			result = append(result, p)
		}
	}
	return result
}
