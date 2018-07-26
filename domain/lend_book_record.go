package domain

import "time"

// LendBookRecord describe a lend book action in system
type LendBookRecord struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
