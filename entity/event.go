package entity

import "time"

type Event struct {
	Id             string          `json:"id"`
	Name           string          `json:"name" binding:"required"`
	StartDate      string          `json:"start_date"`
	EndDate        string          `json:"end_date"`
	EntryDeadline  time.Time       `json:"entry_deadline"`
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
