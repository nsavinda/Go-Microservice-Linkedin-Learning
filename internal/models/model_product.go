package models

type Product struct {
	ProductID string `gorm:"primaryKey" json:"ProductId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	VendorID  string `json:"vendorId"`
}
