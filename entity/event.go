package entity

type Event struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	StartDate   string   `json:"startdate"`
	EndDate     string   `json:"enddate"`
	Disciplines []string `json:"disciplines"`
}
