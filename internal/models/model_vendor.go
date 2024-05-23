package models

type Vendor struct {
	VendorID string `gorm:"primaryKey" json:"vendorid"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phoneNumber"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
