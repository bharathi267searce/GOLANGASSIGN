package Datastructures

type ProductInfo struct {
	Product_id    string         `json:"product_id"`
	Name          string         `json:"name"`
	Sku           string         `json:"sku"`
	Category_id   string         `json:"category_id"`
	CateoryName   string         `json:"categoryname"`
	Price         float64        `json:"price"`
	Specification map[string]any `json:"specification"`
}
