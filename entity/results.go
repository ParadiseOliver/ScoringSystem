package entity

// Holds the result of a single participant from one category.
type Result struct {
	ID            string `json:"id"`
	EventID       string `json:"event_id"`
	Athlete       string `json:"athlete"`
	Club          string `json:"club"`
	CategoryGroup string `json:"category_group"`
	Score         string `json:"score"`
}
