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
	EventID       string  `json:"event_id"`
	AthleteID     string  `json:"athlete_id"`
	ClubID        string  `json:"club_id"`
	CategoryGroup string  `json:"category_group"`
	Routine       string  `json:"routine"`
	E1            float32 `json:"e1" form:"e1"`
	E2            float32 `json:"e2" form:"e2"`
	E3            float32 `json:"e3" form:"e3"`
	E4            float32 `json:"e4" form:"e4"`
	HD            float32 `json:"hd" form:"hd"`
	DD            float32 `json:"dd" form:"dd"`
	Tof           float32 `json:"tof" form:"tof"`
	Pen           float32 `json:"penalty" form:"penalty"`
	Total         float32 `json:"total"`
	Position      int     `json:"position"`
}
