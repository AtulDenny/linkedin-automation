package search

type Criteria struct {
	JobTitle string
	Company  string
	Location string
	Keywords []string
	Page     int
}
