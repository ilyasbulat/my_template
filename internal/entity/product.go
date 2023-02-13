package entity

import "time"

type Product struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Price         int64             `json:"price"`
	CurrencyID    uint64            `json:"currency_id"`
	Rating        int64             `json:"rating"`
	CategoryID    uint64            `json:"category_id"`
	Specification map[string]string `json:"specification"`
	Image         int64             `json:"image"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}
