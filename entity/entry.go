package entity

// Contains a single entry for an event.
type Entry struct {
	ID       string `json:"id"`
	EventID  string `json:"event_id"`
	UserID   string `json:"user_id"`
	ClubID   string `json:"club_id"` // not needed if this can be got from the users info. unless possible to compete for another club?
	Category string `json:"category_group"`
	Guest    bool   `json:"is_guest"`
}
