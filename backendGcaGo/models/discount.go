package models

import "time"

type Discount struct {
	ID        int
	AppUserID int
	Message   string
	CreatedAt time.Time
	Resolved  int
	FirstName string
	LastName  string
	UUID      string
	Type      string
}
