package entity

type Currency struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
