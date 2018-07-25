package domain

// Category describe category in system
type Category struct {
	Model
	Name string `json:"name"`
}