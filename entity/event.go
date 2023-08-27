package entity

import "time"

// Main struct defining an event (competition) and what groups/categories it will involve.
type Event struct {
	ID             string          `json:"id"`
	Event          string          `json:"event" binding:"required"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	EntryDeadline  time.Time       `json:"entry_deadline"`
	IsPrivate      bool            `json:"is_private"`
	Disciplines    []string        `json:"disciplines"`
	Categories     []string        `json:"categories"`
	Agegroups      []AgeGroup      `json:"agegroups"`
	Genders        []string        `json:"genders"`
	CategoryGroups []CategoryGroup `json:"category_groups"`
}

// Combines a discipline, category, age group and gender into a single group that will perform at an event.
type CategoryGroup struct {
	ID           string `json:"id"`
	DisciplineId string `json:"discipline"`
	CategoryId   string `json:"category"`
	AgegroupId   string `json:"agegroup"`
	GenderId     string `json:"gender"`
}

// Defines a single discipline (TRI, TRS, DMT etc) at an event.
type Discipline struct {
	ID         string `json:"id"`
	Discipline string `json:"discipline"`
}

// Defines a 'level' at an event (e.g. novice, advanced, FIG etc.)
type Category struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

// Defines an age group that can be competed in. May need to convert min and max age to DoB.
type AgeGroup struct {
	ID        string `json:"id"`
	MinAge    int    `json:"min_age"`
	MaxAge    int    `json:"max_age"`
	GroupName string `json:"group_name"`
}

// Defines a gender that can be competed in.
type Gender struct {
	ID     string `json:"id"`
	Gender string `json:"gender"`
}
