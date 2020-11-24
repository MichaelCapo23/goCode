package models

type Message struct {
	ID        int    `json:"ID"`
	AppUserID int    `json:"AppUserID"`
	Message   string `json:"Message"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	UUID      string `json:"UUID"`
	Type      string `json:"Type"`
	Resolved  int    `json:"Resolved"`
	CreatedAt string `json:"createdAt"`
}
