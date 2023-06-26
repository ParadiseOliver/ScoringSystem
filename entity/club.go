package entity

// Holds details of a single club registered in our database.
type Club struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Headcoach string `json:"headcoach"`
}
