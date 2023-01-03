type product_master struct {
	Product_id    string         `json:"product_id"`
	Name          string         `json:"name"`
	Sku           string         `json:"sku"`
	Category_id   string         `json:"category_id"`
	Price         float64        `json:"price"`
	Specification map[string]any `json:"specification"`
}
type inventory struct {
	Product_id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}
type category_master struct {
	Category_id   string `json:"category_id"`
	Category_name string `json:"category_name"`
}

type cart_reference struct {
	Reference_id string    `json:"reference_id"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
}

type cart struct {
	Reference_id string `json:"reference_id"`
	Product_id   string `json:"product_id"`
	Quantity     int    `json:"quantity"`
}

