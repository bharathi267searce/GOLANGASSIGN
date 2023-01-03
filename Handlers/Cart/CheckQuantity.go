package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func CheckQuantity(Product_id string, Quantity int) bool {
	rows, err := Support.DB.Query("SELECT quantity from inventory WHERE product_id = $1", Product_id)
	if err != nil {
		fmt.Println("errr", err)
	}
	fmt.Println(Product_id, Quantity)
	var w http.ResponseWriter
	for rows.Next() {
		var stock int

		rows.Scan(&stock)
		fmt.Println(stock, "stock")
		if stock == 0 {
			fmt.Println("h1")
			result := fmt.Sprintf("OUT OF STOCK!!!SOLD OUT")
			return false
			json.NewEncoder(w).Encode(result)

		}
		if stock < Quantity {
			fmt.Println("h2")
			result := fmt.Sprintf("The selected quantity is not available please choose less number of items")
			return false
			json.NewEncoder(w).Encode(result)

		} else {
			fmt.Println("h3")
			cartcount := stock - Quantity
			Support.DB.Exec("UPDATE inventory SET quantity=$1 WHERE product_id =$2", cartcount, Product_id)
			return true
		}
		return false
	}

	defer rows.Close()
	return false

}
