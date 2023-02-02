package entity

import "time"

type Event struct {
	ID             string          `json:"id"`
	Name           string          `json:"name" binding:"required"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        string          `json:"end_date"`
	EntryDeadline  time.Time       `json:"entry_deadline"`
	IsPrivate      bool            `json:"is_private"`
	Disciplines    []string        `json:"disciplines"`
	Categories     []string        `json:"categories"`
	Agegroups      []AgeGroup      `json:"agegroups"`
	Genders        []string        `json:"genders"`
	CategoryGroups []CategoryGroup `json:"category_groups"`
}

type Discipline struct {
	ID         string `json:"id"`
	Discipline string `json:"discipline"`
}

type Category struct {
	ID       string `json:"id"`
	Category string `json:"discipline"`
}

type AgeGroup struct {
	ID           string `json:"id"`
	MinAge       int    `json:"min_age"`
	MaxAge       int    `json:"max_age"`
	CategoryName string `json:"category_name"`
}

type Gender struct {
	ID     string `json:"id"`
	Gender string `json:"gender"`
}

type CategoryGroup struct {
	ID           string `json:"id"`
	DisciplineId string `json:"discipline"`
	CategoryId   string `json:"category"`
	AgegroupId   string `json:"agegroup"`
	GenderId     string `json:"gender"`
}
