package models

type Booking struct {
	Id       string `json:"id"`
	Location string `json:"location"`
	ParkName string `json:"parkname"`
	Slot     string `json:"slot"`
	Rate     string `json:"rate"`
	UserId   string `json:"userid"`
	From     string `json:"from"`
	To       string `json:"to"`
}
