package models

type Location struct {
	Location    string `json:"location"`
	ParkingName string `json:"ParkingName" gorm:"unique"`
	Dist        string `json:"distance"`
	Slots       uint   `json:"slots"`
	Rate        string `json:"rate"`
}
