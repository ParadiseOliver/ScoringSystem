package entity

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	Surname   string `json:"surname"`
	Club      string `json:"club"`
	User_type string `json:"userType"`
}
