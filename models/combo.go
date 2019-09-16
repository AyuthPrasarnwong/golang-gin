package models

type Combo struct {
	ProductId     int     `json:"product_id"`
	ProductName   Locale  `json:"product_name"`
	CategoryId    int     `json:"category_id"`
	CategoryName  Locale  `json:"category_name"`
	InventoryId   int     `json:"inventory_id"`
	InventoryName Locale  `json:"inventory_name"`
	UnitId        int     `json:"unit_id"`
	UnitName      Locale  `json:"unit_name"`
	Type          string  `json:"type"`
	Quantity      int     `json:"quantity"`
	Discount      float64 `json:"discount"`
	Vat           float64 `json:"vat"`
	Cost          float64 `json:"cost"`
	Price         float64 `json:"price"`
	TotalPrice    float64 `json:"total_price"`
	Remark        string  `json:"remark"`
}