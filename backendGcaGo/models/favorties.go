package models

type Favorite struct {
	ID          string `json:"ID"`
	ProductID   string `json:"ProductID"`
	StoreID     string `json:"StoreID"`
	ProductName string `json:"ProductName"`
	ImageURL    string `json:"ImageURL"`
}
