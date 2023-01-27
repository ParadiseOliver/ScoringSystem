package entity

type Result struct {
	Id       string `json:"id"`
	Athlete  string `json:"athlete"`
	Club     string `json:"club"`
	Category string `json:"category"`
	Agegroup string `json:"agegroup"`
	Score    string `json:"score"`
}
