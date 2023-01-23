package entity

import "time"

type Event struct {
	Id             string          `json:"id" binding:"required"`
	Name           string          `json:"name" binding:"required"`
	StartDate      time.Time       `json:"start_date" time_format:"02-01-2006"`
	EndDate        time.Time       `json:"end_date" validate:"is-after"`
	EntryDeadline  time.Time       `json:"entry_deadline" binding:"gtfield=StartDate" time_format:"02-01-2006 03:04.05"`
	IsPrivate      bool            `json:"is_private"`
	Disciplines    []string        `json:"disciplines"`
	Categories     []string        `json:"categories"`
	Agegroups      []Agegroup      `json:"agegroups"`
	Genders        []string        `json:"genders"`
	CategoryGroups []CategoryGroup `json:"category_groups"`
}

type Agegroup struct {
	Id           string `json:"id"`
	MinAge       int    `json:"min_age"`
	MaxAge       int    `json:"max_age"`
	CategoryName string `json:"category_name"`
}

type CategoryGroup struct {
	Id           string `json:"id"`
	DisciplineId string `json:"discipline"`
	CategoryId   string `json:"category"`
	AgegroupId   string `json:"agegroup"`
	GenderId     string `json:"gender"`
}
