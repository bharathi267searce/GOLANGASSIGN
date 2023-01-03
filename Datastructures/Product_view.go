package Datastructures

type Product_view struct {
	Product_id    string  `json:"product_id"`
	Price         float32 `json:"price"`
	Category_name string  `json:"cateory_name"`
	Quantity      int     `json:"quantity"`
}
