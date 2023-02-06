package entity

type Result struct {
	ID       string `json:"id"`
	Athlete  string `json:"athlete"`
	Club     string `json:"club"`
	Category string `json:"category"`
	Agegroup string `json:"agegroup"`
	Score    string `json:"score"`
}
