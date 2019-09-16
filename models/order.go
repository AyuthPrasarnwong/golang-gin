package models

type Order struct {
	BrandId                 int       `json:"brand_id"`
	BrandName               Locale    `json:"brand_name"`
	OutletId                int       `json:"outlet_id"`
	OutletName              Locale    `json:"outlet_name"`
	TerminalId              int       `json:"terminal_id"`
	TerminalNo              int       `json:"terminal_no"`
	RdCode                  string    `json:"rd_code"`
	Total                   float64   `json:"total"`
	VatTotal                float64   `json:"vat_total"`
	Discount                float64   `json:"discount"`
	DiscountTotal           float64   `json:"discount_total"`
	NetTotal                float64   `json:"net_total"`
	Cash                    int       `json:"cash"`
	ProductTotal            float64   `json:"product_total"`
	PaymentChannelId        int       `json:"payment_channel_id"`
	PaymentChannel          string    `json:"payment_channel"`
	Remark                  string    `json:"remark"`
	CustomerId              int       `json:"customer_id"`
	CustomerName            string    `json:"customer_name"`
	CustomerMobile          string    `json:"customer_mobile"`
	Products                []Product `json:"products"`
	PaymentStatus           string    `json:"payment_status"`
	OrderStatus             string    `json:"order_status"`
	PaymentChannelReference string    `json:"payment_channel_reference"`
	OrderCreatedAt          string    `json:"order_created_at"`
	ClientId                int       `json:"client_id"`
	CustomerType            string    `json:"customer_type"`
}