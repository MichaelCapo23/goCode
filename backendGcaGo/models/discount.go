package models

type Discount struct {
	ID              int
	DiscountID      int
	ProductQuantity string
	ProductDiscount float32
	ProductName     string
	StartDate       string
	EndDate         string
	ProductID       int
}
