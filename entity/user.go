package entity

// Contains a user of the site.
// User type defines if the user is an athlete, coach, admin or a combination of these.
type User struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstname"`
	Surname   string   `json:"surname"`
	Club      string   `json:"club"`
	UserType  []string `json:"user_type"`
	Gender    string   `json:"gender"`
}
