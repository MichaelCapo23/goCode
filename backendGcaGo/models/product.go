package models

type Product struct {
	ID            int
	ProductName   string
	Description   string
	CategoryID    int
	SubcategoryID int
	Status        int
	Manufacturer  string
	Recommended   int
	Category      string
	Subcategory   string
}
