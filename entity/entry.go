package entity

type Entry struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	ClubID   string `json:"club_id"`
	Category string `json:"category_group"`
	Guest    bool   `json:"is_guest"`
}
