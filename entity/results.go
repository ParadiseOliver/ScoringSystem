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

type TriScore struct {
	EventID       string   `json:"event_id"`
	AthleteID     string   `json:"athlete_id"`
	ClubID        string   `json:"club_id"`
	CategoryGroup string   `json:"category_group"`
	Routine       string   `json:"routine"`
	E1            string   `json:"e1" form:"e1"`
	E2            string   `json:"e2" form:"e2"`
	E3            string   `json:"e3" form:"e3"`
	E4            string   `json:"e4" form:"e4"`
	HD            []string `json:"hd"`
	DD            string   `json:"dd" form:"dd"`
	Tof           string   `json:"tof" form:"tof"`
	Pen           string   `json:"penalty" form:"penalty"`
	Total         string   `json:"total"`
	Position      int      `json:"position"`
}
