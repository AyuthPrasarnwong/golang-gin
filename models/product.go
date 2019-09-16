package models

type Product struct {
	ProductId     int     `json:"product_id"`
	ProductName   Locale  `json:"product_name"`
	Type          string  `json:"type"`
	Quantity      int     `json:"quantity"`
	Discount      float64 `json:"discount"`
	Vat           float64 `json:"vat"`
	Cost          float64 `json:"cost"`
	Price         float64 `json:"price"`
	TotalPrice    float64 `json:"total_price"`
	Remark        string  `json:"remark"`
	Combos        []Combo `json:"combo"`
}
